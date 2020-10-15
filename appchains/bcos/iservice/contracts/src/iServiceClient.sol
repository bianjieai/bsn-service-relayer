pragma solidity ^0.6.0;

import "../interfaces/iServiceInterface.sol";

/*
 * @title Contract for the iService core extension client
 */
contract iServiceClient {
    iServiceInterface iServiceCore; // iService Core contract address
    
    // mapping the request id to RequestStatus
    mapping(bytes32 => RequestStatus) requests;
    
    // request status
    struct RequestStatus {
        bool sent; // request sent
        bool responded; // request responded
    }
    
    /*
     * @title Event triggered when the iService request is sent
     * @param _requestID Request id
     */
    event IServiceRequestSent(bytes32 _requestID);
    
    /*
     * @title Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "iServiceClient: request does not exist");
        require(!requests[_requestID].responded, "iServiceClient: request has been responded");
        
        _;
    }
    
    /*
     * @title Send iService request
     * @param _serviceName Service name
     * @param _input Service request input
     * @param _timeout Service request timeout
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendIServiceRequest(
        string calldata _serviceName,
        string calldata _input,
        uint256 _timeout,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
        internal
        returns (bytes32 requestID)
    {
        requestID = iServiceCore.callService(serviceName, input, timeout, _callbackAddress, _callbackFunction);
        
        emit IServiceRequestSent(requestID);
        
        requests[requestID].sent = true;
        
        return requestID
    }

    /**
     * @title Set the iService core contract address
     * @param _iServiceCore The address of the iService core contract
     */
    function setIServiceCore(address _iServiceCore) internal {
        iServiceCore = _iServiceCore;
    }
}