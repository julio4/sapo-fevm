// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0;

import "./SapoJob.sol";

/**
 * @title Sapo Bridge
 * @dev Collect jobs requests to dispatch on bacalhau
 *      Perform simple verification
 */
contract SapoBridge {
    address private owner;
    address private bridge;

    event JobExecutionRequest(address sapoJob);

    constructor(address bridgeAddress) {
        owner = msg.sender;
        bridge = bridgeAddress;
    }

    function request(string jobSpec) public {
        SapoJob job = new SapoJob(bridge, jobSpec);
        emit JobExecutionRequest(address(job));
    }
}
