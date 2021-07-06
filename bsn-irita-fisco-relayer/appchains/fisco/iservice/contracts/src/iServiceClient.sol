pragma solidity ^0.4.24;

/**
 * @title iService interface
 */
interface iServiceInterface {
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
    ) external returns (bytes32 requestID);

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
    ) external returns (bool);
}


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
     * @dev Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "iServiceClient: request does not exist");
        require(!requests[_requestID].responded, "iServiceClient: request has been responded");

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
    function sendIServiceRequest(
        string _endpointInfo,
        string _method,
        bytes _callData,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
    internal
    returns (bytes32 requestID)
    {
        requestID = iServiceCore.sendRequest(_endpointInfo, _method, _callData, address(this), this.onResponse.selector);

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
        string _output
    )
    external
    validRequest(_requestID)
    {
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


