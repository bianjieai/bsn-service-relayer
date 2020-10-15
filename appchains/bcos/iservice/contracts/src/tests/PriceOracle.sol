pragma solidity ^0.6.0;

import "../interfaces/iServiceInterface.sol";

/*
 * @title Contract for price oracle powered by iService
 */
contract PriceOracle {
    string public price; // latest price
    
    iServiceInterface iServiceContract; // iService contract address 
    
    // oracle request variables
    string serviceName = "oracle"; // oracle-specific service name
    string input = "btc-usdt"; // feed name
    
    // mapping the request id to RequestStatus
    mapping(bytes32 => RequestStatus) requests;
    
    // request status
    struct RequestStatus {
        bool sent; // request sent
        bool responded; // request responded
    }
    
    /*
     * @title Event triggered when a request is sent
     * @param _requestID Request id
     */
    event RequestSent(bytes32 _requestID);
    
    /*
     * @title Event triggered when a request is responded
     * @param _requestID Request id
     * @param _price Price
     */
    event RequestResponded(bytes32 _requestID, string _price);
    
    /*
     * @title Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _serviceName Service name
     * @param _input Service request input
     * @param _timeout Service request timeout
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
     * @title Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "PriceOracle: request does not exist");
        require(!requests[_requestID].responded, "PriceOracle: request has been responded");
        
        _;
    }
    
    /*
     * @title Send iService request
     */
    function sendRequest()
        external
    {
        bytes32 requestID = iServiceContract.callService(serviceName, input, 1, address(this), this.onResponse.selector);
        
        emit RequestSent(requestID);
        
        requests[requestID].sent = true;
    }
    
    /* 
     * @title Callback function
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
        price = _output;
        requests[_requestID].responded = true;
        
        emit RequestResponded(_requestID, price);
    }
    
}