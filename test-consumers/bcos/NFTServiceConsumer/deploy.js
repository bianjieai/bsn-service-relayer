const Configuration = require('api').Configuration;
const Web3jService = require('api').Web3jService;
const CompileService = require('api').CompileService;
const EventLogService = require('api').EventLogService;
const log4js = require('log4js');

var logger = log4js.getLogger('normal');
logger.level = 'info';

// config
var configPath = './config.json';
var contractPath = './NFTServiceConsumer.sol'; 
var iServiceCoreAddr = '0xD49De1022108932d1f117AE661c18d0299E7e238';

// instantiation
var config = new Configuration(configPath);
var web3jService = new Web3jService(config);
var compileService = new CompileService(config);
var eventLogService = new EventLogService(config);

// build contract instance
var contractClass = compileService.compile(contractPath);
var consumerInstance = contractClass.newInstance();

logger.info('deploying NFT service consumer contract');

// deploy contract
consumerInstance.$deploy(web3jService, iServiceCoreAddr, 0)
.then(contractAddr => {
    logger.info('deployed contract address: %s', contractAddr);
})
.catch(logger.error);