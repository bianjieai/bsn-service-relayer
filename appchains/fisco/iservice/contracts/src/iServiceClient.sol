pragma solidity ^0.6.10;

import "../interfaces/iServiceInterface.sol";

/*
 * @title Contract for the iService core extension client
 */
contract iServiceClient {
    iServiceInterface iServiceCore; // iService Core contract address

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
     * @dev Make sure that the sender is the contract itself
     * @param _requestID Request id
     */
    modifier onlySelf() {
        require(msg.sender == address(this), "iServiceClient: sender must be the contract itself");
        
        _;
    }

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
        string memory _serviceName,
        string memory _input,
        uint256 _timeout,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
        internal
        returns (bytes32 requestID)
    {
        requestID = iServiceCore.callService(_serviceName, _input, _timeout, address(this), this.onResponse.selector);
        
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
        string calldata _output
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
        iServiceCore = iServiceInterface(_iServiceCore);
    }
}