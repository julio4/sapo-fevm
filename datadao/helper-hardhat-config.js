const { ethers } = require("hardhat")

const developmentChains = ["hardhat", "localhost"]

const networkConfig = {
    3141: {
        name: "hyperspace",
        tokensToBeMinted: 12000,
    },
    31337: {
        name: "localhost",
        tokensToBeMinted: 12000,
    },
}

module.exports = {
    networkConfig,
    developmentChains,
}
