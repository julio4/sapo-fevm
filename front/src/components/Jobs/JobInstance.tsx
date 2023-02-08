import { useEffect, useState } from "react";

import { JobProvider } from "@/components/Context/JobContext";
import Layout from "@/components/Layout/layout";
import {
  Text,
  Flex,
  Box,
  useColorModeValue,
  Divider,
  Button,
  Badge,
} from "@chakra-ui/react";

import { CheckCircleIcon, TimeIcon, WarningIcon } from "@chakra-ui/icons";
import { useContractRead, useAccount } from "wagmi";
import { useDisclosure, defineStyleConfig } from "@chakra-ui/react";
import JobResult from "@/components/Jobs/JobResult";
import JobsList from "@/components/Jobs/JobsList";
import JobDetails from "@/components/Jobs/JobDetails";
import JobSummary from "./JobSummary";

const jobsList: JobResult[] = [
  {
    id: 0,
    address: "0x07705e0fFdc102e6144e22261477B4cE46Da341C",
    exitCode: 0,
    outputs: ["QmX2", "QmX3"],
    stderr: "",
    stdout: "Hello world",
  },
  {
    id: 1,
    address: "0x07705e0fFdc102e6144e22261477B4cE46DaABC",
    exitCode: 1,
    outputs: [],
    stderr: "Error",
    stdout: "",
  },
  {
    id: 2,
    address: "0x07705e0fFdc102e6144e22261477BSA8D46aCvC1",
    exitCode: 1,
    outputs: ["funfile.txt", "funfile2.txt"],
    stderr: "Error",
    stdout: "",
  },
  {
    id: 3,
    address: "0x07705e0fFdc102e6144e22261477BSA8D46aCvC1",
    exitCode: 1,
    outputs: [],
    stderr: "",
    stdout: "",
  },
];

const statusText: string[] = ["Pending", "Completed", "Failed", "Loading..."];

const statusColor: string[] = ["blue", "green", "red", "orange"];

const JobInstance = ({
  onSelect,
  jobAddress,
  selected,
  id,
}: {
  onSelect: (select: JobSummary) => void;
  jobAddress: `0x${string}`;
  selected: `0x${string}` | null;
  id: number;
}) => {
  const [outputUrl, setOutputUrl] = useState<string>("");
  // TODO Bacalhau
  let [job, setJob] = useState<JobResult>({
    id: id,
    address: jobAddress,
    exitCode: 1,
    outputs: [""],
    stderr: "",
    stdout: "",
  });

  let {
    data: jobId,
    isLoading: jobIdIsLoading,
    isError: jobIdIsError,
  } = useContractRead({
    address: jobAddress,
    abi: [
      {
        inputs: [],
        name: "getResult",
        outputs: [
          {
            internalType: "string",
            name: "",
            type: "string",
          },
        ],
        stateMutability: "view",
        type: "function",
      },
    ],
    functionName: "getResult",
  });

  let res = useContractRead({
    address: jobAddress,
    abi: [
      {
        inputs: [],
        name: "getStatus",
        outputs: [
          {
            internalType: "enum SapoJob.Status",
            name: "",
            type: "uint8",
          },
        ],
        stateMutability: "view",
        type: "function",
      },
    ],
    functionName: "getStatus",
  });

  let { data: status, isLoading: statusIsLoading } = res;

  const unselectedColor = useColorModeValue("whiteAlpha.700", "transparent");
  const selectedColor = useColorModeValue("teal.50", "teal.800");

  async function getFileFromGateway(/* cid, path */) {
    try {
      const response = await fetch(
        `https://ipfs.io/ipfs/QmcGJ2KkiUwSpbdFiEgzVgYVJDfn11WBg6tUUGWAGeRa2N/outputs/image0.png`
      );
      const formattedUrl = response.url.replace(
        /http.*\/ipfs\//,
        "https://ipfs.io/ipfs/"
      );
      setOutputUrl(formattedUrl);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    getFileFromGateway();
  }, [outputUrl]);

  return (
    <Box
      key={job.id}
      transition="all 0.2s ease-in"
      bg={selected && selected === jobAddress ? selectedColor : unselectedColor}
      _hover={{
        bg: useColorModeValue("blackAlpha.100", "blackAlpha.400"),
        cursor: "pointer",
      }}
      onClick={() => {
        onSelect({
          ...job,
          jobId: jobId ? jobId : null,
          status: status ? status : null,
          outputUrl: outputUrl ? outputUrl : null,
        });
      }}
    >
      <Flex px={4} py={2} justifyContent="space-between" paddingX="5%">
        <Text>{job.id}</Text>

        <Text width={!selected ? "35%" : "50%"}>
          {selected && jobId && jobId.length > 22
            ? jobId.substring(0, 22) + "..."
            : jobId}
        </Text>
        <Badge
          colorScheme={statusColor[statusIsLoading ? 3 : status ? status : 0]}
          style={{ display: "flex", alignItems: "center" }}
          borderRadius="8px"
        >
          <Text width="6vw" align="center">
            {statusText[statusIsLoading ? 3 : status ? status : 0]}
          </Text>
        </Badge>
      </Flex>
      <Divider />
    </Box>
  );
};

export default JobInstance;
