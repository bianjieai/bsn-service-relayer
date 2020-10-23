pragma solidity ^0.6.10;
pragma experimental "ABIEncoderV2";

import "./vendor/Ownable.sol";

/**
 * @title iService Market Extension contract
 */
contract iServiceMarketEx is Ownable {
    // service binding list
    ServiceBinding[] bindings;

    // mapping the service name to index
    mapping(string => ServiceBindingIndex) bindingIndices;

    // service binding
    struct ServiceBinding {
        string serviceName; // service name
        string schemas; // service schemas
        string provider; // service provider
        string serviceFee; // service fee
        uint256 qos; // service quality, in terms of the minimum response time
    }

    // service binding index
    struct ServiceBindingIndex {
        uint256 index; // index
        bool exist; // indicates if the service binding exists
    }

    /**
     * @notice Event triggered when a service binding is added
     * @param _serviceName Service name
     * @param _schemas Input and output schemas of the service definition to which the binding is attached
     * @param _provider Provider address of the binding
     * @param _serviceFee Service fee
     * @param _qos Service quality
     */
    event ServiceBindingAdded(
        string indexed _serviceName,
        string _schemas,
        string _provider,
        string _serviceFee,
        uint256 _qos
    );

    /**
     * @notice Event triggered when the service binding is updated
     * @param _serviceName Service name
     * @param _provider Provider address of the binding
     * @param _serviceFee Service fee
     * @param _qos Service quality
     */
    event ServiceBindingUpdated(
        string indexed _serviceName,
        string _provider,
        string _serviceFee,
        uint256 _qos
    );

    /**
     * @notice Event triggered when the service binding is removed
     * @param _serviceName Service name
     */
    event ServiceBindingRemoved(
        string indexed _serviceName
    );

    /**
     * @notice Constructor
     */
    constructor()
        public
        Ownable()
    {
       // no-op
    }

    /**
     * @notice Make sure that the service binding is valid
     * @param _serviceName Service name
     * @param _schemas Input and output schemas of the service definition to which the binding is attached
     * @param _provider Provider address of the binding
     * @param _serviceFee Service fee
     * @param _qos Service quality of the binding
     */
    modifier validBinding(
        string memory _serviceName,
        string memory _schemas,
        string memory _provider,
        string memory _serviceFee,
        uint256 _qos
    )
    {
        require(bytes(_serviceName).length > 0, "iServiceMarketEx: service name can not be empty");
        require(bytes(_schemas).length > 0, "iServiceMarketEx: service schemas can not be empty");
        require(bytes(_provider).length > 0, "iServiceMarketEx: service provider can not be empty");
        require(bytes(_serviceFee).length > 0, "iServiceMarketEx: service fee can not be empty");
        require(_qos > 0, "iServiceMarketEx: qos should be greater than 0");
        
        _;
    }

    /**
     * @notice Make sure that the specified binding does not exist
     * @param _serviceName Service name
     */
    modifier bindingDoesNotExist(string memory _serviceName) {
        require(!bindingIndices[_serviceName].exist, "iServiceMarketEx: service binding already exists");
        _;
    }

    /**
     * @notice Make sure that the specified binding already exists
     * @param _serviceName Service name
     */
    modifier bindingExists(string memory _serviceName) {
        require(bindingIndices[_serviceName].exist, "iServiceMarketEx: service binding does not exist");
        _;
    }

    /**
     * @notice Add a service binding to the iService market
     * @param _serviceName Service name
     * @param _schemas Input and output schemas of the service definition to which the binding is attached
     * @param _provider Provider address of the binding
     * @param _serviceFee Service fee
     * @param _qos Service quality of the binding
     */
    function addServiceBinding(
        string calldata _serviceName,
        string calldata _schemas,
        string calldata _provider,
        string calldata _serviceFee,
        uint256 _qos
    )
        external
        onlyOwner
        validBinding(_serviceName, _schemas, _provider, _serviceFee, _qos)
        bindingDoesNotExist(_serviceName)
    {
        ServiceBinding memory binding;

        binding.serviceName = _serviceName;
        binding.schemas = _schemas;
        binding.provider = _provider;
        binding.serviceFee = _serviceFee;
        binding.qos = _qos;
        
        _addServiceBinding(binding);
    }

    /**
     * @notice Update the specified service binding
     * @param _serviceName Service name
     * @param _provider Provider address of the binding, not updated if empty
     * @param _serviceFee Service fee, not updated if empty
     * @param _qos Service quality of the binding, not updated if set to zero
     */
    function updateServiceBinding(
        string calldata _serviceName,
        string calldata _provider,
        string calldata _serviceFee,
        uint256 _qos
    )
        external
        onlyOwner
        bindingExists(_serviceName)
    {
        ServiceBinding storage binding = bindings[bindingIndices[_serviceName].index];
        
        if (bytes(_provider).length > 0) {
            binding.provider = _provider;
        }
        
        if (bytes(_serviceFee).length > 0) {
            binding.serviceFee = _serviceFee;
        }
        
        if (_qos > 0) {
            binding.qos = _qos;
        }

        emit ServiceBindingUpdated(_serviceName, _provider, _serviceFee, _qos);
    }

    /**
     * @notice Remove the specified service binding
     * @param _serviceName Service name
     */
    function removeServiceBinding(
        string calldata _serviceName
    )
        external
        onlyOwner
        bindingExists(_serviceName)
    {
        delete bindings[bindingIndices[_serviceName].index]; // delete the binding
        delete bindingIndices[_serviceName]; // delete the index

        emit ServiceBindingRemoved(_serviceName);
    }

    /**
     * @notice Retrieve the specified service binding
     * @param _serviceName Service name
     * @return binding Service binding
     */
    function GetServiceBinding(
        string memory _serviceName
    )
        public
        view
        returns (ServiceBinding memory binding)
    {
        ServiceBindingIndex memory sbi = bindingIndices[_serviceName];

        if (sbi.exist) {
            return bindings[sbi.index];
        }

        return binding;
    }

    /**
     * @notice Query the total number of the service bindings
     * @return count Total number of the service bindings
     */
    function GetServiceBindingCount()
        public
        view
        returns (uint256 count)
    {
        return bindings.length;
    }

    /**
     * @notice Retrieve the specified service binding by index
     * @param _index Index of the service binding
     * @return binding Service binding
     */
    function GetServiceBindingByIndex(
        uint256 _index
    )
        public
        view
        returns (ServiceBinding memory binding)
    {
        if (_index < bindings.length) {
            return bindings[_index];
        }

        return binding;
    }
    
    /**
     * @notice Add the service binding internally
     * @param _binding Service binding to be added
     */
    function _addServiceBinding(
        ServiceBinding memory _binding
    )
        internal
    {
        bindingIndices[_binding.serviceName].index = bindings.length;
        bindingIndices[_binding.serviceName].exist = true;
        
        bindings.push(_binding);

        emit ServiceBindingAdded(_binding.serviceName, _binding.schemas, _binding.provider, _binding.serviceFee, _binding.qos);
    }
}