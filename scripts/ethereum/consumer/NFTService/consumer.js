var Web3 = require('web3');
var Tx = require('ethereumjs-tx');
var fs = require('fs')

// web3 instance
var web3 = new Web3(new Web3.providers.HttpProvider("https://ropsten.infura.io"));

// abi
var abi = JSON.parse(fs.readFileSync("./abi/NFTServiceConsumer.abi")).abi;

// deployed NFTService contract address
var nftServiceAddress = '0xbd2c938b9f6bfc1a66368d08cb44dc3eb2ae27be';

// nft service instance
var nftService = web3.eth.contract(abi).at(nftServiceAddress);

// ethereum config
var fromAddress = "0xe61ec6bbdedd0ab319cf311152fca5f487153481";
var privateKey = "5dee232c8be5cb81f0ae6fddb45243fc6208192c16aef275ef41b019df765d1f";
var nonce = web3.eth.getTransactionCount(fromAddress);
var gasPrice = web3.eth.gasPrice;
var gasLimit = 500000;

// nft variables
var destAddress = "";
var amount = 5;
var metaID = "";
var setPrice = 1; // price in USDT
var isForSale = false;

// transaction data for minting nft 
var data = nftService.mint.getData(
    destAddress, 
    amount, 
    metaID, 
    setPrice, 
    isForSale,
    {from: fromAddress}
);

// build raw transaction
var rawTransaction = {
    "from": fromAddress,
    "nonce": web3.toHex(nonce),
    "gasPrice": web3.toHex(gasPrice),
    "gasLimit": web3.toHex(gasLimit),
    "to": contractAddress,
    "value": "0x0",
    "data": data,
};

// sign transaction
var privKey = new Buffer.from(privateKey, 'hex');
var tx = new Tx(rawTransaction);
tx.sign(privKey);
var serializedTx = tx.serialize();

// initiate nft minting transaction
web3.eth.sendRawTransaction('0x' + serializedTx.toString('hex'), 
    function(err, txHash) {
        if (!err)
            console.log(txHash);
        else
            console.log(err);
});