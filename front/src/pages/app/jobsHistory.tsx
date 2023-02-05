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
  Badge,
  Slide,
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

import { useDisclosure, defineStyleConfig } from "@chakra-ui/react";

const statusText = {
  0: "Completed ",
  1: "Pending",
  2: "Failed",
};

const statusTextIcon = {
  0: (
    <>
      {" "}
      Completed
      <CheckCircleIcon ml={2} size="20px" color={"green.500"} />
    </>
  ),
  1: (
    <>
      {" "}
      Pending
      <TimeIcon ml={2} size="20px" color={"blue.500"} />
    </>
  ),
  2: (
    <>
      {" "}
      Failed
      <WarningIcon ml={2} size="20px" color={"red.500"} />
    </>
  ),
};

const styleDetail = { fontWeight: "bold", marginRight: "1ex" };

type JobResult = {
  id: number;
  address: string;
  status: number; // 0: pending, 1: success, 2: fail
  result: string; // CID
  exitCode: number;
  outputs: string[]; // CIDs
  stderr: string;
  stdout: string;
};

const jobsList: JobResult[] = [
  {
    id: 1,
    address: "0x07705e0fFdc102e6144e22261477B4cE46Da341C",
    status: 0,
    result: "QmRSMeMFFVgjXRSPkSrpJog5ThfrxZKCHboen2uTGZWgtc",
    exitCode: 0,
    outputs: ["QmX2", "QmX3"],
    stderr: "",
    stdout: "Hello world",
  },
  {
    id: 2,
    address: "0x07705e0fFdc102e6144e22261477B4cE46DaABC",
    status: 2,
    result: "Not available",
    exitCode: 1,
    outputs: [],
    stderr: "Error",
    stdout: "",
  },
  {
    id: 3,
    address: "0x07705e0fFdc102e6144e22261477BSA8D46aCvC1",
    status: 0,
    result: "QmqHEDJFFVgjnQpPknrpoOg5ThfrxZKCHboen2uTGZXcf1",
    exitCode: 1,
    outputs: ["funfile.txt", "funfile2.txt"],
    stderr: "Error",
    stdout: "",
  },
  {
    id: 4,
    address: "0x07705e0fFdc102e6144e22261477BSA8D46aCvC1",
    status: 1,
    result: "Not available yet",
    exitCode: 1,
    outputs: [],
    stderr: "",
    stdout: "",
  },
];

