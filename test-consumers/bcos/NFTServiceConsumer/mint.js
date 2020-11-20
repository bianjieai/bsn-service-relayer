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
var iServiceCoreAddr = '0xe9708c47B560AC923E5a9096669fC71E8bD771Cb';

// instantiation
var config = new Configuration(configPath);
var web3jService = new Web3jService(config);
var compileService = new CompileService(config);
var eventLogService = new EventLogService(config);

// build contract instance
var contractClass = compileService.compile(contractPath);
var consumerInstance = contractClass.newInstance();

// event callback function
function eventHandler(status, logs) {
    console.log(status);
    console.log(logs);

    logs.forEach(log => {
        console.log(log);

        switch (log.event) {
            case 'IServiceRequestSent':
                logger.info('iservice request sent, request id: %s', log.returnValues._requestID);
                break;
            
            case 'PriceSet':
                logger.info('usdt-eth price got: %s', log.returnValues._price);
                break;
            
            case 'NFTMinted':
                logger.info('nft minted: %s', log.returnValues._nftID);
                break;
            
            default:
                logger.info('event triggered: %s', log.event);
        }
    });
}

logger.info('deploying NFT service consumer contract');

// deploy contract
consumerInstance.$deploy(web3jService, iServiceCoreAddr, 0)
.then(contractAddr => {
    logger.info('deployed contract address: %s', contractAddr);
    
    // register event log filter
    // eventLogService.registerEventLogFilter({addresses: [contractAddr]}, eventHandler, consumerInstance.abi);

    // nft variables
    var destAddress = '0x791a0073e6dfd9dc5e5061aebc43ab4f7aa4ae8b';
    var amount = 1;
    var metaID = '-Z-2fJxzCoFJ0MOU-zA3-tiIh7dK6FjDruAxgxW6PEs';
    var setPrice = 1; // price in USDT
    var isForSale = true;

    logger.info(
        'starting to mint NFT, to:%s, amount:%d, metaID:%s, setPrice:%s, isForSale:%s',
        destAddress,
        amount,
        metaID,
        setPrice,
        isForSale
    );

    // start to mint NFT
    consumerInstance.mintV2('{"header":{},"body":{"to":"0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b","amount_to_mint":"1","meta_id":"-Z-2fJxzCoFJ0MOU-zA3-tiIh7dK6FjDruAxgxW6PEs"}}')
    .then(res => {
        logger.info('minting transaction succeeded');
    });
})
.catch(console.error);
