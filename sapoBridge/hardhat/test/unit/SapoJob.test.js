const { assert, expect } = require("chai");
const { network, deployments, ethers } = require("hardhat");
const { developmentChains } = require("../../helper-hardhat-config");

!developmentChains.includes(network.name)
    ? describe.skip()
    : describe("Unit tests for SapoJob", function () {
          let sapoJob, deployer, user;

          beforeEach(async () => {
              accounts = await ethers.getSigners();
              deployer = accounts[0];
              user = accounts[1];
              await deployments.fixture(["all"]);

              sapoJob = await ethers.getContract("SapoJob", deployer);
          });

          describe("constructor", function () {
              it("sets the initiator correctly", async function () {
                  const initiatorAddress = await sapoJob.getInitiator();
                  assert.equal(deployer.address, initiatorAddress);
              });
          });

          describe("saveResult", function () {
              it("should revert if not caller is not bridge", async function () {
                  sapoJob = sapoJob.connect(user);
                  await expect(
                      sapoJob.saveResult("Just a simple result")
                  ).to.be.revertedWith("Caller is not bridge");
              });

              it("should revert if job already completed", async function () {
                  const startState = await sapoJob.getState();
                  await sapoJob.saveResult("Just a simple result");
                  const finalState = await sapoJob.getState();

                  assert.equal(startState, false);
                  assert.equal(finalState, true);
                  await expect(
                      sapoJob.saveResult("An other result")
                  ).to.be.revertedWith("Job is already finished");
              });

              it("should emit an event when saving result", async function () {
                  await expect(
                      sapoJob.saveResult("Just a simple result")
                  ).to.emit(sapoJob, "JobSucceeded");
              });

              it("should effectively saving the result", async function () {
                  const result = "Just a simple result";
                  await sapoJob.saveResult("Just a simple result");
                  const resultAnswer = await sapoJob.getResult();
                  assert.equal(resultAnswer, result);
              });
          });

          describe("getResult", function () {
              it("should revert if caller is not initiator", async function () {
                  sapoJob = sapoJob.connect(user);
                  expect(sapoJob.getResult()).to.be.revertedWith(
                      "Caller is not initiator"
                  );
              });
              it("should revert if not completed", async function () {
                  expect(sapoJob.getResult()).to.be.revertedWith(
                      "Job is not finished"
                  );
              });
              it("should effectively returning the result", async function () {
                  const result = "Just a simple result";
                  await sapoJob.saveResult("Just a simple result");
                  const resultAnswer = await sapoJob.getResult();
                  assert.equal(resultAnswer, result);
              });
          });
          describe("failAndRefund", function () {
              it("should revert if caller is not bridge", async function () {
                  sapoJob = sapoJob.connect(user);
                  expect(sapoJob.failAndRefund()).to.be.revertedWith(
                      "Caller is not bridge"
                  );
              });
              it("should revert if completed", async function () {
                  const startState = await sapoJob.getState();
                  await sapoJob.saveResult("Just a simple result");
                  const finalState = await sapoJob.getState();

                  assert.equal(startState, false);
                  assert.equal(finalState, true);
                  await expect(sapoJob.failAndRefund()).to.be.revertedWith(
                      "Job is already finished"
                  );
              });
              it("should emit an event", async function () {
                  await expect(sapoJob.failAndRefund()).to.emit(
                      sapoJob,
                      "JobFailed"
                  );
              });
              it("should refund the amount paid to initiator", async function () {
                  const amountPaid = ethers.utils
                      .parseUnits("1", "ether")
                      .toHexString();
                  const initiator = deployer;
                  await sapoJob.fund({ value: amountPaid });

                  const initiatorStartBalance =
                      await ethers.provider.getBalance(initiator.address);

                  const transactionResponse = await sapoJob.failAndRefund();
                  const transactionReceipt = await transactionResponse.wait(1);
                  const { gasUsed, effectiveGasPrice } = transactionReceipt;
                  const gasCost = gasUsed.mul(effectiveGasPrice);

                  const initiatorFinalBalance =
                      await ethers.provider.getBalance(initiator.address);

                  assert.equal(
                      initiatorStartBalance.toString(),
                      initiatorFinalBalance.add(gasCost).toString()
                  );
              });
          });
      });
