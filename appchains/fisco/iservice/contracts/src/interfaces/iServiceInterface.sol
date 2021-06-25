pragma solidity ^0.4.24;

/**
 * @title iService interface
 */
interface iServiceInterface {
    /**
     * @dev Send cross chain request
     * @param _endpointInfo information of endpoint
     * @param _method Target method name
     * @param _methodAndArgs Target method name and arguments
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendRequest(
        string _endpointInfo,
        string _method,
        bytes _methodAndArgs,
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
