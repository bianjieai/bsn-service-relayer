package main

import (
	"github.com/spf13/cobra"
	"relayer/appchains"
	cfg "relayer/config"
	"relayer/hub"
	"relayer/logging"
	"relayer/server"
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

			//path :=config.GetString(cfg.ConfigKeyStorePath)
			//store, err := store.NewStore(path)
			//if err != nil {
			//	return err
			//}

			//appChainFactory := appchains.NewAppChainFactory(store)

			hubChain := hub.BuildIritaHubChain(hub.NewConfig(config))

			//relayerInstance := core.NewRelayer(appChainType, hubChain, appChainFactory, logging.Logger)

			chainManager := appchains.NewAppChainHandler(appChainType, &hubChain, logging.Logger, config) //server.NewChainManager(relayerInstance)

			httpPort := config.GetInt(cfg.ConfigKeyHttpPort)

			if httpPort == 0 {
				httpPort = 80
			}

			server.StartWebServer(chainManager, httpPort)

			return nil
		},
	}

	return cmd
}
