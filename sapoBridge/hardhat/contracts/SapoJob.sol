// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

/**
 * @title Sapo Job
 * @dev Representation of a Sapo job execution
 *      Owner is bridge
 */
contract SapoJob {
    address payable private initiator;
    address private bridge;

    string private image;
    string private cid;

    bool public completed = false;
    string private result;

    event JobSuceeded();
    event JobFailed();

    modifier isBridge() {
        require(msg.sender == bridge, "Caller is not bridge");
        _;
    }

    modifier isInitiator() {
        require(msg.sender == initiator, "Caller is not initiator");
        _;
    }

    constructor(address sapoBridge, string memory dockerImg, string memory inputCid) payable {
        initiator = payable(msg.sender);
        bridge = sapoBridge;
        image = dockerImg;
        cid = inputCid;
    }

    function saveResult(string memory executionResult) public isBridge {
        require(!completed);
        completed = true;
        result = executionResult;
        emit JobSuceeded();
    }

    function getResult() public view isInitiator returns (string memory) {
        require(completed);
        return result;
    }

    function failAndRefund() public isBridge {
        require(!completed);
        (bool success, ) = initiator.call{value: address(this).balance}("");
        require(success, "Failed to send Ether");
        emit JobFailed();
    }
}
