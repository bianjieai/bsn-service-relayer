pragma solidity ^0.4.24;

import "./vendor/Ownable.sol";
import "./interfaces/iServiceInterface.sol";

/**
 * @title iService Core Extension contract
 */
contract iServiceCoreEx is iServiceInterface, Ownable {
    string private sourceChainID ;
    // if the request has been response,the request id will mapped to true
    mapping(bytes32 => bool) requests;

    // mapping the request id to Callback
    mapping(bytes32 => Callback) callbacks;

    // if the icRequest has been received in dest chain,the icRequest id will mapped to true
    mapping(bytes32 => bool) icRequests;

    // global request count
    uint256 public requestCount;

    // address allowed to relay the interchain requests
    address public relayer;

    // request callback
    struct Callback {
        address callbackAddress; // callback contract address
        bytes4 functionSelector; // callback function selector
    }

    /**
     * @dev Event triggered when the request is sent
     * @param _requestID Request id
     * @param _endpointInfo information of endpoint
     * @param _method Target method name
     * @param _callData abi decode of target method name and arguments
     * @param _sender Message sender
     */
    event CrossChainRequestSent(
        string _eventName,
        bytes32 _requestID,
        string _endpointInfo,
        string _method,
        bytes _callData, // target method name and json string of arguments
        address _sender
    );

    /**
    * @dev Event triggered when the request is sent
    * @param _icRequestID Request id
    * @param _result result bytes
    */
    event CrossChainResponseSent(
        string _eventName,
        bytes32 _icRequestID,
        bytes _result
    );

    /**
     * @dev Constructor
     * @param _relayer Relayer address
     */
    constructor(address _relayer, string _sourceChainID) public Ownable() {
        sourceChainID = _sourceChainID;
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
        bytes _callData
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
            _callData.length > 0,
            "iServiceCoreEx: callData can not be empty"
        );
        _;
    }

    /**
     * @dev Make sure that the request has not been responded
     * @param _requestID Request id
     */
    modifier validateRequest(bytes32 _requestID) {
        require(
            requests[_requestID] == false,
            "iServiceCoreEx: request has been responded"
        );

        _;
    }

    /**
    * @dev Make sure that the icRequest  has not been responded
    * @param _icRequestID Request id
    */
    modifier validateIcRequest(bytes32 _icRequestID) {
        require(
            icRequests[_icRequestID] == false,
            "iServiceCoreEx: icRequest has been received"
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
     * @param _callData Target method callData
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendRequest(
        string _endpointInfo,
        string _method,
        bytes _callData,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
    external
    checkRequest(_endpointInfo, _method, _callData)
    returns (bytes32 requestID)
    {
        requestID = keccak256(abi.encodePacked(sourceChainID, requestCount));

        requestCount ++;

        emit CrossChainRequestSent(
            "CrossChainRequestSent",
            requestID,
            _endpointInfo,
            _method,
            _callData,
            msg.sender
        );

        _saveRequestCallback(requestID, _callbackAddress, _callbackFunction);
        requests[requestID] = false;
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

        if (bytes(_errMsg).length > 0) {
            result = _errMsg;
        } else {
            result = _output;
        }

        requests[_requestID] = true;
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
     * @dev call service/contract in dest chain
     * @param _icRequestID Request id
     * @param _endpointAddress endpointAddress
     * @param _callData call data from source chain
     */
    function callService(
        bytes32 _icRequestID,
        address _endpointAddress,
        bytes _callData
    ) public validateIcRequest(_icRequestID) {
        uint callDataLength = _callData.length;
        bytes memory result;
        uint success;
        assembly {
        // call
            let d := add(_callData, 32)
            success := call(
                gas(),
                _endpointAddress,
                callvalue,
                d,
                callDataLength,
                0,
                0
            )

        // handle result
            switch success
            case 1 {
                let size := returndatasize
                result := mload(0x40)
                mstore(0x40, add(result, and(add(add(size, 0x20), 0x1f), not(0x1f))))
                mstore(result, size)
                returndatacopy(add(result, 0x20), 0, size)
            }
            case 0 {
            // call failed and revert
                revert(0, 0)
            }
        }
        if (success == 1) {
            emit CrossChainResponseSent(
                "CrossChainResponseSent",
                _icRequestID,
                result
            );
        }
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
