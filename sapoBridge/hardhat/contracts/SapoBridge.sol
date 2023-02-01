// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import './SapoJob.sol';

/**
 * @title   Sapo Bridge
 * @dev     Collect bacalhau jobs requests and send them to a bridge.
 */
contract SapoBridge {
    /**
     * @dev     The owner of the contract.
     *          He can change bridge address.
     */
    address private owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only Owner can call this function.");
        _;
    }

    /**
     * @dev     The address of the bridge.
     *          The bridge is responsible for executing the jobs, sending back
     *          results, and refunding the user if the job fails.
     */
    address private bridge;

    modifier onlyBridge() {
        require(msg.sender == bridge, "Only Bridge can call this function.");
        _;
    }

    /**
     * @dev     Change the bridge address.
     *          Only the owner can call this function.
     * @param   bridgeAddress   The new bridge address.
     */
    function setBridge(address bridgeAddress) public onlyOwner {
        bridge = bridgeAddress;
    }

    constructor(address bridgeAddress) {
        owner = msg.sender;
        bridge = bridgeAddress;
    }
    
    /**
     * @dev     The addresses of the bacalhau jobs owned by each users.
     */
    mapping(address => address[]) private jobs;

    /**
     * @dev     Get the addresses of the bacalhau jobs owned by a user.
     * @param   user    The user address.
     * @return  The addresses of the bacalhau jobs owned by the user.
     */
    function getJobs(address user) public view returns (address[] memory) {
        return jobs[user];
    }

    /**
     * @dev     The event emitted when a job execution is requested.
     *          The bridge is listening for these events.
     * @param   sapoJob     The address of the job as an identifier.
     * @param   cid         The cid of the job specification.
     */
    event JobExecutionRequest(
        address sapoJob,
        string cid
    );

    /**
     * @dev     The user can request a job execution.
     *          The job will be sent to the bridge.
     * @param   cid     The cid of the job.
     */
    function request(string memory cid) public {
        SapoJob job = new SapoJob(bridge);
        emit JobExecutionRequest(address(job), cid);
        jobs[msg.sender].push(address(job));
    }

    /**
     * @dev     The bridge can send the result of a job execution.
     *          The result will be saved on the job.
     * @param   job     The job address.
     * @param   result  The job result.
     */
    function saveResult(address job, string memory result) public onlyOwner {
        SapoJob(job).saveResult(result);
    }

    /**
     * @dev     The bridge can fail a job execution.
     *          The job will be refunded.
     * @param   job     The job address.
     */
    function failAndRefund(address job) public onlyOwner {
        SapoJob(job).failAndRefund();
    }
}