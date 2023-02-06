import {
  Box,
  Flex,
  Text,
  useColorModeValue,
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  Td,
  Button,
  Heading,
  Checkbox,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Code,
} from "@chakra-ui/react";

import {
  ChevronDownIcon,
  CheckIcon,
  WarningTwoIcon,
  WarningIcon,
  QuestionIcon,
} from "@chakra-ui/icons";

import SelectFromDaoInput from "@/components/Flow/SelectFromDaoInput";
import SelectFromIpfsInput from "@/components/Flow/SelectFromIpfsInput";
import { useJobContext, File } from "@/components/Context/JobContext";
import { useEffect, useState } from "react";
import { useDisclosure } from "@chakra-ui/react";

const NoFileFound = () => {
  return (
    <Flex
      direction="column"
      justifyContent="center"
      h="full"
      p={12}
      rounded={6}
      alignItems="center"
    >
      <Heading
        as="h1"
        size="2xl"
        mb={2}
        color={useColorModeValue("gray.300", "gray.700")}
      >
        No files found
      </Heading>
      <Heading
        as="h2"
        size="md"
        color={useColorModeValue("gray.300", "gray.700")}
      >
        Use import buttons to add your files
      </Heading>
    </Flex>
  );
};

const FilesTable = ({ files }) => {
  const {
    job,
    jobRequest,
    setJobRequest,
    setStep,
    seeIncompatibleFiles,
    setSeeIncompatibleFiles,
  } = useJobContext();
  const [showWarning, setShowWarning] = useState(true);
  const [warningTitle, setWarningTitle] = useState("");
  const [warningMessage, setWarningMessage] = useState("");
  const { isOpen, onOpen, onClose } = useDisclosure();

  const getOpacity = (file, job) => {
    if (file.cid.includes("ERROR")) {
      return 0.4;
    } else if (
      !(
        job.inputTypes.includes("all") ||
        job.inputTypes.some((ft) => file.type.includes(ft))
      )
    ) {
      return 0.4;
    }
    return 1;
  };

  const sortFiles = (files) => {
    return files.sort((a, b) => {
      if (a.cid.includes("ERROR")) {
        return 1;
      } else if (b.cid.includes("ERROR")) {
        return -1;
      }
      if (!a.type.includes(job.inputTypes)) {
        return 1;
      } else if (!b.type.includes(job.inputTypes)) {
        return -1;
      }
      return 0;
    });
  };

  function humanFileSize(size) {
    var i = size == 0 ? 0 : Math.floor(Math.log(size) / Math.log(1024));
    return (
      (size / Math.pow(1024, i)).toFixed(2) * 1 +
      " " +
      ["B", "kB", "MB", "GB", "TB"][i]
    );
  }

  const handleClick = (file) => {
    if (!job) return;
    if (showWarning) {
      if (file.cid.includes("ERROR")) {
        setWarningTitle("Warning !");
        setWarningMessage(
          "This file is not reachable, it may have been deleted or moved"
        );
        onOpen();
        return;
      } else if (
        !(
          job.inputTypes.includes("all") ||
          job.inputTypes.some((ft) => file.type.includes(ft))
        )
      ) {
        setWarningTitle("Warning !");
        setWarningMessage(
          "This file is not compatible with the job, please select another file"
        );
        onOpen();
        return;
      }
    }
    setJobRequest({
      ...jobRequest,
      job: job,
      usrInput: file,
    });
    setStep(2);
  };

  const fileBg = useColorModeValue("green.100", "green.900");

  return (
    <TableContainer>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader
            textAlign={"center"}
            color={
              warningMessage.includes("reachable") ? "red.500" : "yellow.500"
            }
            fontWeight={"semibold"}
          >
            {warningTitle}
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text>{warningMessage}</Text>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="teal" mr={3} onClick={onClose}>
              Close
            </Button>
            <Button
              variant="ghost"
              onClick={() => {
                setShowWarning(false);
                onClose();
              }}
            >
              Do not show me again
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
      <Box overflowY="auto" maxHeight="50vh">
        <Table size="sm">
          <Thead>
            <Tr>
              <Th />
              <Th>CID</Th>
              <Th>File type</Th>
              <Th>Size</Th>
            </Tr>
          </Thead>
          <Tbody>
            {sortFiles(files)
              .filter((file) => {
                if (seeIncompatibleFiles) {
                  return true;
                } else if (file.cid.includes("ERROR")) {
                  return false;
                }
                return (
                  job?.inputTypes.some((ft) => file.type?.includes(ft)) ||
                  job?.inputTypes.includes("all")
                );
              })
              .map((file) => (
                <Box
                  key={file.cid}
                  transition={"all 0.2s ease-in-out"}
                  as="tr"
                  borderRadius="md"
                  _hover={{
                    cursor: "pointer",
                    bg: fileBg,
                  }}
                  onClick={() => handleClick(file)}
                >
                  <Td
                    _hover={{ transform: "scale(1.01)" }}
                    transition={"all 0.2s ease-in-out"}
                    opacity={getOpacity(file, job) + 0.2}
                  >
                    {file.cid.includes("ERROR") ? (
                      <WarningTwoIcon size="20px" color={"red.500"} />
                    ) : job.inputTypes.includes("all") ||
                      job.inputTypes.some((ft) => file.type.includes(ft)) ? (
                      <CheckIcon size="20px" color={"green.500"} />
                    ) : (
                      <WarningIcon size="20px" color={"yellow.500"} />
                    )}
                  </Td>
                  <Td
                    _hover={{ transform: "scale(1.01)" }}
                    transition={"all 0.2s ease-in-out"}
                    opacity={getOpacity(file, job)}
                  >
                    {file.cid}
                  </Td>
                  <Td
                    _hover={{ transform: "scale(1.01)" }}
                    transition={"all 0.2s ease-in-out"}
                    opacity={getOpacity(file, job)}
                  >
                    {file.type ? file.type : "Unknown"}
                  </Td>
                  <Td
                    _hover={{ transform: "scale(1.01)" }}
                    transition={"all 0.2s ease-in-out"}
                    opacity={getOpacity(file, job)}
                  >
                    {file.size ? humanFileSize(file.size) : "Not Specified"}
                  </Td>
                </Box>
              ))}
          </Tbody>
        </Table>
      </Box>
    </TableContainer>
  );
};

