pragma solidity ^0.6.0;

/**
 * @title iService interface
 */
interface iServiceInterface {
    /**
     * @title Initiate a service invocation
     * @param _serviceName Service name
     * @param _input Request input
     * @param _timeout Request timeout
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return Request id
     */
    function callService(
        string calldata _serviceName,
        string calldata _input,
        uint256 _timeout,
        address _callbackAddress,
        bytes4 _callbackFunction
    ) external returns (bytes32 requestID);

    /**
     * @title Set the response of the specified service request
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
    ) external returns (bool);
}

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
     * @return Request id
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

/*
 * @title Contract for inter-chain NFT minting powered by iService
 * The service is supported by price service
 */
contract NFTService is iServiceClient {
    // price service variables
    string private priceServiceName = "oracle_price"; // price service name
    string private priceRequestInput = "{\"pair\":\"eth-usdt\"}"; // price request input

    // nft service variables
    string private nftServiceName = "nft"; // nft service name
    address private to;
    uint256 private amount;
    string private metaID;
    uint256 private setPrice;
    bool private isForSale;
    
    uint256 public price; // price for minting nft 
    string public nftID; // id of the minted nft

    uint256 private defaultTimeout = 100; // maximum number of irita-hub blocks to wait for; default to 100
    
    /*
     * @title Event triggered when the nft is minted
     * @param _requestID Request id
     * @param _nftID NFT id
     */
    event NFTMinted(string _requestID, string _nftID);
    
    /*
     * @title Event triggered when the price is set
     * @param _requestID Request id
     * @param _price Price
     */
    event PriceSet(string _requestID, uint256 _price);

    /*
     * @title Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _defaultTimeout Default service request timeout
     */
    constructor(
        address _iServiceCore,
        uint256 _defaultTimeout
    )
        public
    {
        setIServiceCore(_iServiceCore);
        
        if (_defaultTimeout > 0) {
            defaultTimeout = _defaultTimeout;
        }
    }

    /*
     * @title Start workflow for minting nft
     * @param _to Destination address to mint to
     * @param _amount Amount of NFTs to be minted
     * @param _metaID Meta id
     * @param _setPrice Price
     * @param _isForSale Whether or not for sale
     */
    function mint (
        address _to,
        uint256 _amount,
        string calldata _metaID,
        uint256 _setPrice,
        bool _isForSale,
    )
        external
    {
        to = _to;
        amount = _amount;
        metaID = _metaID;
        setPrice = _setPrice;
        isForSale = _isForSale;

        _requestPrice()
    }

    /*
     * @title Request Eth price for minting NFT 
     */
    function _requestPrice () internal {
        sendIServiceRequest(
            priceServiceName,
            priceRequestInput,
            defaultTimeout,
            this,
            this.onPriceSet.selector
        );
    }

    /*
     * @title Request to mint an NFT 
     */
    function _requestMint () internal {
        bytes memory nftRequestInput = _buildMintRequest(to, amount, metaID, setPrice, isForSale);
        
        sendIServiceRequest(
            nftServiceName,
            nftRequestInput,
            defaultTimeout,
            this,
            this.onNFTMinted.selector
        );
    }

    /* 
     * @title Price service callback function
     * @param _requestID Request id
     * @param _price Price
     */
    function onPriceSet(
        bytes32 _requestID,
        string calldata _output
    )
        external
    {
        uint256 price = _parsePrice(_output);

        emit PriceSet(_requestID, _price);

        _requestMint();
    }

    /* 
     * @title NFT service callback function
     * @param _requestID Request id
     * @param _nftID NFT id
     */
    function onNFTMinted(
        bytes32 _requestID,
        string calldata _nftID
    )
        external
    {
        nftID = _nftID;

        emit NFTMinted(_requestID, nftID);
    }
    
    /*
     * @title Parse the price from output
     * @param _to Destination address to mint to
     * @param _amount Amount of NFTs to be minted
     * @param _metaID Meta id
     * @param _setPrice Price
     * @param _isForSale Whether or not for sale
     */
    function _parsePrice(
        string memory _output
    ) 
        internal
        pure
        returns (uint256)
    {
        return 0;
    }
    
    /*
     * @title Build the nft minting request
     * @param _to Destination address to mint to
     * @param _amount Amount of NFTs to be minted
     * @param _metaID Meta id
     * @param _setPrice Price
     * @param _isForSale Whether or not for sale
     */
    function _buildMintRequest(
        address _to,
        uint256 _amount,
        string calldata _metaID,
        uint256 _setPrice,
        bool _isForSale
    )
        internal
        view
        returns (bytes)
    {
        return abi.encodePacked(_to, amount, _metaID, _setPrice, _isForSale);
    }
}