import {
  Text,
  Flex,
  Box,
  useColorModeValue,
  Heading,
  Divider,
  Badge,
} from "@chakra-ui/react";
import JobInstance from "./JobInstance";

import JobResult from "./JobResult";
import JobSummary from "./JobSummary";

const JobsList = ({
  jobAddresses,
  selected,
  onSelect,
}: {
  jobAddresses: readonly `0x${string}`[];
  onSelect: (job: JobSummary) => void;
  selected: `0x${string}` | null;
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

      <Divider />

      {Object.entries(jobAddresses).map((tab) => {
        let [key, jobAddress] = tab;
        return <JobInstance {...{ onSelect, jobAddress, selected, id: parseInt(key) }} key={key} />
      })}
    </Box>
  );
};

export default JobsList;