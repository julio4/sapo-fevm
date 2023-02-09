import { useState } from "react";

import {
  Text,
  Box,
  useColorModeValue,
  Heading,
  Divider,
  Button,
  Accordion,
  AccordionItem,
  AccordionButton,
  AccordionPanel,
  AccordionIcon,
  Container,
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

const statusIcon = [
  <CheckCircleIcon ml={2} size="20px" color={"green.500"} key="success" />,
  <WarningIcon ml={2} size="20px" color={"red.500"} key="failed" />,
  <WarningIcon ml={2} size="20px" color={"blue.500"} key="unknown" />,
];

const JobDetails = ({ job }: { job: JobSummary | null }) => {
  const jobBgGradient = useColorModeValue(
    "linear(to-r, blue.200, blue.200)",
    "linear(to-r, blue.900, blue.900)"
  );
  const jobColor = useColorModeValue("gray.700", "gray.200");

  function jobStatusIcon(status: number) {
    let icon;
    switch (job?.exitCode) {
      case "0":
        icon = statusIcon[0];
        break;
      case "1":
        icon = statusIcon[1];
        break;
      default:
        icon = statusIcon[2];
        break;
    }
    return icon;
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
      maxHeight={"95%"}
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
          <span style={styleDetail}>Cid Result:</span> {job?.cidResult}
        </Text>
        <Text mb={"5%"} display={"flex"}>
          <span style={styleDetail}>Exit Code:</span>{" "}
          {job?.exitCode ? (
            <>
              {job.exitCode} {jobStatusIcon(Number(job.exitCode))}
            </>
          ) : (
            <Text>
              NaN
              <WarningIcon ml={2} size="20px" color={"blue.500"} />{" "}
            </Text>
          )}
        </Text>
        <Text>
          <span style={styleDetail}>About the results:</span>
        </Text>
        <Accordion allowToggle>
          <AccordionItem>
            <h2>
              <AccordionButton>
                <Box
                  as="span"
                  flex="1"
                  textAlign="left"
                  fontWeight={"semibold"}
                  fontSize={"lg"}
                >
                  Outputs
                </Box>
                <AccordionIcon />
              </AccordionButton>
            </h2>
            <AccordionPanel pb={4}>
              <Container
                py="2%"
                maxW="95%"
                bg={useColorModeValue("green.300", "green.500")}
                borderRadius="md"
                color={useColorModeValue("gray.700", "gray.200")}
              >
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
                          bg={useColorModeValue("blue.300", "blue.500")}
                          onClick={() =>
                            window.open(`https://ipfs.io/ipfs/${output.cid}`)
                          }
                        >
                          Show on IPFS
                        </Button>
                      </Box>
                    </>
                  ))
                ) : (
                  <Text>No outputs</Text>
                )}
              </Container>
            </AccordionPanel>
          </AccordionItem>

          <AccordionItem>
            <h2>
              <AccordionButton>
                <Box
                  as="span"
                  flex="1"
                  textAlign="left"
                  fontWeight={"semibold"}
                  fontSize={"lg"}
                >
                  Stderr
                </Box>
                <AccordionIcon />
              </AccordionButton>
            </h2>
            <AccordionPanel pb={4}>
              {job?.stderr ? (
                <>
                  <Box
                    display={"flex"}
                    alignItems={"center"}
                    flexDirection={"column"}
                  >
                    <Text mb={2}>
                      <span style={styleDetail}> CID:</span> {job.cidStderr}
                    </Text>
                    <Container
                      py="2%"
                      maxW="95%"
                      bg={useColorModeValue("red.300", "red.500")}
                      borderRadius="md"
                      color="BlackAlpha.900"
                    >
                      <Text>Stderr:{job.stderr}</Text>
                    </Container>
                  </Box>
                </>
              ) : (
                <Text>No stderr</Text>
              )}
            </AccordionPanel>
          </AccordionItem>
          <AccordionItem>
            <h2>
              <AccordionButton>
                <Box
                  as="span"
                  flex="1"
                  textAlign="left"
                  fontWeight={"semibold"}
                  fontSize={"lg"}
                >
                  Stdout
                </Box>
                <AccordionIcon />
              </AccordionButton>
            </h2>
            <AccordionPanel pb={4}>
              {job?.stdout ? (
                <>
                  <Box
                    display={"flex"}
                    alignItems={"center"}
                    flexDirection={"column"}
                  >
                    <Text mb={2}>
                      <span style={styleDetail}> CID:</span> {job.cidStdout}
                    </Text>
                    <Container
                      py="2%"
                      maxW="95%"
                      bg={useColorModeValue("green.300", "green.500")}
                      borderRadius="md"
                      color={useColorModeValue("gray.700", "gray.200")}
                    >
                      <Text>Stderr:{job.stdout}</Text>
                    </Container>
                  </Box>
                </>
              ) : (
                <Text>No stdout</Text>
              )}
            </AccordionPanel>
          </AccordionItem>
        </Accordion>

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
                width="500"
                height="500"
                alt="Image of the output"
                layout="responsive"
              />
              <Text fontWeight={"bold"} mt={2}>
                {job?.outputsNamesAndCids[0]?.name}
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
