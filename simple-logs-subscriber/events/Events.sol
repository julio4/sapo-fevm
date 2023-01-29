// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Events
 * @dev Emit events
 */
contract Events {

    address private owner;

    // event for EVM logging
    event NumberEvent(int number);

    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    /**
     * @dev Set contract deployer as owner
     */
    constructor() {
        owner = msg.sender;
        emit NumberEvent(0);
    }

    /**
     * @dev Emit Number event
     * @param nb the nb
     */
    function emitNb(int nb) public isOwner {
        emit NumberEvent(nb);
    }
}