export default function SelectCidInput() {
  const { allFiles } = useJobContext();
  const [sourceInput, setSourceInput] = useState("");
  const { seeIncompatibleFiles, setSeeIncompatibleFiles } = useJobContext();
  const [menuSelected, setMenuSelected] = useState("");
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <Flex direction="column" h="full" w="full" p={8}>
      <Modal isOpen={isOpen} onClose={onClose} size={"xl"}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader
            textAlign={"center"}
            fontWeight={"bold"}
            fontSize={"2xl"}
          >
            {"How to import your CID"}
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text mb={5}>
              <Heading size={"md"} mb={1}>
                {"From your DataDAO"}
              </Heading>
              <Text mb={2}>
                CIDs should be stored as bytes in your Smart Contract. <br />
                You should have a public view function named{" "}
                <Code>getAllCids</Code> that returns an array of all your cids.
              </Text>
              <Code p={3} borderRadius={6} mb={2}>
                {" "}
                function getAllCids() public view returns (bytes[] memory ){
                  "{"
                }{" "}
                <br />
                return cids;
                {"}"}{" "}
              </Code>
              <Text mb={3}>
                You can use our DataDAO example address to test:
                <Code borderRadius={2} padding={1} px={2}>
                  0xC9605dAe1a25598d3FC548336693aBedEe77a58A
                </Code>
              </Text>
              <Text fontSize="md" fontStyle={"italic"}>
                Possibility to import with strings is coming soon
              </Text>
            </Text>
            <Text mb={4}>
              <Heading size={"sm"}>{"From a single IPFS CID"}</Heading>
              <Text>
                When importing a single CID, user simply needs to enter the CID
                of the file.
              </Text>
            </Text>
            <Text>
              <Heading size={"sm"} mb={1}>
                {"More information"}
              </Heading>
              <Text>
                A file is designated as inaccessible when no response has been
                received after a 5 second waiting period. If you are confident
                that the file is present on the IPFS network, you can attempt to
                reload it.
                <br />
                Check the box to view files that are incompatible with the
                selected job.
              </Text>
            </Text>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="teal" mr={3} onClick={onClose}>
              Close
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
      <Box alignContent="center" justifyContent="center">
        <Text
          align="center"
          fontSize="2xl"
          fontFamily="system-ui"
          fontWeight="bold"
        >
          Select input file
        </Text>
      </Box>

      <Box alignContent="center" justifyContent="center" mb={4}>
        <Menu>
          <MenuButton
            bg={useColorModeValue("green.100", "teal.700")}
            as={Button}
            rightIcon={<ChevronDownIcon />}
          >
            Import from {menuSelected}
          </MenuButton>
          <MenuList>
            <MenuItem
              onClick={() => {
                setMenuSelected("IPFS");
                setSourceInput("ipfs");
              }}
            >
              IPFS
            </MenuItem>
            <MenuItem
              onClick={() => {
                setMenuSelected("DataDAO");
                setSourceInput("datadao");
              }}
            >
              DataDAO
            </MenuItem>
          </MenuList>
        </Menu>

        <Button
          ml={6}
          borderRadius={10}
          bg={useColorModeValue("teal.50", "whiteAlpha.400")}
          pt={1}
          onClick={() => onOpen()}
        >
          How to import{" "}
          <QuestionIcon size="20px" color={"blue.400"} ml={3} mb={1} />
        </Button>

        {sourceInput == "datadao" ? <SelectFromDaoInput /> : <div></div>}
        {sourceInput == "ipfs" ? <SelectFromIpfsInput /> : <div></div>}
      </Box>

      <Checkbox
        size={"sm"}
        pb={3}
        onChange={() => {
          setSeeIncompatibleFiles(!seeIncompatibleFiles);
        }}
      >
        Load incompatible files
      </Checkbox>

      <Box
        h="full"
        w="full"
        bg={useColorModeValue("whiteAlpha.700", "blackAlpha.400")}
        backdropFilter="auto"
        backdropBlur="2px"
        border="1px"
        borderColor={useColorModeValue("gray.200", "gray.700")}
        borderRadius="lg"
      >
        {allFiles.length === 0 ? (
          <NoFileFound />
        ) : (
          <FilesTable files={allFiles} />
        )}
      </Box>
    </Flex>
  );
}
