package main

import (
	"encoding/json"
	"github.com/spf13/cobra"

	"relayer/appchains"
	"relayer/appchains/fisco"
	cfg "relayer/config"
	"relayer/core"
	"relayer/hub"
	"relayer/logging"
	"relayer/server"
	"relayer/store"
)

// StartCmd implements the start command
func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [config-file]",
		Short: "Start the relayer daemon",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 0 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[0]
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			appChainType := config.GetString(cfg.ConfigKeyAppChainType)

			store, err := store.NewStore(config.GetString(cfg.ConfigKeyStorePath))
			if err != nil {
				return err
			}

			appChainFactory := appchains.NewAppChainFactory(store)
			hubChain := hub.BuildIritaHubChain(hub.NewConfig(config))
			relayerInstance := core.NewRelayer(appChainType, hubChain, appChainFactory, logging.Logger)

			if appChainType=="fisco" {
				fiscoBaseConfig, err := fisco.NewBaseConfig(config)
				if err != nil {
					return err
				}
				baseConfigByte, _ := json.Marshal(fiscoBaseConfig)
				appChainFactory.StoreBaseConfig(appChainType, baseConfigByte)
				chainIDsbz, _ := store.Get([]byte("chainIDs"))
				if chainIDsbz == nil {
					chainIDsbz,err = json.Marshal(map[string]string{})
					if err != nil {
						return err
					}
					store.Set([]byte("chainIDs"), chainIDsbz)
				}else{
					chainIDs:= map[string]string{}
					json.Unmarshal(chainIDsbz, &chainIDs)
					for chainID, chainType := range chainIDs{
						if chainType == appChainType{
							chainParams,err := store.Get(fisco.ChainParamsKey(chainID))
							if err != nil {
								return err
							}
							chain, err := relayerInstance.AppChainFactory.BuildAppChain(chainType, chainParams)
							if err != nil {
								return err
							}

							if err := chain.Start(relayerInstance.HandleInterchainRequest); err != nil {
								return err
							}

							relayerInstance.AppChains[chainID] = chain
							relayerInstance.AppChainStates[chainID] = true
						}
					}
				}
			}
			chainManager := server.NewChainManager(relayerInstance)
			marketManger := server.NewMarketManager(relayerInstance)

			server.StartWebServer(chainManager, marketManger)

			return nil
		},
	}

	return cmd
}
