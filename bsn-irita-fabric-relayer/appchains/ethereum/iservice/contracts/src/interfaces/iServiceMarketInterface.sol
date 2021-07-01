pragma solidity ^0.6.10;

/**
 * @title iService market interface for query the service binding by iService Core
 */
interface iServiceMarketInterface {
    /**
     * @dev Check if the given service binding exists
     * @param _serviceName Service name
     * @return exist Indicates if the service binding exists
     */
    function serviceBindingExists(
        string calldata _serviceName
    ) external view returns (bool exist);

    /**
     * @dev Query the service provider of the specified service binding
     * @param _serviceName Service name
     * @return provider Service provider
     */
    function getServiceProvider(
        string calldata _serviceName
    ) external view returns (string memory provider);

    /**
     * @dev Query the service fee of the specified service binding
     * @param _serviceName Service name
     * @return fee Service fee
     */
    function getServiceFee(
        string calldata _serviceName
    ) external view returns (string memory fee);

    /**
     * @dev Query the service quality of the specified service binding
     * @param _serviceName Service name
     * @return qos Service quality
     */
    function getQoS(
        string calldata _serviceName
    ) external view returns (uint256 qos);
}
