import { useJobContext } from "@/components/Context/JobContext";
import SelectJob from "@/components/Flow/SelectJob";
import SelectCidInput from "@/components/Flow/SelectCidInput";
import SubmitJob from "@/components/Flow/SubmitJob";
import JobSubmitted from "@/components/Flow/JobSubmitted";

import { Box, Progress, useColorModeValue } from "@chakra-ui/react";

export default function JobRequestFlow() {
  const { step } = useJobContext();

  const Step: JSX.Element = [
    <SelectJob key={0}/>,
    <SelectCidInput key={1}/>,
    <SubmitJob key={2}/>,
    <JobSubmitted key={3}/>,
  ][step];

  return (
    <Box flexGrow={1} p={6}>
      <Box
        h="full"
        w="full"
        bg={useColorModeValue("white", "whiteAlpha.200")}
        backdropFilter="auto"
        backdropBlur="2px"
        border="1px"
        borderColor={useColorModeValue("gray.200", "gray.700")}
        borderRadius="lg"
      >
        <Progress
          colorScheme="green"
          size="xs"
          borderRadius="lg"
          value={step * 34}
        />
        {Step}
      </Box>
    </Box>
  );
}
