module relayer

go 1.14

require (
	github.com/ethereum/go-ethereum v1.9.18
	github.com/irisnet/service-sdk-go v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc3.0.20200907055413-3359e0bf2f84 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	// github.com/irisnet/service-sdk-go => github.com/secret2830/service-sdk-go v0.0.0-20200930025908-91ed6ca17b1b
	github.com/irisnet/service-sdk-go => /Users/bianjie/develop/service-sdk-go-2/service-sdk-go
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703.0.20200925112439-d4196a88a285
)
