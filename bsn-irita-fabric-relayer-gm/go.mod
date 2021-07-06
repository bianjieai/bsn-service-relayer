module relayer

go 1.14

require (
	github.com/BSNDA/fabric-sdk-go-gm v0.0.0
	github.com/Shopify/sarama v1.28.0 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/aristanetworks/goarista v0.0.0-20210107181124-fad53805024e // indirect
	github.com/cockroachdb/pebble v0.0.0-20201118202804-75ede898b66c
	github.com/ethereum/go-ethereum v1.9.18
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3 // indirect
	github.com/hyperledger/fabric v1.4.3
	github.com/hyperledger/fabric-protos-go v0.0.0-20200124220212-e9cfc186ba7b
	github.com/irisnet/service-sdk-go v1.0.1-0.20210416090657-1bdf41efe743
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil v3.20.12+incompatible // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tjfoc/gmsm v1.4.0
)

replace (
	github.com/BSNDA/fabric-sdk-go-gm => github.com/chenxifun/fabric-sdk-go-gm v1.4.3-bsn-0.2
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
	github.com/tjfoc/gmsm => github.com/chenxifun/gmsm v1.4.0
	github.com/tjfoc/gmtls => github.com/chenxifun/gmtls v1.2.1-0.20210427064604-124283070ca7
	github.com/ugorji/go => github.com/ugorji/go v1.1.2
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20200513103714-09dca8ec2884
	google.golang.org/grpc => google.golang.org/grpc v1.31.0
)
