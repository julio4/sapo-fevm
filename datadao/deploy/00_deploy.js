require("hardhat-deploy")
require("hardhat-deploy-ethers")

const { networkConfig } = require("../helper-hardhat-config")

const private_key = network.config.accounts[0]
const wallet = new ethers.Wallet(private_key, ethers.provider)

module.exports = async ({ deployments }) => {
    console.log("Wallet Ethereum Address:", wallet.address)
    const chainId = network.config.chainId

    //deploy MembershipNFT
    const membershipNFT = await ethers.getContractFactory("MembershipNFT")
    console.log("Deploying MembershipNFT...")
    const membershipNFTContract = await membershipNFT.deploy()
    console.log("a")
    await membershipNFTContract.deployed()
    console.log("MembershipNFT deployed to:", membershipNFTContract.address)

    //deploy SimpleCIDTracker
    var adminsList = ["0x07705e0fFdc102e6144e22261477B4cE46Da350A"]
    const dataDAOExample = await ethers.getContractFactory("DataDAOExample")
    console.log("Deploying DataDAOExample...")
    const dataDAOExampleContract = await dataDAOExample.deploy(adminsList, membershipnftaddress, {
        gasLimit: 1000000000,
    })

    await dataDAOExampleContract.deployed()
    console.log("DataDAOExample deployed to:", dataDAOExampleContract.address)
}
