var Web3 = require('web3');
var Tx = require('ethereumjs-tx').Transaction;
var fs = require('fs')

// ethereum config
var chainID = 'ropsten';
var nodeRPCAddr = 'wss://ropsten.infura.io/ws/v3/56e89587eacb4fbe8655e4c44b146237'
var fromAddress = '0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b';
var privateKey = '5dee232c8be5cb81f0ae6fddb45243fc6208192c16aef275ef41b019df765d1f';
var gasPrice = 20000000000;
var gasLimit = 500000;

// contract config
var consumerContractAddr = '0x248d41Aab5c95f02912B63dE2047305B65d34606';
var abiPath = './artifacts/NFTServiceConsumer.json';

// web3 instance
var web3 = new Web3(nodeRPCAddr);

// abi
var abi = JSON.parse(fs.readFileSync(abiPath));

// contract address passed in
var args = process.argv.splice(2);
if (args.length > 0) {
    consumerContractAddr = args[0];
}

// consumer contract instance
var consumerContract = new web3.eth.Contract(abi, consumerContractAddr);

// nft variables
var destAddress = '0xaa27bb5ef6e54a9019be7ade0d0fc514abb4d03b';
var amount = 1;
var metaID = 'test-id';
var setPrice = 1; // price in USDT
var isForSale = true;

// transaction data for minting nft 
var data = consumerContract.methods.mint(
    destAddress,
    amount,
    metaID,
    setPrice,
    isForSale
).encodeABI();

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
            to: consumerContractAddr,
            value: '0x0',
            data: data,
        };
        
        // sign transaction
        var privKey = new Buffer.from(privateKey, 'hex');
        var tx = new Tx(rawTx, {chain: chainID});
        tx.sign(privKey);
        var serializedTx = tx.serialize();
        
        // initiate nft minting transaction
        web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'))
        .on('transactionHash', function(hash){
            console.log('tx hash: %s', hash);
        })
        .on('error', console.error);
})
.catch(console.error);

// listen to events
consumerContract.events.allEvents()
.on('data', function(event){
    switch (event.event) {
        case 'IServiceRequestSent':
            console.log('request sent: %s', event.returnValues._requestID);
            break;
        
        case 'PriceSet':
            console.log('price got: %s', event.returnValues._price);
            break;
        
        case 'NFTMinted':
            console.log('nft minted: %s', event.returnValues._nftID);
            break;
        
        default:
            console.log('event triggered: %s', event.event);
    }
})
.on('error', console.error);
