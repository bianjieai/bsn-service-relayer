var Web3 = require('web3');
var Tx = require('ethereumjs-tx');
var solc = require('solc');
var fs = require('fs');
var sleep = require('sleep');

// web3 instance
var web3 = new Web3(new Web3.providers.HttpProvider("https://ropsten.infura.io"));

// ethereum config
var fromAddress = "0xe61ec6bbdedd0ab319cf311152fca5f487153481";
var privateKey = "";
var nonce = web3.eth.getTransactionCount(fromAddress);
var gasPrice = web3.eth.gasPrice;
var gasLimit = 500000;

// compile contract source code
var source = fs.readFileSync("NFTService.sol");
var compilied = solc.compile(source, 1);
var bytecode = compilied.contracts[':NFTService'].bytecode;

// build raw transaction
var rawTx = {
    'from': fromAddress,
    'nonce': web3.toHex(count),
    'gasPrice': web3.toHex(gasPrice),
    'gasLimit': web3.toHex(gasLimit),
    'value': '0x0',
    'data': '0x'+bytecode
};

// sign raw transaction
var privKey = new Buffer.from(privateKey, 'hex');
var tx = new Tx(rawTx);
tx.sign(privateKey);
var serializedTx = tx.serialize();

// build raw tx
var rawTransaction = {
    "from": fromAddress,
    "nonce": web3.toHex(nonce),
    "gasPrice": web3.toHex(gasPrice),
    "gasLimit": web3.toHex(gasLimit),
    "to": contractAddress,
    "value": "0x0",
    "data": data,
};

// sign tx
var privKey = new Buffer.from(privateKey, 'hex');
var tx = new Tx(rawTransaction);
tx.sign(privKey);
var serializedTx = tx.serialize();

// initiate nft minting transaction
web3.eth.sendRawTransaction('0x' + serializedTx.toString('hex'), 
    function(err, txHash) {
        if (!err) {
            console.log("tx hash: " + txHash);
            watch(txHash);
        } else {
            console.log(err);
        }
});

// watch contract deployment
function watch(txHash) {
    var tx;

    while (true) {
        tx = web3.eth.getTransaction(txHash);
        
        if (tx["blockNumber"] !== null) {
            var receipt = web3.eth.getTransactionReceipt(txHash);
            console.log("contract address: " + receipt["contractAddress"]);
            
            break;
        }

        sleep(1);
    }
}
