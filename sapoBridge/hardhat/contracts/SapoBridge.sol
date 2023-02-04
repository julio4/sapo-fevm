// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "./SapoJob.sol";

/**
 * @title   Sapo Bridge
 * @dev     Collect bacalhau jobs requests and send them to a bridge.
 */
contract SapoBridge {
    bytes32 constant MAX_BYTES32 = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

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
     * @param   cid1        The cid of the job (part 1).
     * @param   cid2        The cid of the job (part 2).
     */
    event JobExecutionRequest(address sapoJob, bytes32 cid1, bytes32 cid2);

    /**
     * @dev     The user can request a job execution.
     *          The job will be sent to the bridge.
     * @param   cid1    The cid of the job (part 1).
     * @param   cid2    The cid of the job (part 2).
     */
    function request(bytes32 cid1, bytes32 cid2) public payable {
        require(msg.value >= 0.1 ether, "Minimum colateral is 0.1TFIL");
        SapoJob job = new SapoJob(msg.sender, bridge);
        emit JobExecutionRequest(address(job), cid1, cid2);
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
        bytes32 mask = MAX_BYTES32;
        uint shift = 0;

        for (uint j = 0; j < 32; j++) {
            bytes1 char = bytes1((x & mask) << shift);

            if (char != 0) {
                bytesString[charCount] = char;
                charCount++;
            }

            shift += 8;
            mask = mask >> 8;
        }

        mask = MAX_BYTES32;
        shift = 0;

        for (uint j = 0; j < 32; j++) {
            bytes1 char = bytes1((y & mask) << shift);

            if (char != 0) {
                bytesString[charCount] = char;
                charCount++;
            }

            shift += 8;
            mask = mask >> 8;
        }

        bytes memory bytesStringTrimmed = new bytes(charCount);

        for (uint j = 0; j < charCount; j++) {
            bytesStringTrimmed[j] = bytesString[j];
        }

        return string(bytesStringTrimmed);
    }
}
