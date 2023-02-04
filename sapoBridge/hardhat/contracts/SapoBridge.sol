// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "./SapoJob.sol";

/**
 * @title   Sapo Bridge
 * @dev     Collect bacalhau jobs requests and send them to a bridge.
 */
contract SapoBridge {
    uint constant MAX_UINT = 115792089237316195423570985008687907853269984665640564039457584007913129639935;

    /**
     * @dev     The owner of the contract.
     *          He can change bridge address.
     */
    address private owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only Owner can call this function");
        _;
    }

    /**
     * @dev     The address of the bridge.
     *          The bridge is responsible for executing the jobs, sending back
     *          results, and refunding the user if the job fails.
     */
    address private bridge;

    modifier onlyBridge() {
        require(msg.sender == bridge, "Only Bridge can call this function");
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
     * @dev     The event emitted when a job execution is requested.
     *          The bridge is listening for these events.
     * @param   sapoJob     The address of the job as an identifier.
     * @param   cid         The cid of the job specification.
     */
    event JobExecutionRequest(address sapoJob, string cid);

    /**
     * @dev     The user can request a job execution.
     *          The job will be sent to the bridge.
     * @param   cid     The cid of the job.
     */
    function request(string memory cid) public payable {
        require(msg.value >= 0.1 ether, "Minimum colateral is 0.1TFIL");
        SapoJob job = new SapoJob(msg.sender, bridge);
        emit JobExecutionRequest(address(job), cid);
        jobs[msg.sender].push(address(job));

        payable(bridge).transfer(msg.value);
    }

    /**
     * @dev     The bridge can send the result of a job execution.
     *          The result will be saved on the job.
     * @param   job     The job address.
     * @param   result1 The job result jobId (part1).
     * @param   result2 The job result jobId (part2).
     */
    function saveResult(address job, bytes32 result1, bytes32 result2) public onlyBridge {
        SapoJob(job).saveResult(result1, result2);
    }

    /**
     * @dev     The bridge can fail a job execution.
     *          The job will be refunded.
     * @param   job     The job address.
     */
    function failAndRefund(address job) public onlyBridge {
        SapoJob(job).failAndRefund();
    }

    /* Getters */

    function getBridgeAddress() public view returns (address) {
        return bridge;
    }

    /**
     * @dev     Get the addresses of the bacalhau jobs owned by a user.
     * @param   user    The user address.
     * @return  The addresses of the bacalhau jobs owned by the user.
     */
    function getJobs(address user) public view returns (address[] memory) {
        return jobs[user];
    }

    /**
     * converts two bytes 32 to string.
     * @param x bytes32 to convert to string (part 1).
     * @param y bytes32 to convert to string (part 2).
     */
    function bytes64ToString(bytes32 x, bytes32 y) public pure returns (string memory) {
        bytes memory bytesString = new bytes(64);
        uint charCount = 0;

        for (uint j = 0; j < 32; j++) {
            uint shift = 8 * j;
            uint andOp = shift == 0 ? MAX_UINT : 2 ** (256 - 8 * j) - 1;
            uint timesOp = 2 ** shift;
            bytes1 char = bytes1(bytes32(uint(x & bytes32(andOp)) * timesOp));

            if (char != 0) {
                bytesString[charCount] = char;
                charCount++;
            }
        }

        for (uint j = 0; j < 32; j++) {
            uint shift = 8 * j;
            uint andOp = shift == 0 ? MAX_UINT : 2 ** (256 - 8 * j) - 1;
            uint timesOp = 2 ** shift;
            bytes1 char = bytes1(bytes32(uint(y & bytes32(andOp)) * timesOp));

            if (char != 0) {
                bytesString[charCount] = char;
                charCount++;
            }
        }

        bytes memory bytesStringTrimmed = new bytes(charCount);

        for (uint j = 0; j < charCount; j++) {
            bytesStringTrimmed[j] = bytesString[j];
        }

        return string(bytesStringTrimmed);
    }
}
