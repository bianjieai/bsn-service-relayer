package bcos_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	bcosclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"

	"relayer/appchains/bcos/iservice"
)

type BCOSTestSuite struct {
	suite.Suite
	client *bcosclient.Client
}

func TestBCOSTestSuite(t *testing.T) {
	suite.Run(t, new(BCOSTestSuite))
}

func (suite *BCOSTestSuite) SetupTest() {
	configs, err := conf.ParseConfigFile("config.toml")
	suite.NoError(err)

	client, err := bcosclient.Dial(&configs[0])
	suite.NoError(err)

	suite.client = client
}

func (suite *BCOSTestSuite) TestDeployIServiceContracts() {
	transactOpts := suite.client.GetTransactOpts()

	// deploy iservice core extension contract
	address, tx, _, err := iservice.DeployIServiceCoreEx(transactOpts, suite.client, transactOpts.From)
	suite.NoError(err)

	fmt.Printf("iservice core extension deployed, contract address: %s, tx hash: %s", address.Hex(), tx.Hash().Hex())

	// deploy iservice market extension contract
	address, tx, _, err = iservice.DeployIServiceMarketEx(transactOpts, suite.client)
	suite.NoError(err)

	fmt.Printf("iservice market extension deployed, contract address: %s, tx hash: %s", address.Hex(), tx.Hash().Hex())
}
