import { useEffect, useState } from "react";

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
  Badge,
} from "@chakra-ui/react";

import { CheckCircleIcon, TimeIcon, WarningIcon } from "@chakra-ui/icons";
import { useContractRead, useAccount } from "wagmi";
import { useDisclosure, defineStyleConfig } from "@chakra-ui/react";
import JobResult from "@/components/Jobs/JobResult";
import JobsList from "@/components/Jobs/JobsList";
import JobDetails from "@/components/Jobs/JobDetails";

const jobsList: JobResult[] = [
  {
    id: 0,
    address: "0x07705e0fFdc102e6144e22261477B4cE46Da341C",
    status: 0,
    jobId: "QmRSMeMFFVgjXRSPkSrpJog5ThfrxZKCHboen2uTGZWgtc",
    exitCode: 0,
    outputs: ["QmX2", "QmX3"],
    stderr: "",
    stdout: "Hello world",
  },
  {
    id: 1,
    address: "0x07705e0fFdc102e6144e22261477B4cE46DaABC",
    status: 2,
    jobId: "Not available",
    exitCode: 1,
    outputs: [],
    stderr: "Error",
    stdout: "",
  },
  {
    id: 2,
    address: "0x07705e0fFdc102e6144e22261477BSA8D46aCvC1",
    status: 0,
    jobId: "QmqHEDJFFVgjnQpPknrpoOg5ThfrxZKCHboen2uTGZXcf1",
    exitCode: 1,
    outputs: ["funfile.txt", "funfile2.txt"],
    stderr: "Error",
    stdout: "",
  },
  {
    id: 3,
    address: "0x07705e0fFdc102e6144e22261477BSA8D46aCvC1",
    status: 1,
    jobId: "Not available yet",
    exitCode: 1,
    outputs: [],
    stderr: "",
    stdout: "",
  },
];


const statusText : string[] = [
  "Completed",
  "Pending",
  "Failed",
  "Loading...",
];

const statusColor : string[] = [
  "green",
  "blue",
  "red",
  "orange",
]

const JobInstance = ({
  onSelect,
  jobAddress,
  selected,
  key,
}: {
  onSelect: (select: JobResult) => void;
  jobAddress: `0x${string}`;
  selected: JobResult | null;
  key: number;
}) => {
  useEffect(() => {

  }, []);

  let [job, setJob] = useState<JobResult>({
    id: key,
    address: jobAddress,
    status: 3,
    jobId: "Loading",
    exitCode: 1,
    outputs: ["funfile.txt", "funfile2.txt"],
    stderr: "Error",
    stdout: "hi world",
  });

  let { data, isError, isLoading } = useContractRead({
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
    onSuccess(result) {
      handleGetResultSuccess(result);
    },
  });

  function handleGetResultSuccess(data: string) {
    setJob({...job, jobId: data});
  }

  console.log("test: ", {jobAddress, job});

  return (
        <Box
          key={job.id}
          transition="all 0.2s ease-in"
          bg={
            selected && selected.id === job.id
              ? useColorModeValue("teal.50", "teal.800")
              : useColorModeValue("whiteAlpha.700", "transparent")
          }
          _hover={{
            bg: useColorModeValue("blackAlpha.100", "blackAlpha.400"),
            cursor: "pointer",
          }}
          onClick={() => onSelect(job)}
        >
          <Flex px={4} py={2} justifyContent="space-between" paddingX="5%">
            <Text>{job.id}</Text>

            <Text width={!selected ? "35%" : "50%"}>
              {selected && job.jobId.length > 22
                ? job.jobId.substring(0, 22) + "..."
                : job.jobId}
            </Text>

            <Badge
              colorScheme={
                statusColor[job.status]
              }
              style={{ display: "flex", alignItems: "center" }}
              borderRadius="8px"
            >
              <Text width="6vw" align="center">
                {statusText[job.status]}
              </Text>
            </Badge>
          </Flex>
          <Divider />
        </Box>
      );
}

export default JobInstance;