const JobsList = ({
  jobs,
  selected,
  onSelect,
}: {
  jobs: JobResult[];
  onSelect: (job: JobResult) => void;
  selected: JobResult | null;
}) => {
  return (
    <Box
      w="full"
      m={4}
      mr={2}
      py={4}
      bg={useColorModeValue("whiteAlpha.700", "transparent")}
      backdropFilter="auto"
      backdropBlur="2px"
      border="1px"
      borderColor={useColorModeValue("gray.200", "gray.700")}
      borderRadius="lg"
    >
      <Box px={4}>
        <Heading size="md" mb={2}>
          Jobs List
        </Heading>
      </Box>
      <Button
        onClick={() => {
          console.log("allResults", jobs);
        }}
      >
        TEST
      </Button>
      <Divider />

      {jobs.map((job) => (
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
              {selected && job.result.length > 22
                ? job.result.substring(0, 22) + "..."
                : job.result}
            </Text>

            <Badge
              colorScheme={
                job.status === 0 ? "green" : job.status === 1 ? "blue" : "red"
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
      ))}
    </Box>
  );
};

const JobDetails = ({ job }: { job: JobResult | null }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [modalTitle, setModalTitle] = useState("");
  const [modalContent, setModalContent] = useState("");

  function handleClickShow(from: string) {
    console.log("from:", from);
    console.log(from === "outputs");
    if (from == "result") {
      setModalTitle("Result");
      setModalContent(job?.result);
    } else if (from == "outputs") {
      setModalTitle("outputs");
      setModalContent(job?.outputs.join(", "));
    } else if (from === "stderr") {
      setModalTitle("Stderr");
      setModalContent(job?.stderr);
    } else if (from === "stdout") {
      setModalTitle("Stdout");
      setModalContent(job?.stdout);
    }
    onOpen();
  }

  return (
    <Box
      w={job ? "full" : "0"}
      m={4}
      ml={2}
      py={4}
      bg={useColorModeValue("whiteAlpha.700", "transparent")}
      backdropFilter="auto"
      backdropBlur="2px"
      border="1px"
      borderColor={useColorModeValue("gray.200", "gray.700")}
      borderRadius="lg"
    >
      <Box px={4}>
        <Heading size="md" mb={2}>
          Job Detail
        </Heading>
      </Box>
      <Divider />

      <Box px={4} py={2}>
        <Text>
          <span style={styleDetail}>Job ID:</span> {job?.id}
        </Text>
        <Text>
          <span style={styleDetail}>Job Address:</span> {job?.address}
        </Text>
        <Text>
          <span style={styleDetail}>Job Status:</span>{" "}
          {statusTextIcon[job.status]}
        </Text>
        <Text>
          <span style={styleDetail}>Cid Result:</span> {job?.result}
        </Text>
        <Text>
          <span style={styleDetail}>Exit Code:</span> {job?.exitCode}
        </Text>
        <Text display="flex" alignItems="center" mb={2}>
          <span style={styleDetail}>Outputs: </span> {job?.outputs}
          {!(job?.status === 1) ? (
            <Button
              mb={1}
              ml={6}
              onClick={() => handleClickShow("outputs")}
              size="xs"
              bgGradient={useColorModeValue(
                "linear(to-r, blue.200, blue.200)",
                "linear(to-r, blue.900, blue.900)"
              )}
              color={useColorModeValue("gray.700", "gray.200")}
              fontWeight="bold"
            >
              Show
            </Button>
          ) : (
            <Text fontStyle={"italic"}>Waiting for result...</Text>
          )}
        </Text>
        <Text display="flex" alignItems="center" mb={2}>
          <span style={styleDetail}>Stderr:</span> {job?.stderr}
          {!(job?.status === 1) ? (
            <Button
              mb={1}
              ml={6}
              onClick={() => handleClickShow("outputs")}
              size="xs"
              bgGradient={useColorModeValue(
                "linear(to-r, blue.200, blue.200)",
                "linear(to-r, blue.900, blue.900)"
              )}
              color={useColorModeValue("gray.700", "gray.200")}
              fontWeight="bold"
            >
              Show
            </Button>
          ) : (
            <Text fontStyle={"italic"}>Waiting for result...</Text>
          )}
        </Text>
        <Text display="flex" alignItems="center" mb={2}>
          <span style={styleDetail}>Stdout:</span> {job?.stdout}
          {!(job?.status === 1) ? (
            <Button
              mb={1}
              ml={6}
              onClick={() => handleClickShow("outputs")}
              size="xs"
              bgGradient={useColorModeValue(
                "linear(to-r, blue.200, blue.200)",
                "linear(to-r, blue.900, blue.900)"
              )}
              color={useColorModeValue("gray.700", "gray.200")}
              fontWeight="bold"
            >
              Show
            </Button>
          ) : (
            <Text fontStyle={"italic"}>Waiting for result...</Text>
          )}
        </Text>
      </Box>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>{modalTitle}</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text>{modalContent}</Text>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={onClose}>
              Close
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default function JobsHistory() {
  const [job, setJob] = useState<JobResult | null>(null);

  return (
    <JobProvider>
      <Layout>
        <Flex w="full">
          <JobsList
            jobs={jobsList}
            onSelect={(selected) => setJob(selected)}
            selected={job}
          />
          {job && <JobDetails job={job} />}
        </Flex>
      </Layout>
    </JobProvider>
  );
}
