pragma solidity ^0.4.24;

import "./vendor/Ownable.sol";
import "./interfaces/iServiceInterface.sol";

/**
 * @title iService Core Extension contract
 */
contract iServiceCoreEx is iServiceInterface, Ownable {
    string private sourseChainID = "fisco-1-1";
    // mapping the request id to Request
    mapping(bytes32 => Request) requests;

    // mapping the request id to Callback
    mapping(bytes32 => Callback) callbacks;

    // global request count
    uint256 public requestCount;

    // address allowed to relay the interchain requests
    address public relayer;

    // service request
    struct Request {
        bytes32 id; // request id
        bool responded; // if request has been responded
    }

    // request callback
    struct Callback {
        address callbackAddress; // callback contract address
        bytes4 functionSelector; // callback function selector
    }
    // service response
    struct Response {
        bytes32 id; // response id, equals to request id
        string errMsg; // error message of the service invocation
        string output; // response output
    }


    /**
     * @dev Event triggered when the request is sent
     * @param _requestID Request id
     * @param _endpointInfo information of endpoint
     * @param _method Target method name
     * @param _methodAndArgs abi decode of target method name and arguments
     * @param _sender Message sender
     */
    event CrossChainRequestSent(
        bytes32 _requestID,
        string _endpointInfo,
        string _method,
        bytes _methodAndArgs, // target method name and json string of arguments
        address _sender
    );

    /**
     * @dev Constructor
     * @param _relayer Relayer address
     */
    constructor(address _relayer) public Ownable() {
        if (_relayer != address(0)) {
            relayer = _relayer;
        } else {
            relayer = owner();
        }
    }

    /**
     * @dev Make sure that the request is valid
     */
    modifier checkRequest(
        string _endpointInfo,
        string _method,
        bytes  _methodAndArgs
    ) {
        require(
            bytes(_endpointInfo).length > 0,
            "iServiceCoreEx: destChainID can not be empty"
        );
        require(
            bytes(_method).length > 0,
            "iServiceCoreEx: method can not be empty"
        );
        require(
            _methodAndArgs.length > 0,
            "iServiceCoreEx: methodAndArgs can not be empty"
        );
        _;
    }

    /**
     * @dev Make sure that the request exists and has not been responded
     * @param _requestID Request id
     */
    modifier validateRequest(bytes32 _requestID) {
        require(
            requests[_requestID].id.length > 0,
            "iServiceCoreEx: request does not exist"
        );
        require(
            requests[_requestID].responded == false,
            "iServiceCoreEx: request has been responded"
        );

        _;
    }


    /**
     * @dev Make sure that the sender is the relayer
     */
    modifier onlyRelayer() {
        require(
            msg.sender == relayer,
            "iServiceCoreEx: sender is not the relayer"
        );
        _;
    }

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
    )
    external
    checkRequest(_endpointInfo, _method, _methodAndArgs)
    returns (bytes32 requestID)
    {
        requestID = keccak256(abi.encodePacked(sourseChainID, requestCount));

        requestCount ++;

        emit CrossChainRequestSent(
            requestID,
            _endpointInfo,
            _method,
            _methodAndArgs,
            msg.sender
        );

        _saveRequestCallback(requestID, _callbackAddress, _callbackFunction);
        return requestID;
    }

    /**
     * @dev Set the response of the specified service request
     * @param _requestID Request id
     * @param _errMsg Error message of the service invocation
     * @param _output Response output
     * @return True on success, false otherwise
     */
    function setResponse(
        bytes32 _requestID,
        string _errMsg,
        string _output
    ) external onlyRelayer validateRequest(_requestID) returns (bool) {
        Callback memory cb = callbacks[_requestID];

        string memory result;

        if(bytes(_errMsg).length > 0){
            result = _errMsg;
        }else{
            result = _output;
        }

        requests[_requestID].responded = true;
        bool success =
        cb.callbackAddress.call(
            abi.encodeWithSelector(
                cb.functionSelector,
                _requestID,
                result
            )
        );

        return success;
    }

    /**
     * @notice Set the relayer address
     * @param _address Relayer address
     */
    function setRelayer(address _address) external onlyOwner {
        require(
            _address != address(0),
            "iServiceCoreEx: relayer address can not be zero"
        );
        relayer = _address;
    }

    function _saveRequest(
        bytes32 _requestID
    ) internal {
        Request memory req;

        req.id = _requestID;
        req.responded = false;

        requests[_requestID] = req;
    }

    /**
     * @notice Save the request callback
     * @param _requestID Request id
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     */
    function _saveRequestCallback(
        bytes32 _requestID,
        address _callbackAddress,
        bytes4 _callbackFunction
    ) internal {
        Callback memory cb;

        cb.callbackAddress = _callbackAddress;
        cb.functionSelector = _callbackFunction;

        callbacks[_requestID] = cb;
    }
}
