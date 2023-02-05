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
} from "@chakra-ui/react";

import {
  ChevronDownIcon,
  CheckIcon,
  WarningTwoIcon,
  WarningIcon,
} from "@chakra-ui/icons";

import SelectFromDaoInput from "@/components/Flow/SelectFromDaoInput";
import SelectFromIpfsInput from "@/components/Flow/SelectFromIpfsInput";
import { useJobContext, File } from "@/components/Context/JobContext";
import { useEffect, useState } from "react";

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
  const [showWarning, setShowWarning] = useState(false);

  const getOpacity = (file, job) => {
    if (file.cid.includes("ERROR")) {
      return 0.4;
    } else if (!file.type.includes(job.inputTypes)) {
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

  return (
    <TableContainer>
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
              }
              return file.type?.includes(job?.inputTypes);
            })
            .map((file) => (
              <Box
                key={file.cid}
                transition={"all 0.2s ease-in-out"}
                as="tr"
                borderRadius="md"
                _hover={{
                  cursor: "pointer",
                  bg: useColorModeValue("green.100", "green.900"),
                }}
                onClick={() => {
                  if (!job) return;
                  if (file.cid.includes("ERROR")) {
                    if (!showWarning) {
                      alert("Warning, this file is not reachable");
                      setShowWarning(true);
                      return;
                    }
                  }
                  if (!file.type?.includes(job?.inputTypes)) {
                    if (!showWarning) {
                      alert(
                        "Warning, this file is not compatible with the job"
                      );
                      setShowWarning(true);
                      return;
                    }
                  }
                  setJobRequest({
                    ...jobRequest,
                    job: job,
                    usrInput: file,
                  });
                  setStep(2);
                }}
              >
                <Td
                  _hover={{ transform: "scale(1.01)" }}
                  transition={"all 0.2s ease-in-out"}
                  opacity={getOpacity(file, job) + 0.2}
                >
                  {file.cid.includes("ERROR") ? (
                    <WarningTwoIcon size="20px" color={"red.500"} />
                  ) : !file.type.includes(job.inputTypes) ? (
                    <WarningIcon size="20px" color={"yellow.500"} />
                  ) : (
                    <CheckIcon size="20px" color={"green.500"} />
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
    </TableContainer>
  );
};

export default function SelectCidInput() {
  const { allFiles } = useJobContext();
  const [sourceInput, setSourceInput] = useState("");
  const { seeIncompatibleFiles, setSeeIncompatibleFiles } = useJobContext();

  return (
    <Flex direction="column" h="full" w="full" p={8}>
      <Box alignContent="center" justifyContent="center">
        <Text
          align="center"
          fontSize="2xl"
          fontFamily="monospace"
          fontWeight="bold"
        >
          Select input file
        </Text>
      </Box>

      <Box alignContent="center" justifyContent="center" mb={4}>
        <Menu>
          <MenuButton as={Button} rightIcon={<ChevronDownIcon />}>
            Import from
          </MenuButton>
          <MenuList>
            <MenuItem onClick={() => setSourceInput("ipfs")}>IPFS</MenuItem>
            <MenuItem onClick={() => setSourceInput("datadao")}>
              DataDAO
            </MenuItem>
          </MenuList>
        </Menu>

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
        bg={useColorModeValue("whiteAlpha.700", "blackAlpha.700")}
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
