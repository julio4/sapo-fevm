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
  Accordion,
  AccordionItem,
  AccordionButton,
  AccordionPanel,
  AccordionIcon,
} from "@chakra-ui/react";

import Image from "next/image";

import { useDisclosure } from "@chakra-ui/react";
import JobResult from "./JobResult";
import { CheckCircleIcon, TimeIcon, WarningIcon } from "@chakra-ui/icons";
import JobSummary from "./JobSummary";

const styleDetail = { fontWeight: "bold", marginRight: "1ex" };

const statusTextIcon = [
  <>
    {" "}
    pending
    <TimeIcon ml={2} size="20px" color={"blue.500"} />
  </>,
  <>
    {" "}
    completed
    <CheckCircleIcon ml={2} size="20px" color={"green.500"} />
  </>,
  <>
    {" "}
    failed
    <WarningIcon ml={2} size="20px" color={"red.500"} />
  </>,
];

const JobDetails = ({ job }: { job: JobSummary | null }) => {
  const jobBgGradient = useColorModeValue(
    "linear(to-r, blue.200, blue.200)",
    "linear(to-r, blue.900, blue.900)"
  );
  const jobColor = useColorModeValue("gray.700", "gray.200");

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
        <Heading
          size="md"
          mb={2}
          display={"flex"}
          justifyContent={"space-between"}
        >
          Job Detail
        </Heading>
      </Box>
      <Divider />

      <Box px={4} py={2}>
        <Text>
          <span style={styleDetail}>
            Job n° {job ? job.id : "Nan"} {statusTextIcon[job.status]}{" "}
          </span>
        </Text>
        <Text>
          <span style={styleDetail}>Bacalhau ID:</span> {job?.jobId}
        </Text>
        <Text>
          <span style={styleDetail}>Job Address:</span> {job?.address}
        </Text>
        <Text>
          <span style={styleDetail}>Job Status:</span>{" "}
          {statusTextIcon[job.status]}
        </Text>
        <Text mb={"5%"}>
          <span style={styleDetail}>Exit Code:</span> {job?.exitCode}
        </Text>
        <Text>
          <span style={styleDetail}>About the results:</span>
        </Text>
        <Accordion allowToggle>
          <AccordionItem>
            <h2>
              <AccordionButton>
                <Box as="span" flex="1" textAlign="left">
                  Outputs
                </Box>
                <AccordionIcon />
              </AccordionButton>
            </h2>
            <AccordionPanel pb={4}>
              {job?.outputsNamesAndCids?.length ? (
                job.outputsNamesAndCids.map((output) => (
                  <>
                    <Box
                      display={"flex"}
                      alignItems={"center"}
                      justifyContent={"flex-start"}
                    >
                      <Text mr={"5%"}>
                        <span style={styleDetail}>{output.name}:</span>{" "}
                        {output.cid}
                      </Text>
                      <Button
                        size={"sm"}
                        onClick={() =>
                          window.open(`https://ipfs.io/ipfs/${output.cid}`)
                        }
                      >
                        Show on IPFS.io
                      </Button>
                    </Box>
                  </>
                ))
              ) : (
                <Text>No outputs</Text>
              )}
            </AccordionPanel>
          </AccordionItem>

          <AccordionItem>
            <h2>
              <AccordionButton>
                <Box as="span" flex="1" textAlign="left">
                  Stderr
                </Box>
                <AccordionIcon />
              </AccordionButton>
            </h2>
            <AccordionPanel pb={4}>
              {job?.stderr ? (
                <>
                  <Text>Here is the stderr of the job:</Text>
                  <Text style={styleDetail}>Stderr:{job.stderr}</Text>
                </>
              ) : (
                <Text>No stderr</Text>
              )}
            </AccordionPanel>
          </AccordionItem>
          <AccordionItem>
            <h2>
              <AccordionButton>
                <Box as="span" flex="1" textAlign="left">
                  Stdout
                </Box>
                <AccordionIcon />
              </AccordionButton>
            </h2>
            <AccordionPanel pb={4}>
              {job?.stdout ? (
                <>
                  <Text>Here is the stdout of the job:</Text>
                  <Text>
                    <span style={styleDetail}>Stdout:</span> {job.stdout}
                  </Text>
                </>
              ) : (
                <Text>No stdout</Text>
              )}
            </AccordionPanel>
          </AccordionItem>
        </Accordion>

        {/*         <Text display="flex" alignItems="center" mb={2}>
          <span style={styleDetail}>Stdout:</span>{" "}
          {!job || job.status === 0 || job.status === null ? (
            job?.stdout
          ) : (
            <></>
          )}
          {job?.status === 1 ? (
            <Button
              mb={1}
              ml={6}
              onClick={() => handleClickShow("stdout")}
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
        </Text> */}
        {job?.outputUrl ? (
          <Box display={"flex"} justifyContent={"center"}>
            <Box
              borderRadius={"lg"}
              width={"40%"}
              display={"flex"}
              flexDirection={"column"}
              overflow={"hidden"}
              alignItems={"center"}
              mt={4}
            >
              <Image
                loader={() => job?.outputUrl}
                src={job?.outputUrl}
                width="100"
                height="100"
                alt="Image of the output"
                layout="responsive"
              />
              <Text fontWeight={"bold"} mt={2}>
                Your beautiful result
              </Text>
            </Box>
          </Box>
        ) : (
          <></>
        )}
      </Box>
    </Box>
  );
};

export default JobDetails;
