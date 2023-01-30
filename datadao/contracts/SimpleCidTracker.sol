// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./CoinDAO.sol";

contract SimpleCidTracker {
    address public owner;

    CoinDAO coinContract;

    mapping(string => bool) public cidSet;
    // mapping(bytes => uint) public cidSizes;
    mapping(bytes => mapping(uint64 => bool)) public cidProviders;
    mapping(bytes => mapping(uint64 => bool)) public cidToValidate;

    string[] public cids;

    /* Events */
    event EventCalled(address caller, uint256 counter);

    constructor(address _addressCoinContract) {
        owner = msg.sender;
        coinContract = CoinDAO(_addressCoinContract);
    }

    function isCIDStored(bytes memory cidraw, uint64 provider) internal view returns (bool) {
        bool alreadyStoring = cidProviders[cidraw][provider];
        return !alreadyStoring;
    }

    function proposeCID(bytes memory cidraw, uint64 provider /* , uint size */) public {
        require(coinContract.isMember(), "caller must be a member");
        cidToValidate[cidraw][provider] = true;
    }

    // function authorizeData(bytes memory cidraw, uint64 provider /* , uint size */) public {
    //     require(coinContract.isMember(), "caller must be a member");
    //     require(cidSet[cidraw], "cid must be added before authorizing");
    //     // require(cidSizes[cidraw] == size, "data size must match expected");
    //     require(isCIDStored(cidraw, provider), "cid is already stored by this provider");
    //     cidProviders[cidraw][provider] = true;
    // }

    function addCID(string memory cidraw) public {
        cids.push(cidraw);
    }

    function getCidStored() public view returns (string[] memory) {
        return cids;
    }

    // function addCIDForBounty(bytes calldata cidraw, uint size) public {
    //     require(msg.sender == owner);
    //     cidSet[cidraw] = true;
    //     cidSizes[cidraw] = size;
    //     cids.push(cidraw);
    // }
}
