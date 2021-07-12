package opb

import (
	"fmt"

	"relayer/store"
)

const (
	StorePrefix = ChainType

	KeyBaseConfig        = "baseconfig"
	KeyPrefixChainParams = "params"
	KeyPrefixHeight      = "height"
)

// BaseConfigKey returns the key for the FISCO base config
func BaseConfigKey() []byte {
	return []byte(fmt.Sprintf("%s:%s", StorePrefix, KeyBaseConfig))
}

// ChainParamsKey returns the key for the params of the given chain
func ChainParamsKey(chainID string) []byte {
	return []byte(fmt.Sprintf("%s:%s:%s", StorePrefix, KeyPrefixChainParams, chainID))
}

// HeightKey returns the key for the height of the specified chain
func HeightKey(chainID string) []byte {
	return []byte(fmt.Sprintf("%s:%s:%s", StorePrefix, KeyPrefixHeight, chainID))
}

// StoreBaseConfig stores the base config
func StoreBaseConfig(store *store.Store, baseConfig []byte) error {
	err := ValidateBaseConfig(baseConfig)
	if err != nil {
		return err
	}

	return store.Set(BaseConfigKey(), baseConfig)
}
