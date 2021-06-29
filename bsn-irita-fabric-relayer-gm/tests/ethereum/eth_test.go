package ethereum_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cfg "relayer/config"
	ethtesting "relayer/tests/ethereum"
)

type EthTestSuite struct {
	suite.Suite

	config *viper.Viper
	client *ethclient.Client

	exchangeRateSvcContract *ethtesting.ExchangeRateService
}

func TestEthTestSuite(t *testing.T) {
	suite.Run(t, new(EthTestSuite))
}

func (suite *EthTestSuite) SetupTest() {
	config, err := cfg.LoadYAMLConfig("../config.yaml")
	suite.NoError(err)

	suite.config = config

	client, err := ethclient.Dial(config.GetString("eth.node_rpc_addr"))
	suite.NoError(err)

	suite.client = client

	exchangeRateSvcAddr := ethcmn.HexToAddress(config.GetString("eth.exchange_rate_service_addr"))
	exchangeRateSvcContract, err := ethtesting.NewExchangeRateService(exchangeRateSvcAddr, client)
	suite.NoError(err)

	suite.exchangeRateSvcContract = exchangeRateSvcContract
}

func (suite *EthTestSuite) buildAuthTransactor() *bind.TransactOpts {
	privKey, err := crypto.HexToECDSA(suite.config.GetString("eth.key"))
	suite.NoError(err)

	auth := bind.NewKeyedTransactor(privKey)

	nextNonce, err := suite.client.PendingNonceAt(context.Background(), auth.From)
	suite.NoError(err)

	auth.GasLimit = suite.config.GetUint64("eth.gas_limit")
	auth.GasPrice = big.NewInt(int64(suite.config.GetUint64("eth.gas_price")))
	auth.Nonce = big.NewInt(int64(nextNonce))

	return auth
}

func (suite *EthTestSuite) TestSendInterchainRequest() {
	auth := suite.buildAuthTransactor()

	tx, err := suite.exchangeRateSvcContract.SendRequest(auth)
	suite.NoError(err)

	fmt.Printf("tx sent successfully: %s", tx.Hash().Hex())

	_, err = bind.WaitMined(context.Background(), suite.client, tx)
	suite.NoError(err)
}
