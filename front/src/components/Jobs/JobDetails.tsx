import { useState } from "react";

import {
  Text,
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

import { useDisclosure } from "@chakra-ui/react";
import JobResult from "./JobResult";
import { CheckCircleIcon, TimeIcon, WarningIcon } from "@chakra-ui/icons";
import JobSummary from "./JobSummary";

const styleDetail = { fontWeight: "bold", marginRight: "1ex" };

const statusTextIcon = [
  (
    <>
      {" "}
      Pending
      <TimeIcon ml={2} size="20px" color={"blue.500"} />
    </>
  ),
  (
    <>
      {" "}
      Completed
      <CheckCircleIcon ml={2} size="20px" color={"green.500"} />
    </>
  ),
  (
    <>
      {" "}
      Failed
      <WarningIcon ml={2} size="20px" color={"red.500"} />
    </>
  ),
];

const JobDetails = ({ job }: { job: JobSummary | null }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [modalTitle, setModalTitle] = useState("");
  const [modalContent, setModalContent] = useState("");

  function handleClickShow(from: string) {
    if (from == "result") {
      setModalTitle("Result");
      setModalContent(job?.jobId);
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

  const jobBgGradient = useColorModeValue(
    "linear(to-r, blue.200, blue.200)",
    "linear(to-r, blue.900, blue.900)"
  )
  const jobColor = useColorModeValue("gray.700", "gray.200")

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
          <span style={styleDetail}>Job ID:</span> {job ? job.id : "no job"}
        </Text>
        <Text>
          <span style={styleDetail}>Job Address:</span> {job?.address}
        </Text>
        <Text>
          <span style={styleDetail}>Job Status:</span>{" "}
          {statusTextIcon[job.status]}
        </Text>
        <Text>
          <span style={styleDetail}>JobId Result:</span> {job?.jobId}
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
              bgGradient={jobBgGradient}
              color={jobColor}
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
              bgGradient={jobBgGradient}
              color={jobColor}
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
              bgGradient={jobBgGradient}
              color={jobColor}
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

export default JobDetails;