// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

/**
 * @title   Sapo Job
 * @dev     A bacalhau job contract.
 *          The initiator of the job is the owner of the contract.
 *          The bridge is the only one who can set the result.
 *          The initiator is the only one who can get the result.
 */
contract SapoJob {
    /**
     * @dev     The address of the initiator of the job.
     *          The initiator is the owner of the contract.
     *          He's the one who requested this job execution
     */
    address private initiator;

    modifier isInitiator() {
        require(msg.sender == initiator, "Caller is not initiator");
        _;
    }

    /**
     * @dev     The address of the bridge.
     *          The bridge is the only one who can set the result.
     */
    address private bridge;

    modifier isBridge() {
        require(msg.sender == bridge, "Caller is not bridge");
        _;
    }

    /**
     * @dev     The status of the job execution.
     */
    bool public completed = false;

    /**
     * @dev     The result of the job execution.
     */
    string private result;

    /**
     * @dev     The amount paid by the used for the job execution.
     */
    uint256 private amountPaid;

    event JobSucceeded();
    event JobFailed();

    constructor(address sapoBridge) payable {
        initiator = msg.sender;
        bridge = sapoBridge;
    }

    /**
     * @dev     Set the result of the job execution.
     *          Only the bridge can call this function.
     * @param   executionResult The result of the job execution.
     */
    function saveResult(string memory executionResult) public isBridge {
        require(!completed, "Job is already finished");
        completed = true;
        result = executionResult;
        emit JobSucceeded();
    }

    /**
     * @dev     Get the result of the job execution.
     *          Only the initiator can call this function.
     * @return  The result of the job execution.
     */
    function getResult() public view isInitiator returns (string memory) {
        require(completed, "Job is not finished");
        return result;
    }

    /**
     * @dev     Refund the initiator of the job execution.
     *          Only the bridge can call this function.
     */
    function failAndRefund() public isBridge {
        require(!completed, "Job is already finished");
        bool success = payable(initiator).send(amountPaid);
        require(success, "Failed to send Ether");
        emit JobFailed();
    }

    function fund() public payable isBridge {
        require(msg.value > 0, "You need to send a positive value");
    }

    /* Getters */

    function getBridgeAddress() public view returns (address) {
        return bridge;
    }

    function getInitiator() public view returns (address) {
        return initiator;
    }

    function getState() public view returns (bool) {
        return completed;
    }
}
