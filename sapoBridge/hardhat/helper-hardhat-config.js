const { ethers } = require("hardhat");

const networkConfig = {
  3141: {
    name: "hyperspace",
    waitConfirmations: "3",
  },
  31337: {
    name: "localhost",
  },
};

const developmentChains = ["hardhat", "localhost"];

module.exports = {
  networkConfig,
  developmentChains,
};
