var Web3 = require('web3');
var Tx = require('ethereumjs-tx').Transaction;
var fs = require('fs');
var log4js = require('log4js');

var logger = log4js.getLogger('normal');
logger.level = 'info';

// ethereum config
var chainID = 'ropsten';
var nodeRPCAddr = 'wss://ropsten.infura.io/ws/v3/56e89587eacb4fbe8655e4c44b146237'
var fromAddress = '0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b';
var privateKey = '5dee232c8be5cb81f0ae6fddb45243fc6208192c16aef275ef41b019df765d1f';
var gasPrice = 20000000000;
var gasLimit = 500000;

// contract config
var consumerContractAddr = '0x87eFCd928Fa49B50177FFd7db144f6cEc051cB55';
var abiPath = './artifacts/NFTServiceConsumer.json';

// web3 instance
var web3 = new Web3(nodeRPCAddr);

// abi
var abi = JSON.parse(fs.readFileSync(abiPath));

// contract address passed in
var args = process.argv.splice(2);
if (args.length > 0) {
    consumerContractAddr = args[0];
}

// consumer contract instance
var consumerContract = new web3.eth.Contract(abi, consumerContractAddr);

// nft variables
var destAddress = '0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b';
var amount = 1;
var metaID = '-Z-2fJxzCoFJ0MOU-zA3-tiIh7dK6FjDruAxgxW6PEs';
var setPrice = 1; // price in USDT
var isForSale = true;

// transaction data for minting nft 
// var data = consumerContract.methods.mint(
//     destAddress,
//     amount,
//     metaID,
//     setPrice,
//     isForSale
// ).encodeABI();

var data = consumerContract.methods.mintV2(
    '{"header":{},"body":{"to":"0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b","amount_to_mint":"1","meta_id":"test"}}'
).encodeABI();

// build, sign and send transaction
web3.eth.getTransactionCount(fromAddress)
.then(
    nonce => {
        // build raw tx
        var rawTx = {
            from: fromAddress,
            nonce: nonce,
            gasPrice: web3.utils.toHex(gasPrice),
            gasLimit: web3.utils.toHex(gasLimit),
            to: consumerContractAddr,
            value: '0x0',
            data: data,
        };
        
        // sign transaction
        var privKey = new Buffer.from(privateKey, 'hex');
        var tx = new Tx(rawTx, {chain: chainID});
        tx.sign(privKey);
        var serializedTx = tx.serialize();
        
        logger.info(
            'starting to mint NFT, to:%s, amount:%d, metaID:%s, setPrice:%s, isForSale:%s',
            destAddress,
            amount,
            metaID,
            setPrice,
            isForSale
        );

        // initiate nft minting transaction
        web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'))
        .on('transactionHash', function(hash){
            logger.info('nft minting tx sent, tx hash: %s', hash);
        })
        .on('error', logger.error);
})
.catch(console.error);

// listen to events
consumerContract.events.allEvents()
.on('data', function(event){
    switch (event.event) {
        case 'IServiceRequestSent':
            logger.info('iservice request sent, request id: %s', event.returnValues._requestID);
            break;
        
        case 'PriceSet':
            logger.info('usdt-eth price got: %s', event.returnValues._price);
            break;
        
        case 'NFTMinted':
            logger.info('nft minted: %s', event.returnValues._nftID);
            break;
        
        default:
            logger.info('event triggered: %s', event.event);
    }
})
.on('error', logger.error);
