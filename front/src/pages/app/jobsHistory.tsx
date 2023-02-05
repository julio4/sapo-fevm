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

export default function JobsHistory() {
  const [job, setJob] = useState<JobResult | null>(null);
  const [jobAddress, setJobAddress] = useState<JobResult | null>(null);
  const [listJobsAddress, setListJobsAddress] = useState<string[]>([]);
  const [currentJobAddress, setCurrentJobAddress] = useState<string>("");
  const { address, isConnecting, isDisconnected } = useAccount();

  // const jobsList: JobResult[] = [];

  // Step 1: Get list of jobs address with readContract getJobs
  // Step 2: Get job result (cid) with readContract getResult
  // Step 3: Create a new job result object with all data, and result is cid
  // Step 4: Push new job result object to jobsList

  let { data, isError, isLoading } = useContractRead({
    address: "0x0621B6d8d05CA0158429371ADCE09586965BA299",
    abi: [
      {
        inputs: [{ internalType: "address", name: "user", type: "address" }],
        name: "getJobs",
        outputs: [{ internalType: "address[]", name: "", type: "address[]" }],
        stateMutability: "view",
        type: "function",
      },
    ],
    args: ["0x07705e0fFdc102e6144e22261477B4cE46Da350A"],
    functionName: "getJobs",
  });

  return (
    <JobProvider>
      <Layout>
        {data && (
          <Flex w="full">
            <JobsList
              jobAddresses={data}
              onSelect={(selected) => setJob(selected)}
              selected={jobAddress}
            />
            {job && <JobDetails job={job} />}
          </Flex>
        )}
      </Layout>
    </JobProvider>
  );
}
