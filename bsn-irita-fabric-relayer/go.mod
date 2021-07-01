module relayer

go 1.14

require (
	github.com/BSNDA/fabric-sdk-go-gm v0.0.0
	github.com/FISCO-BCOS/go-sdk v0.10.0
	github.com/cockroachdb/pebble v0.0.0-20201118202804-75ede898b66c
	github.com/ethereum/go-ethereum v1.9.18 // indirect
	github.com/fsouza/go-dockerclient v1.7.3 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.0
	github.com/hyperledger/fabric v1.4.3
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric v0.0.0-20190822125948-d2b42602e52e
	github.com/irisnet/service-sdk-go v1.0.1-0.20210416090657-1bdf41efe743
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/BSNDA/fabric-sdk-go-gm => github.com/chenxifun/fabric-sdk-go-gm v1.4.3-bsn-0.2
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
	github.com/ugorji/go => github.com/ugorji/go v1.1.2
)
