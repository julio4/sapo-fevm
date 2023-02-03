const {
  networkConfig,
  developmentChains,
} = require("../helper-hardhat-config");
const { network } = require("hardhat"); //network coming from hardhat
require("dotenv").config();

// same as :   [...] = async (hre) => {const { getNamedAccounts, deployments } = hre;}
module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log } = deployments;
  const { deployer } = await getNamedAccounts();
  const { verify } = require("../utils/verify");
  const chainId = network.config.chainId;

  const SapoBridge = await deploy("SapoBridge", {
    from: deployer,
    log: true,
    args: [deployer],
    waitConfirmations: network.config.blockConfirmations || 1,
  });

  log("===============================================================\n");
  const SapoJob = await deploy("SapoJob", {
    from: deployer,
    log: true,
    args: [deployer],
    waitConfirmations: network.config.blockConfirmations || 1,
  });

  log("===============================================================\n");
};
module.exports.tags = ["all", "sapobridge"];
