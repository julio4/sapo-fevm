require("@nomicfoundation/hardhat-toolbox")
require("hardhat-deploy")
require("hardhat-deploy-ethers")
require("./tasks")
require("dotenv").config()

const PRIVATE_KEY = process.env.PRIVATE_KEY || "key"

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: "0.8.17",
    defaultNetwork: "hyperspace",
    networks: {
        hyperspace: {
            chainId: 3141,
            url: "https://api.hyperspace.node.glif.io/rpc/v1",
            accounts: [PRIVATE_KEY],
        },
        localhost: {
            chainId: 31337,
            blockConfirmations: 1,
        },
    },
    paths: {
        sources: "./contracts",
        tests: "./test",
        cache: "./cache",
        artifacts: "./artifacts",
    },
}
