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

const JobsList = ({
  jobAddresses,
  selected,
  onSelect,
}: {
  jobAddresses: readonly `0x${string}`[];
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

      <Divider />

      {Object.entries(jobAddresses).map(([key, jobAddress]) => <JobInstance {...{ onSelect, jobAddress, selected, key }} />)}
    </Box>
  );
};

export default JobsList;