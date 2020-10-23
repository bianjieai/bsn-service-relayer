pragma solidity ^0.6.10;
pragma experimental "ABIEncoderV2";

import "./vendor/Ownable.sol";
import "./interfaces/iServiceInterface.sol";

/**
 * @title iService Core Extension contract
 */
contract iServiceCoreEx is iServiceInterface, Ownable {
    // mapping the request id to Request
    mapping(bytes32 => Request) requests;

    // mapping the request id to Response
    mapping(bytes32 => Response) responses;

    // global request count
    uint256 public requestCount;

    // address allowed to relay the interchain requests
    address relayer;
    
    // service request
    struct Request {
        string serviceName; // service name
        string input; // request input
        uint256 timeout; // request timeout
        address callbackAddress; // callback contract address
        bytes4 callbackFunction; // callback function selector
    }

    // service response
    struct Response {
        string errMsg; // error message of the service invocation
        string output; // response output
        string icRequestID; // the interchain request id on Irita-Hub
    }

    /**
     * @notice Event triggered when the service invocation is initiated
     * @param _requestID Request id
     * @param _serviceName Service name
     * @param _input Request input
     * @param _timeout Request timeout
     */
    event ServiceInvoked(bytes32 indexed _requestID, string _serviceName, string _input, uint256 _timeout);

    /**
     * @notice Constructor
     * @param _relayer Relayer address
     */
    constructor(address _relayer)
        public
        Ownable()
    {
        if (_relayer != address(0)) {
            relayer = _relayer;
        }
    }

    /**
     * @notice Make sure that the request is valid
     * @param _serviceName Service name
     * @param _input Request input
     * @param _timeout Request timeout
     */
    modifier checkRequest(
        string memory _serviceName,
        string memory _input,
        uint256 _timeout
    )
    {
        require(bytes(_serviceName).length > 0, "iServiceCoreEx: service name can not be empty");
        require(bytes(_input).length > 0, "iServiceCoreEx: request input can not be empty");
        require(_timeout > 0, "iServiceCoreEx: request timeout must be greater than 0");
        
        _;
    }

    /**
     * @notice Make sure that the request exists and has not been responded
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(bytes(requests[_requestID].serviceName).length > 0, "iServiceCoreEx: request does not exist");
        require(bytes(responses[_requestID].icRequestID).length == 0, "iServiceCoreEx: request has been responded");
        
        _;
    }
    
    /** 
     * @notice Make sure that the sender is the relayer
     */
    modifier onlyRelayer() {
        require(msg.sender == relayer, "iServiceCoreEx: sender is not the relayer");
        _;
    }

    /**
     * @notice Initiate a service invocation
     * @param _serviceName Service name
     * @param _input Request input
     * @param _timeout Request timeout
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function callService(
        string calldata _serviceName,
        string calldata _input,
        uint256 _timeout,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
        external
        override
        checkRequest(_serviceName, _input, _timeout)
        returns (bytes32 requestID)
    {
        Request memory req;
        
        req.serviceName = _serviceName;
        req.input = _input;
        req.timeout= _timeout;
        req.callbackAddress = _callbackAddress;
        req.callbackFunction = _callbackFunction;
        
        return _sendRequest(req);
    }

    /**
     * @notice Set the response of the specified service request
     * @param _requestID Request id
     * @param _errMsg Error message of the service invocation
     * @param _output Response output
     * @param _icRequestID Request id on Irita-Hub
     * @return True on success, false otherwise
     */
    function setResponse(
        bytes32 _requestID,
        string calldata _errMsg,
        string calldata _output,
        string calldata _icRequestID
    )
        external
        override
        onlyRelayer
        validRequest(_requestID)
        returns (bool)
    {
        Response memory resp;

        resp.errMsg = _errMsg;
        resp.icRequestID = _icRequestID;

        if (bytes(_errMsg).length == 0) {
            resp.output = _output;
        }
        
        responses[_requestID] = resp;

        Request memory req = requests[_requestID];
        (bool success, ) = req.callbackAddress.call(abi.encodeWithSelector(req.callbackFunction, _requestID, resp.output));
        
        return success;
    }

    /**
     * @notice Retrieve the response of the given service request 
     * @param _requestID Request id
     * @return Response
     */
    function getResponse(
        bytes32 _requestID
    )
        public
        view
        returns (Response memory)
    {
        return responses[_requestID];
    }

    /**
     * @notice Set the relayer address
     * @param _address Relayer address
     */
    function setRelayer(
        address _address
    )
        external
        onlyOwner
    {
        require(_address != address(0), "iServiceCoreEx: relayer address can not be zero");
        relayer = _address;
    }
    
    /**
     * @notice Send the service request
     * @param _req Request
     */
    function _sendRequest(
        Request memory _req
    )
        internal
        returns (bytes32 requestID)
    {
        requestID = keccak256(abi.encodePacked(this, requestCount));
        
        requests[requestID] = _req;
        requestCount += 1;
        
        emit ServiceInvoked(requestID, _req.serviceName, _req.input, _req.timeout);
        
        return requestID;
    }
}
