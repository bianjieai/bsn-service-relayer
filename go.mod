module relayer

go 1.14

require (
	github.com/bianjieai/irita-sdk-go v0.0.0-20200807094603-3750279a4002
	github.com/ethereum/go-ethereum v1.9.18
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	google.golang.org/genproto v0.0.0-20191216205247-b31c10ee225f // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200529.0.20200616020441-2ea9cf4018bb
