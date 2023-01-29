// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0;

/**
 * @title Sapo Job
 * @dev Representation of a Sapo job execution
 *      Owner is bridge
 */
contract SapoJob {
    address private bridge;
    address private initiator;
    string private jobSpec;

    bool public completed = false;
    string private result;

    modifier isBridge() {
        require(msg.sender == bridge, "Caller is not bridge");
        _;
    }

    modifier isInitiator() {
        require(msg.sender == initiator, "Caller is not initiator");
        _;
    }

    constructor(address sapoBridge) {
        initiator = msg.sender;
        bridge = sapoBridge;
    }

    function complete(string memory executionResult) public isBridge {
        require(!completed);
        completed = true;
        result = executionResult;
    }

    function getResult() public view isInitiator returns (string memory) {
        require(completed);
        return result;
    }
}
