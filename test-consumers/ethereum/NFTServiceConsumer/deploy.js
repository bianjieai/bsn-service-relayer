var Web3 = require('web3');
var Tx = require('ethereumjs-tx').Transaction;
var solc = require('solc');
var fs = require('fs');
var log4js = require('log4js');

var logger = log4js.getLogger('normal');
logger.level = 'info';

// ethereum config
var chainID = 'ropsten';
var nodeRPCAddr = 'wss://ropsten.infura.io/ws/v3/56e89587eacb4fbe8655e4c44b146237'
var fromAddress = '0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b';
var privateKey = '5dee232c8be5cb81f0ae6fddb45243fc6208192c16aef275ef41b019df765d1f';
var gasPrice = 20000000000;
var gasLimit = 5000000;

// contract config
var iserviceCoreAddr = '0x79b6c1ab5dbeba4879bfbea35a78fac8e6c73c92';
var abiPath = './artifacts/NFTServiceConsumer.json';
var bytecodePath = './artifacts/NFTServiceConsumer.bytecode';

// abi and bytecode
var abi = JSON.parse(fs.readFileSync(abiPath));
var bytecode = fs.readFileSync(bytecodePath);

// web3 instance
var web3 = new Web3(nodeRPCAddr);

// consumer contract instance
var consumerContract = new web3.eth.Contract(abi);

// contract deployment data
var data = consumerContract.deploy({
    data: '0x'+bytecode,
    arguments: [iserviceCoreAddr, 0]
}).encodeABI();

// build, sign and send transaction
web3.eth.getTransactionCount(fromAddress)
.then(
    nonce => {
        // build raw tx
        var rawTx = {
            from: fromAddress,
            nonce: nonce,
            gasPrice: web3.utils.toHex(gasPrice),
            gasLimit: web3.utils.toHex(gasLimit),
            value: '0x0',
            data: data,
        };

        // sign transaction
        var privKey = new Buffer.from(privateKey, 'hex');
        var tx = new Tx(rawTx, {chain: chainID});
        tx.sign(privKey);
        var serializedTx = tx.serialize();
        
        logger.info('starting to deploy contract');

        // initiate contract deployment transaction
        web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'))
        .on('transactionHash', function(hash){
            logger.info('contract deployment tx sent, tx hash: %s', hash);
        })
        .on('receipt', function(receipt){
            logger.info('contract deployment tx minted, contract address: %s', receipt.contractAddress);
        })
        .on('error', logger.error);
})
.catch(logger.error);
