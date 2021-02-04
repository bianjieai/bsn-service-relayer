pragma solidity ^0.6.10;

import "../interfaces/iServiceInterface.sol";

/**
 * @title Contract for exchange rate consumer powered by iService
 */
contract ExchangeRateConsumer {
    string public rate; // latest exchange rate
    
    iServiceInterface iServiceContract; // iService contract address 
    
    // iService request variables
    string serviceName = "exchange-rate"; // service name
    string input = "{\"header\":{},\"body\":{\"pair\":\"USD-CNY\"}}"; // request input
    uint256 timeout = 50; // request timeout
    
    // mapping the request id to RequestStatus
    mapping(bytes32 => RequestStatus) requests;
    
    // request status
    struct RequestStatus {
        bool sent; // request sent
        bool responded; // request responded
    }
    
    /**
     * @notice Event triggered when a request is sent
     * @param _requestID Request id
     */
    event RequestSent(bytes32 _requestID);
    
    /**
     * @notice Event triggered when a request is responded
     * @param _requestID Request id
     * @param _rate Exchange rate
     */
    event RequestResponded(bytes32 _requestID, string _rate);
    
    /**
     * @notice Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _serviceName Service name
     * @param _input Service request input
     * @param _timeout Service request timeout
     */
    constructor(
        address _iServiceContract,
        string memory _serviceName,
        string memory _input,
        uint256 _timeout
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
        
        if (_timeout > 0) {
            timeout = _timeout;
        }
    }
    
    /**
     * @notice Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "ExchangeRateService: request does not exist");
        require(!requests[_requestID].responded, "ExchangeRateService: request has been responded");
        
        _;
    }
    
    /**
     * @notice Send iService request
     */
    function sendRequest()
        external
    {
        bytes32 requestID = iServiceContract.callService(serviceName, input, timeout, address(this), this.onResponse.selector);
        
        emit RequestSent(requestID);
        
        requests[requestID].sent = true;
    }
    
    /**
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
        rate = _output;
        requests[_requestID].responded = true;

        emit RequestResponded(_requestID, rate);
    }
}