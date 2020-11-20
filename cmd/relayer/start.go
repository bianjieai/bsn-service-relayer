package main

import (
	"github.com/spf13/cobra"

	"relayer/appchains"
	cfg "relayer/config"
	"relayer/core"
	"relayer/hub"
	"relayer/logging"
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

			store, err := store.NewStore(config.GetString(cfg.ConfigKeyStorePath))
			if err != nil {
				return err
			}

			appChainFactory := appchains.NewAppChainFactory(config, store)
			appChain, err := appChainFactory.Make(config.GetString(cfg.ConfigKeyAppChainName))
			if err != nil {
				return err
			}

			hubChain := hub.MakeIritaHubChain(hub.NewConfig(config))

			relayerInstance := core.NewRelayer(hubChain, appChain, logging.Logger)
			relayerInstance.Start()

			return nil
		},
	}

	return cmd
}
