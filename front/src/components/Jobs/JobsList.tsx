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

import { useRouter } from "next/router";
import JobResult from "./JobResult";
import JobSummary from "./JobSummary";
import { CloseIcon } from "@chakra-ui/icons";

const JobsList = ({
  jobAddresses,
  selected,
  onSelect,
}: {
  jobAddresses: readonly `0x${string}`[];
  onSelect: (job: JobSummary) => void;
  selected: `0x${string}` | null;
}) => {
  const router = useRouter();
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
        <Heading
          size="md"
          mb={2}
          display={"flex"}
          justifyContent={"space-between"}
          px={6}
        >
          <Text>Jobs History</Text>
          <CloseIcon
            onClick={() => {
              router.push("/app/");
            }}
            size="20px"
            _hover={{
              cursor: "pointer",
            }}
            color={"teal.500"}
          />
        </Heading>
      </Box>

      <Divider />

      {Object.entries(jobAddresses).map((tab) => {
        let [key, jobAddress] = tab;
        return (
          <JobInstance
            {...{ onSelect, jobAddress, selected, id: parseInt(key) }}
            key={key}
          />
        );
      })}
    </Box>
  );
};

export default JobsList;
