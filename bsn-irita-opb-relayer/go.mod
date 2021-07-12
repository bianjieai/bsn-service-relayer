module relayer

go 1.14

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/bianjieai/irita-sdk-go v1.1.1-0.20210707070124-79ed0124b3de
	github.com/cockroachdb/pebble v0.0.0-20201118202804-75ede898b66c
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.0
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.3
)

replace (
	github.com/bianjieai/irita-sdk-go => /Users/bianjie/irita/irita-sdk-go
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
	github.com/ugorji/go => github.com/ugorji/go v1.1.2
)
