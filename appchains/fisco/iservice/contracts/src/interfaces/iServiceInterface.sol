pragma solidity ^0.4.24;

/**
 * @title iService interface
 */
interface iServiceInterface {
    /**
     * @dev Send cross chain request
     * @param _destChainID Target chain address
     * @param _endpointAddress Target contract address or service bind address
     * @param _endpointType service/contract/offchain
     * @param _methodAndArgs Target method name and json string of arguments
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
