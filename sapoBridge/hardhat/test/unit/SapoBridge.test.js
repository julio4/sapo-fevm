const { assert, expect } = require("chai");
const { network, deployments, ethers } = require("hardhat");
const { developmentChains } = require("../../helper-hardhat-config");

!developmentChains.includes(network.name)
    ? describe.skip()
    : describe("Unit tests for SapoBridge", function () {
          let sapoBridge, sapoJob, deployer, user;

          beforeEach(async () => {
              accounts = await ethers.getSigners();
              deployer = accounts[0];
              user = accounts[1];
              await deployments.fixture(["all"]);

              sapoBridge = await ethers.getContract("SapoBridge", deployer);
              sapoJob = await ethers.getContract("SapoJob", deployer);
          });

          describe("constructor", function () {
              it("sets the bridge correctly", async function () {
                  const bridgeAddress = await sapoBridge.getBridgeAddress();
                  assert.equal(deployer.address, bridgeAddress);
              });
          });

          describe("setBridge", function () {
              it("reverts if caller is not owner", async function () {
                  sapoBridge = sapoBridge.connect(user);
                  await expect(
                      sapoBridge.setBridge(user.address)
                  ).to.be.revertedWith("Only Owner can call this function");
              });
              it("sets a new bridge address", async function () {
                  const initialAddress = await sapoBridge.getBridgeAddress();
                  await sapoBridge.setBridge(user.address);
                  const finalAddress = await sapoBridge.getBridgeAddress();

                  assert.equal(initialAddress, deployer.address);
                  assert.equal(finalAddress, user.address);
              });
          });

          describe("request", function () {
              it("reverts if value sent is lower than expected", async function () {
                  const value = ethers.utils
                      .parseUnits("5", "ether")
                      .toHexString();
                  await expect(
                      sapoBridge.request(
                          "QmYnaUZLWmbRTJzpx6kgxoAVT3ZAmhqWY6qWZm33v8PjNm",
                          value
                      )
                  ).to.be.revertedWith(
                      "The sent value is lower than the required value"
                  );
              });

              it("effectively deploy the job contract", async function () {
                  const value = ethers.utils
                      .parseUnits("5", "ether")
                      .toHexString();
                  const cid = "QmYnaUZLWmbRTJzpx6kgxoAVT3ZAmhqWY6qWZm33v8PjNm";

                  const jobAddressesBefore = await sapoBridge.getJobs(
                      deployer.address
                  );
                  await sapoBridge.request(cid, value, { value: value });
                  const jobAddressesAfter = await sapoBridge.getJobs(
                      deployer.address
                  );

                  assert.notEqual(
                      jobAddressesBefore,
                      jobAddressesAfter,
                      "A new job contract was not deployed"
                  );
              });

              it("effectively emits the event", async function () {
                  const value = ethers.utils
                      .parseUnits("5", "ether")
                      .toHexString();
                  const cid = "QmYnaUZLWmbRTJzpx6kgxoAVT3ZAmhqWY6qWZm33v8PjNm";

                  await expect(
                      sapoBridge.request(cid, value, { value: value })
                  ).to.emit(sapoBridge, "JobExecutionRequest");
              });

              it("bridge job got fund transfered", async function () {
                  const value = ethers.utils
                      .parseUnits("5", "ether")
                      .toString();
                  const cid = "QmYnaUZLWmbRTJzpx6kgxoAVT3ZAmhqWY6qWZm33v8PjNm";
                  await sapoBridge.setBridge(
                      "0x000000000000000000000000000000000000dEaD"
                  );
                  const bridge = await sapoBridge.getBridgeAddress();
                  const startBridgeBalance = (
                      await ethers.provider.getBalance(bridge)
                  ).toString();

                  await sapoBridge.request(cid, value, { value: value });

                  const finalBridgeBalance = (
                      await ethers.provider.getBalance(bridge)
                  ).toString();

                  assert.equal(finalBridgeBalance - startBridgeBalance, value);
              });
          });

          describe("saveResult", function () {
              it("reverts if caller is not bridge", async function () {
                  const value = ethers.utils
                      .parseUnits("5", "ether")
                      .toString();
                  const cid = "QmYnaUZLWmbRTJzpx6kgxoAVT3ZAmhqWY6qWZm33v8PjNm";
                  await sapoBridge.setBridge(user.address);
                  await sapoBridge.request(cid, value, { value: value });
                  const jobAddress = (
                      await sapoBridge.getJobs(deployer.address)
                  )[0];

                  await expect(
                      sapoBridge.saveResult(jobAddress, "cid of the result")
                  ).to.be.revertedWith("Only Bridge can call this function");
              });
          });

          describe("failAndRefund", function () {
              it("reverts if caller is not bridge", async function () {
                  const value = ethers.utils
                      .parseUnits("5", "ether")
                      .toString();
                  const cid = "QmYnaUZLWmbRTJzpx6kgxoAVT3ZAmhqWY6qWZm33v8PjNm";
                  await sapoBridge.setBridge(user.address);
                  await sapoBridge.request(cid, value, { value: value });
                  const jobAddress = (
                      await sapoBridge.getJobs(deployer.address)
                  )[0];

                  await expect(
                      sapoBridge.failAndRefund(jobAddress)
                  ).to.be.revertedWith("Only Bridge can call this function");
              });
          });
      });
