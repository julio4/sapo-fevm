// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import './SapoJob.sol';

/**
 * @title Sapo Bridge
 * @dev Collect jobs requests to dispatch on bacalhau
 *      Perform simple verification
 */
contract SapoBridge {
    address private owner;
    address private bridge;
    mapping(address => address[]) public jobs;

    event JobExecutionRequest(
        address sapoJob,
        string cid
    );

    modifier onlyOwner() {
        require(msg.sender == owner, "Only bridge can call this function.");
        _;
    }

    constructor(address bridgeAddress) {
        owner = msg.sender;
        bridge = bridgeAddress;
    }

    function request(string memory cid) public {
        SapoJob job = new SapoJob(bridge);
        emit JobExecutionRequest(address(job), cid);
        jobs[msg.sender].push(address(job));
    }

    function saveResult(address job, string memory result) public onlyOwner {
        SapoJob(job).saveResult(result);
    }

    function failAndRefund(address job) public onlyOwner {
        SapoJob(job).failAndRefund();
    }

    function getJobs(address user) public view returns (address[] memory) {
        return jobs[user];
    }
}