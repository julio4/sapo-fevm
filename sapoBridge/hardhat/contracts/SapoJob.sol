// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "./SapoBridge.sol";

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
     * @dev     The address of wallet of the bridge.
     */
    address private bridge;

    /**
     * @dev     The address of the bridge contract.
     */
    address private owner;

    modifier isBridge() {
        require(msg.sender == bridge || msg.sender == owner, "Caller is not bridge");
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
     * @dev     The result (jobId) of the job execution. Has two parts
     */
    bytes32 private result1;
    bytes32 private result2;

    /**
     * @dev     The result (jobId) of the job execution. Has two parts
     */
    bytes32 private result1CID;
    bytes32 private result2CID;

    /**
     * @dev     The amount paid by the used for the job execution.
     */
    uint256 private amountPaid;

    event JobSucceeded();
    event JobFailed();

    constructor(address requestInitiator, address sapoBridge) payable {
        initiator = requestInitiator;
        bridge = sapoBridge;
        owner = msg.sender;
    }

    /**
     * @dev     Set the result of the job execution.
     *          Only the bridge can call this function.
     * @param   exResult1 result jobId (part 1).
     * @param   exResult2 result jobId (part 2).
     */
    function saveResult(bytes32 exResult1, bytes32 exResult2, bytes32 exCid1, bytes32 exCid2) public isBridge isPending {
        (bool success, ) = payable(bridge).call{value: address(this).balance}("");
        require(success, "Failed to send Ether");
        completed = Status.Completed;
        result1 = exResult1;
        result2 = exResult2;
        result1CID = exCid1;
        result2CID = exCid2;
        emit JobSucceeded();
    }

    /**
     * @dev     Refund the initiator of the job execution.
     *          Only the bridge can call this function.
     */
    function failAndRefund() public isBridge isPending {
        (bool success, ) = payable(initiator).call{value: 0.05 ether}("");
        require(success, "Failed to send Ether to bridge");
        (success, ) = payable(initiator).call{value: address(this).balance}("");
        require(success, "Failed to send Ether to initiator");
        completed = Status.Rejected;
        emit JobFailed();
    }

    /* Getters */
    function getResult() public view isDone returns (string memory) {
        SapoBridge sb = SapoBridge(owner);
        string memory resString = sb.bytes64ToString(result1, result2);
        return resString;
    }

    function getResultCid() public view isDone returns (string memory) {
        SapoBridge sb = SapoBridge(owner);
        string memory resString = sb.bytes64ToString(result1CID, result2CID);
        return resString;
    }

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
