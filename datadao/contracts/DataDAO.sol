// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./lib/filecoin-solidity/contracts/v0.8/MarketAPI.sol";

contract DataDAO {
    address public owner;

    mapping(uint256 => Proposal) public proposals;

    uint256 public proposalsCounter;

    event NewProposal(uint256 proposalId);

    struct Proposal {
        uint256 Id;
        address assignedProvider;
        bytes cidRaw;
        uint256 size;
        uint256 bountyAmount;
        uint256 endOfPin;
        uint256 upVote;
        uint256 downVote;
    }

    constructor() {
        owner = msg.sender;
    }

    function submitCid(
        address assignedProvider,
        bytes memory cidRaw,
        uint256 size,
        uint256 bountyAmount,
        uint256 endOfPin
    ) public {
        uint256 proposalId = proposalsCounter++;

        Proposal storage proposal = proposals[proposalId];
        proposal.Id = proposalId;
        proposal.assignedProvider = assignedProvider;
        proposal.cidRaw = cidRaw;
        proposal.size = size;
        proposal.bountyAmount = bountyAmount;
        proposal.endOfPin = endOfPin;

        emit NewProposal(proposalId);
    }
}
