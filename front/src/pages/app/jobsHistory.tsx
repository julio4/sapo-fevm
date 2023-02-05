import { useState } from "react";

import { JobProvider } from "@/components/Context/JobContext";
import Layout from "@/components/Layout/layout";
import {
  Text,
  Flex,
  Box,
  useColorModeValue,
  Heading,
  Divider,
  Button,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
} from "@chakra-ui/react";

import { CheckCircleIcon, TimeIcon, WarningIcon } from "@chakra-ui/icons";
import { useContractRead, useAccount } from "wagmi";
import { useDisclosure, defineStyleConfig } from "@chakra-ui/react";
import JobResult from "@/components/Jobs/JobResult";
import JobsList from "@/components/Jobs/JobsList";
import JobDetails from "@/components/Jobs/JobDetails";
import JobSummary from "@/components/Jobs/JobSummary";

export default function JobsHistory() {
  const [selectedJob, setSelectedJob] = useState<JobSummary | null>(null);
  const [listJobsAddress, setListJobsAddress] = useState<string[]>([]);
  const [currentJobAddress, setCurrentJobAddress] = useState<string>("");
  const { address, isConnecting, isDisconnected } = useAccount();
  const bridgeAddress = "0x40cBC46277b1a875027649D97239f65e5090D86B";

  // const jobsList: JobResult[] = [];

  // Step 1: Get list of jobs address with readContract getJobs
  // Step 2: Get job result (cid) with readContract getResult
  // Step 3: Create a new job result object with all data, and result is cid
  // Step 4: Push new job result object to jobsList

  let { data, isError, isLoading } = useContractRead({
    address: bridgeAddress,
    abi: [
      {
        inputs: [{ internalType: "address", name: "user", type: "address" }],
        name: "getJobs",
        outputs: [{ internalType: "address[]", name: "", type: "address[]" }],
        stateMutability: "view",
        type: "function",
      },
    ],
    args: [address],
    functionName: "getJobs",
  });

  return (
    <JobProvider>
      <Layout>
        {data && (
          <Flex w="full">
            <JobsList
              jobAddresses={data}
              onSelect={(selected) => setSelectedJob(selected)}
              selected={selectedJob ? selectedJob.address : null}
            />
            {selectedJob && <JobDetails job={selectedJob} />}
          </Flex>
        )}
      </Layout>
    </JobProvider>
  );
}