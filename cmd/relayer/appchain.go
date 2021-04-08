package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"relayer/appchains"
	cfg "relayer/config"
	"relayer/store"
	"relayer/appchains/fisco"
)

var (
	AppChainCmd = &cobra.Command{
		Use:   "appchain",
		Short: "appchain commands",
	}
)

// AddServiceBindingCmd implements the appchain add-svc-binding command
func AddServiceBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-svc-binding [chain-id] [svc-name] [schemas] [provider] [svc-fee] [qos] [config-file]",
		Short: "Add the specified service binding",
		Args:  cobra.RangeArgs(6, 7),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 6 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[6]
			}

			qos, err := strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				return err
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}
			store, err := store.NewStore(config.GetString(cfg.ConfigKeyStorePath))
			if err != nil {
				return err
			}

			appChainFactory := appchains.NewAppChainFactory(store)
			chainParams,err := store.Get(fisco.ChainParamsKey(args[0]))
			if err != nil {
				return err
			}
			appChain, err := appChainFactory.BuildAppChain(config.GetString(cfg.ConfigKeyAppChainType), chainParams)
			if err != nil {
				return err
			}

			err = appChain.AddServiceBinding(args[1], args[2], args[3], args[4], qos)
			if err != nil {
				return err
			}

			fmt.Printf("AddServiceBinding transaction minted\n")

			return nil
		},
	}

	return cmd
}

// UpdateServiceBindingCmd implements the appchain update-svc-binding command
func UpdateServiceBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-svc-binding [chain-id] [svc-name] [provider] [svc-fee] [qos] [config-file]",
		Short: "Update the specified service binding",
		Args:  cobra.RangeArgs(5, 6),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 5 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[5]
			}

			qos, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return err
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			store, err := store.NewStore(config.GetString(cfg.ConfigKeyStorePath))
			if err != nil {
				return err
			}
			appChainFactory := appchains.NewAppChainFactory(store)
			chainParams,err := store.Get(fisco.ChainParamsKey(args[0]))
			if err != nil {
				return err
			}
			appChain, err := appChainFactory.BuildAppChain(config.GetString(cfg.ConfigKeyAppChainType), chainParams)
			if err != nil {
				return err
			}

			err = appChain.UpdateServiceBinding(args[1], args[2], args[3], qos)
			if err != nil {
				return err
			}

			fmt.Printf("UpdateServiceBinding transaction minted\n")

			return nil
		},
	}

	return cmd
}

func init() {
	AppChainCmd.AddCommand(
		AddServiceBindingCmd(),
		UpdateServiceBindingCmd(),
	)
}
