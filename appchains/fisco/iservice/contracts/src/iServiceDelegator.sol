pragma solidity ^0.6.10;

/**
 * @title iServiceDelegator is intended to be a proxy to the underlying iService Core contract
 */
contract iServiceDelegator {
    address public iServiceCore; // the underlying contract address
    address public owner; // owner
    
    /**
     * @dev Assert caller has access to a function of the modifier
     */
    modifier hasPermission() {
        require(msg.sender == owner);
        _;
    }

    /**
     * @dev Check if the given address is valid
     */
    modifier valid(address addr) {
        require(addr != address(0), "iServiceDelegator: address must not be zero");
        
        _;
    }

    /**
     * @dev Constructor
     */
    constructor() public {
        owner = msg.sender;
    }

    /**
     * @dev Forward call to the underlying iService Core contract
     */
    function() public {
        assembly {
            let dataOffset := mload(0x40) // allocate free memory for input data
            let dataLen := calldatasize // get the data length
            calldatacopy(dataOffset, 0, dataLen) // copy data to allocated memory above
            mstore(0x40, add(dataOffset, dataLen)) // update free memory pointer

            // call
            let success := call(sub(gas, 10000), sload(iServiceCore_slot), callvalue, dataOffset, dataLen, 0, 0)
            
            // handle result
            switch success
            case 1 {
                // call succeeded and return result data
                
                let resultLen := returndatasize // get result data length
                let resultOffset := mload(0x40) // allocate free memory for result data
                returndatacopy(resultOffset, 0, resultLen) // copy result data to memory

                return(resultOffset, resultLen) // return result data
            }
            case 0 {
                // call failed and revert
                revert(0, 0)
            }
        }
    }

    /**
     * @dev set the iService Core address
     * @param _iServiceCore iService Core contract address
     */
    function setIServiceCore(address _iServiceCore) public valid(_iServiceCore) hasPermission {
        iServiceCore = _iServiceCore;
    }
}