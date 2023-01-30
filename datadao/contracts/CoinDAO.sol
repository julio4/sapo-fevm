// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

error CoinDAO__NotEnoughBalance();

contract CoinDAO {
    mapping(address => uint) balances;
    uint256 private i_tokensToBeMinted;

    constructor(uint256 tokensToBeMinted) {
        balances[tx.origin] = tokensToBeMinted;
        i_tokensToBeMinted = tokensToBeMinted;
    }

    function isMember() public returns (bool membership) {
        return (getBalance(msg.sender) > 0);
    }

    function sendCoin(address receiver, uint amount) public returns (bool sufficient) {
        if (balances[msg.sender] < amount) {
            revert CoinDAO__NotEnoughBalance();
        }
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        return true;
    }

    function getBalance(address addr) public view returns (uint) {
        return balances[addr];
    }

    function getMintedTokenBalance() public view returns (uint256) {
        return i_tokensToBeMinted;
    }
}
