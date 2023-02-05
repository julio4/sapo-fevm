import { useState } from "react";

import { JobProvider } from "@/components/Context/JobContext";
import Layout from "@/components/Layout/layout";
import {
  Flex,
} from "@chakra-ui/react";

import { useContractRead, useAccount } from "wagmi";
import JobsList from "@/components/Jobs/JobsList";
import JobDetails from "@/components/Jobs/JobDetails";
import JobSummary from "@/components/Jobs/JobSummary";
import AddressSapoBridge from "@/constants/AddressSapoBridge.json";

export default function JobsHistory() {
  const [selectedJob, setSelectedJob] = useState<JobSummary | null>(null);
  const { address, isConnecting, isDisconnected } = useAccount();

  // const jobsList: JobResult[] = [];

  // Step 1: Get list of jobs address with readContract getJobs
  // Step 2: Get job result (cid) with readContract getResult
  // Step 3: Create a new job result object with all data, and result is cid
  // Step 4: Push new job result object to jobsList

  let { data, isError, isLoading } = useContractRead({
    address: AddressSapoBridge.address,
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
