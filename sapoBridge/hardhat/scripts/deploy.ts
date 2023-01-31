import hre from 'hardhat';

async function main() {
  console.log('Deploying....');

  const owner = new hre.ethers.Wallet(
    process.env.WALLET_PRIVATE_KEY || 'undefined',
    hre.ethers.provider
  );
  const factory = (
    await hre.ethers.getContractFactory('SapoBridge', owner)
  );

  const bridge = (
    await factory.deploy()
  );

  await bridge.deployed();
  console.log('Bridge deployed to ', bridge.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
