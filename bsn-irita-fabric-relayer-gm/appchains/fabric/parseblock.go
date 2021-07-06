package fabric

import (
	"crypto/sha256"
	x509 "github.com/tjfoc/gmsm/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/BSNDA/fabric-sdk-go-gm/third_party/github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/BSNDA/fabric-sdk-go-gm/third_party/github.com/hyperledger/fabric/core/ledger/util"
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"relayer/appchains/fabric/utils"
)

func ParseBlock(block *cb.Block) (*BlockData, error) {

	blockData := &BlockData{}

	blockHeader := block.GetHeader()
	blockData.BlockNumber = blockHeader.GetNumber()

	blockData.BlockHash = hex.EncodeToString(GetBlockHASH(block))
	blockData.BlockPreviousHash = hex.EncodeToString(blockHeader.GetPreviousHash())

	blockData.BlockSize = uint64(len(block.String()))

	var transactions []*TransactionInfo
	var tranNo int64 = -1
	txsFilter := util.TxValidationFlags(block.Metadata.Metadata[cb.BlockMetadataIndex_TRANSACTIONS_FILTER])
	if len(txsFilter) == 0 {
		txsFilter = util.NewTxValidationFlags(len(block.Data.Data))
		block.Metadata.Metadata[cb.BlockMetadataIndex_TRANSACTIONS_FILTER] = txsFilter
	}

	for _, envBytes := range block.Data.Data {
		//fmt.Println("envBytes",envBytes)
		tranNo++
		vcode := txsFilter.Flag(int(tranNo))
		trans, err := parseTransaction(envBytes)

		if err == nil {
			trans.Status = int(vcode)
			transactions = append(transactions, trans)
		}

	}

	blockData.Transactions = transactions
	blockData.BlockTxCount = len(transactions)

	return blockData, nil

}

func parseTransaction(envBytes []byte) (*TransactionInfo, error) {
	trans := &TransactionInfo{}
	var err error
	var env *cb.Envelope

	if env, err = utils.GetEnvelopeFromBlock(envBytes); err != nil {
		return nil, err
	}

	var payload *cb.Payload
	if payload, err = utils.GetPayload(env); err != nil {
		return nil, err
	}

	var chdr *cb.ChannelHeader
	if chdr, err = utils.UnmarshalChannelHeader(payload.Header.ChannelHeader); err != nil {
		return nil, err
	}

	trans.ChannelId = chdr.ChannelId
	trans.TxId = chdr.TxId

	trans.TimeSpanSec = chdr.Timestamp.Seconds
	trans.TimeSpanNsec = int64(chdr.Timestamp.Nanos)
	//orderer提交者

	//var shdr *common.SignatureHeader
	//if shdr, err = utils.GetSignatureHeader(payload.Header.SignatureHeader); err != nil {
	//	return nil,err
	//}

	if cb.HeaderType(chdr.Type) == cb.HeaderType_ENDORSER_TRANSACTION {

		var tx *peer.Transaction
		if tx, err = utils.GetTransaction(payload.Data); err != nil {
			return nil, err
		}

		if len(tx.Actions) > 0 {

			trans.IsTranasction = true

			action := tx.Actions[0]

			//AShdr, err := utils.GetSignatureHeader(action.Header)
			_, err := utils.GetSignatureHeader(action.Header)

			if err != nil {
				return nil, err
			}

			////获取交易的提交者
			//var subject string //mspid
			//if _, subject, err = decodeSerializedIdentity(AShdr.Creator); err != nil {
			//	return nil, err
			//}
			//trans.CreateName = subject

			//var capayload *peer.ChaincodeActionPayload
			var ca *peer.ChaincodeAction
			if _, ca, err = utils.GetPayloads(action); err != nil {
				return nil, err
			}
			if ca.Events != nil {
				ev, err := utils.GetChaincodeEvents(ca.Events)
				if err == nil {
					event := &ChaincodeEvent{
						EventName:   ev.EventName,
						ChaincodeId: ev.ChaincodeId,
						TxId:        ev.TxId,
					}

					trans.Events = event
				}
			}

			txRWSet := &rwsetutil.TxRwSet{}
			if err = txRWSet.FromProtoBytes(ca.Results); err != nil {
				return nil, err
			}

			for _, nsRWSet := range txRWSet.NsRwSets {
				ns := nsRWSet.NameSpace

				if ns != "lscc" {
					r, w := parseReadWrite(nsRWSet.KvRwSet)

					nss := &NameSpaceSet{
						NameSpace: ns,
						Reads:     r,
						Writes:    w,
					}

					trans.NameSpaceSets = append(trans.NameSpaceSets, nss)

				}
			}
		}

	}

	return trans, nil

}

func parseReadWrite(kvrw *kvrwset.KVRWSet) ([]*ReadSet, []*WriteSet) {

	var rs []*ReadSet
	var ws []*WriteSet

	for _, kvRead := range kvrw.Reads {
		r := &ReadSet{
			Key: kvRead.Key,
		}
		rs = append(rs, r)
	}

	for _, kvWrite := range kvrw.Writes {
		w := &WriteSet{
			Key:      kvWrite.Key,
			Value:    string(kvWrite.Value),
			IsDelete: kvWrite.IsDelete,
		}
		ws = append(ws, w)
	}

	return rs, ws

}

func decodeSerializedIdentity(creator []byte) (string, string, error) {

	si := &msp.SerializedIdentity{}

	err := proto.Unmarshal(creator, si)
	if err != nil {
		return "", "", err
	}

	mspId := si.Mspid

	dcert, _ := pem.Decode(si.IdBytes)

	x509Cert, err := x509.ParseCertificate(dcert.Bytes)

	if err != nil {
		return "", "", err
	}

	subject := x509Cert.Subject.CommonName

	return mspId, subject, nil

}

type BlockData struct {
	//块号
	BlockNumber uint64
	//块哈希
	BlockHash string
	//上一个块的哈希
	BlockPreviousHash string
	//块大小
	BlockSize uint64
	//块的交易数量
	BlockTxCount int

	Transactions []*TransactionInfo
}

func (b *BlockData) GetTrans(txid string) *TransactionInfo {

	fmt.Println(len(b.Transactions))

	for i, tran := range b.Transactions {
		fmt.Println(tran.TxId)
		if tran.TxId == txid {
			return b.Transactions[i]
		}
	}
	return nil
}

type TransactionInfo struct {
	TxId string

	Status int

	ChannelId string

	ChaincodeId   Chaincode
	TimeSpanSec   int64
	TimeSpanNsec  int64
	IsTranasction bool

	Events *ChaincodeEvent

	CreateName string
	//CreateCert string

	NameSpaceSets []*NameSpaceSet
}

type ChaincodeEvent struct {
	EventName   string
	ChaincodeId string
	TxId        string
}

type Chaincode struct {
	Name    string
	Version string
}

type NameSpaceSet struct {
	NameSpace string

	Reads  []*ReadSet
	Writes []*WriteSet
}

type ReadSet struct {
	Key string
}

type WriteSet struct {
	Key      string
	Value    string
	IsDelete bool
}

type asn1Header struct {
	Number       int64
	PreviousHash []byte
	DataHash     []byte
}

func GetBlockHASH(info *cb.Block) []byte {
	asn1Header := asn1Header{
		PreviousHash: info.Header.PreviousHash,
		DataHash:     info.Header.DataHash,
		Number:       int64(info.Header.Number),
	}

	result, err := asn1.Marshal(asn1Header)
	if err != nil {

	}
	hash := GetSHA256HASH(result)
	return hash
}

func GetSHA256HASH(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hash := h.Sum(nil)
	return hash
}
