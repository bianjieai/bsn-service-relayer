pragma solidity ^0.4.24;

/**
 * @title core interface
 */
interface coreInterface {
    /**
     * @dev Send cross chain request
     * @param _destChainID Target chain address
     * @param _endpointAddress Target contract address or service bind address
     * @param _endpointType service/contract/offchain
     * @param _methodAndArgs Target method name and arguments
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendRequest(
        string _destChainID,
        string _endpointAddress,
        string _endpointType,
        string _methodAndArgs,
        address _callbackAddress,
        bytes4 _callbackFunction
    ) external returns (bytes32 requestID);

    /**
     * @dev Set the response of the specified service request
     * @param _requestID Request id
     * @param _errMsg Error message of the service invocation
     * @param _output Response output
     */
    function setResponse(
        bytes32 _requestID,
        string  _errMsg,
        string  _output
    ) external returns (bool);
}

/*
 * @title Contract for the iService core extension client
 */
contract iServiceClient {
    coreInterface iServiceCore; // iService Core contract address

    // mapping the request id to Request
    mapping(bytes32 => Request) requests;

    // request
    struct Request {
        address callbackAddress; // callback contract address
        bytes4 callbackFunction; // callback function selector
        bool sent; // request sent
        bool responded; // request responded
    }

    /*
     * @dev Event triggered when the iService request is sent
     * @param _requestID Request id
     */
    event RequestSent(bytes32 _requestID);

    /*
     * @dev Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "iServiceClient: request does not exist");
        require(!requests[_requestID].responded, "iServiceClient: request has been responded");

        _;
    }

    /*
     * @dev Send the iService request
     * @param _serviceName Service name
     * @param _input Service request input
     * @param _timeout Service request timeout
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendIServiceRequest(
        string _destChainID,
        string _endpointAddress,
        string _endpointType,
        string _methodAndArgs,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
    internal
    returns (bytes32 requestID)
    {
        requestID = iServiceCore.sendRequest(_destChainID, _endpointAddress, _endpointType, _methodAndArgs, address(this), this.onResponse.selector);

        Request memory request = Request(
            _callbackAddress,
            _callbackFunction,
            true,
            false
        );

        requests[requestID] = request;

        emit RequestSent(requestID);

        return requestID;
    }

    /*
     * @dev Callback function
     * @param _requestID Request id
     * @param _output Response output
     */
    function onResponse(
        bytes32 _requestID,
        string _output
    )
    external
    validRequest(_requestID)
    {
        requests[_requestID].responded = true;

        address cbAddr = requests[_requestID].callbackAddress;
        bytes4 cbFunc = requests[_requestID].callbackFunction;

        cbAddr.call(abi.encodeWithSelector(cbFunc, _requestID, _output));
    }

    /**
     * @dev Set the iService core contract address
     * @param _iServiceCore Address of the iService core contract
     */
    function setIServiceCore(address _iServiceCore) internal {
        require(_iServiceCore != address(0), "iServiceClient: iService core address can not be zero");
        iServiceCore = coreInterface(_iServiceCore);
    }
}

/*
 * @title Contract for inter-chain NFT minting powered by iService
 * The service is supported by price service
 */
contract ServiceConsumer is iServiceClient {
    string private destChainID = "fisco-1-1";
    string private endpointAddress = "0xe9708c47B560AC923E5a9096669fC71E8bD771Cb";
    string private endpointType = "contract";
    string public helloMsg;

    event Hello(bytes32 _requestID, string _helloMsg);
    /*
     * @notice Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _defaultTimeout Default service request timeout
     */
    constructor(
        address _iServiceCore
    )
    public
    {
        setIServiceCore(_iServiceCore);
    }

    /*
     * @notice Start workflow for minting nft
     * @param _args String arguments for minting nft
     */
    function helloWorld(
        string _hello
    )
    external
    {
        sendIServiceRequest(
            destChainID,
            endpointAddress,
            endpointType,
            _hello,
            address(this),
            this.onHello.selector
        );
    }

    /*
     * @notice NFT service callback function
     * @param _requestID Request id
     * @param _output NFT service response output
     */
    function onHello(
        bytes32 _requestID,
        string _output
    )
    external
    validRequest(_requestID)
    {
        emit Hello(_requestID, _output);
    }
}


