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

    modifier isDone() {
        require(completed != Status.Pending, "Job is not completed");
        _;
    }

    modifier isPending() {
        require(completed == Status.Pending, "Job is not pending");
        _;
    }

    enum Status {
        Pending,
        Completed,
        Rejected
    }

    /**
     * @dev     The status of the job execution.
     */
    Status private completed = Status.Pending;

    /**
     * @dev     The result (jobId) of the job execution.
     */
    string private result;

    /**
     * @dev     The amount paid by the used for the job execution.
     */
    uint256 private amountPaid;

    event JobSucceeded();
    event JobFailed();

    constructor(address requestInitiator, address sapoBridge) payable {
        initiator = requestInitiator;
        bridge = sapoBridge;
    }

    /**
     * @dev     Set the result of the job execution.
     *          Only the bridge can call this function.
     * @param   executionResult The result of the job execution.
     */
    function saveResult(string memory executionResult) public isBridge isPending {
        (bool success, ) = payable(initiator).call{value: address(this).balance}("");
        require(success, "Failed to send Ether");
        completed = Status.Completed;
        result = executionResult;
        emit JobSucceeded();
    }

    /**
     * @dev     Get the result of the job execution.
     *          Only the initiator can call this function.
     * @return  The result of the job execution.
     */
    function getResult() public view isInitiator isDone returns (string memory) {
        return result;
    }

    /**
     * @dev     Refund the initiator of the job execution.
     *          Only the bridge can call this function.
     */
    function failAndRefund(string memory reason) public isBridge isPending {
        (bool success, ) = payable(initiator).call{value: address(this).balance}("");
        require(success, "Failed to send Ether");
        completed = Status.Rejected;
        result = reason;
        emit JobFailed();
    }

    /* Getters */
    function getBridgeAddress() public view returns (address) {
        return bridge;
    }

    function getInitiator() public view returns (address) {
        return initiator;
    }

    function getStatus() public view returns (Status) {
        return completed;
    }
}
