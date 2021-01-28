package fisco_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	fiscoclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"

	"relayer/appchains/fisco/iservice"
)

type FISCOTestSuite struct {
	suite.Suite
	client *fiscoclient.Client
}

func TestFISCOTestSuite(t *testing.T) {
	suite.Run(t, new(FISCOTestSuite))
}

func (suite *FISCOTestSuite) SetupTest() {
	configs, err := conf.ParseConfigFile("config.toml")
	suite.NoError(err)

	client, err := fiscoclient.Dial(&configs[0])
	suite.NoError(err)

	suite.client = client
}

func (suite *FISCOTestSuite) TestDeployIServiceContracts() {
	transactOpts := suite.client.GetTransactOpts()

	// deploy iservice market extension contract
	serviceMarketAddr, tx, _, err := iservice.DeployIServiceMarketEx(transactOpts, suite.client)
	suite.NoError(err)

	fmt.Printf("iservice market extension deployed, contract address: %s, tx hash: %s", serviceMarketAddr.String(), tx.Hash().String())

	// deploy iservice core extension contract
	serviceCoreAddr, tx, _, err := iservice.DeployIServiceCoreEx(transactOpts, suite.client, "fisco-1-1", serviceMarketAddr, transactOpts.From)
	suite.NoError(err)

	fmt.Printf("iservice core extension deployed, contract address: %s, tx hash: %s", serviceCoreAddr.String(), tx.Hash().String())
}
