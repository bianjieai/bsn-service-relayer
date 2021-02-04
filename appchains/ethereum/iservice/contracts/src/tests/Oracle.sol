pragma solidity ^0.6.10;

import "../interfaces/iServiceInterface.sol";

/*
 * @title Oracle contract powered by iService
 */
contract Oracle {
    string public result; // result
    
    iServiceInterface iServiceContract; // iService contract address 
    
    // oracle request variables
    string serviceName = "oracle-price"; // oracle-specific service name
    string input = "{\"header\":{},\"body\":{\"pair\":\"ETH-USDT\"}}"; // request input
    
    // mapping the request id to RequestStatus
    mapping(bytes32 => RequestStatus) requests;
    
    // request status
    struct RequestStatus {
        bool sent; // request sent
        bool responded; // request responded
    }
    
    /*
     * @notice Event triggered when a request is sent
     * @param _requestID Request id
     */
    event RequestSent(bytes32 _requestID);
    
    /*
     * @notice Event triggered when a request is responded
     * @param _requestID Request id
     * @param _result Result
     */
    event RequestResponded(bytes32 _requestID, string _result);
    
    /*
     * @notice Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _serviceName Service name
     * @param _input Service request input
     */
    constructor(
        address _iServiceContract,
        string memory _serviceName,
        string memory _input,
    )
        public
    {
        iServiceContract = iServiceInterface(_iServiceContract);
        
        if (bytes(_serviceName).length > 0) {
            serviceName = _serviceName;
        }
        
        if (bytes(_input).length > 0) {
            input = _input;
        }
    }
    
    /* 
     * @notice Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "Oracle: request does not exist");
        require(!requests[_requestID].responded, "Oracle: request has been responded");
        
        _;
    }
    
    /*
     * @notice Send iService request
     */
    function sendRequest()
        external
    {
        bytes32 requestID = iServiceContract.callService(serviceName, input, 1, address(this), this.onResponse.selector);
        
        emit RequestSent(requestID);
        
        requests[requestID].sent = true;
    }
    
    /* 
     * @notice Callback function
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
        result = _output;
        requests[_requestID].responded = true;
        
        emit RequestResponded(_requestID, result);
    }
}