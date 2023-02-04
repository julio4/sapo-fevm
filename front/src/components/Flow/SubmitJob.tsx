import {
  Box,
  Flex,
  Text,
  useColorModeValue,
  Button,
  Input,
  SimpleGrid,
  FormControl,
  FormLabel,
  FormHelperText,
} from "@chakra-ui/react";

import { useJobContext, } from "../Context/JobContext";

const SpecCard = ({ title, value, desc }: { title: string, value: string, desc: string }) => (
  <FormControl>
    <FormLabel>{title}</FormLabel>
    <Input
      placeholder="ERROR"
      value={value}
      bg={useColorModeValue('gray.100', 'gray.800')}
      border={0}
      color={useColorModeValue('gray.500', 'gray.200')}
      _placeholder={{
        color: 'gray.500',
      }}
      _disabled={{
        bg: useColorModeValue('gray.100', 'gray.800'),
        color: useColorModeValue('gray.500', 'gray.200'),
        cursor: 'not-allowed'
      }}
      disabled
    />
    <FormHelperText>{desc}</FormHelperText>
  </FormControl>
)

export default function SubmitJob() {
  const { step, setStep, jobRequest, setJobRequest } = useJobContext();
  return (
    <Flex direction='column' h='full' w='full' p={8}>
      <Box alignContent='center' justifyContent='center' mb={4}>
        <Text align='center' fontSize="2xl" fontFamily="monospace" fontWeight="bold">
          Submit Job
        </Text>
      </Box>

      <Flex
        as={'form'}
        flexDir='column'
        w="full"
        bg={useColorModeValue('whiteAlpha.700', 'transparent')}
        backdropFilter='auto'
        backdropBlur='2px'
        border="1px"
        borderColor={useColorModeValue('gray.200', 'gray.700')}
        borderRadius="lg"
      >
        <SimpleGrid columns={{xl: 2, lg: 1}} spacing={8} p={4}>

          <SpecCard 
            title="Job name" 
            value={jobRequest.job?.name || (jobRequest.custom && "Custom") || ""} 
            desc={(jobRequest.custom && "Your custom job") || "TODO Job desc"} />

          <SpecCard 
            title="Docker image" 
            value={jobRequest.job?.image || ""} 
            desc="The docker image used for the job execution" />

          <SpecCard 
            title="Input file" 
            value={jobRequest.usrInput?.cid || ""} 
            desc="Source file from IPFS/Filecoin CID" />

          <SpecCard 
            title="Approximate price" 
            value='TODO ~0.2 FIL'
            desc="TODO: Job complexity + Input Size + Retry)?" />
        </SimpleGrid>

        <Box p={4}>
          <Button
            fontFamily={'heading'}
            mt={8}
            w={'full'}
            bgGradient={useColorModeValue(
              'linear(to-r, green.300, teal.300)',
              'linear(to-r, green.500, teal.500)')}
            transition={'all 0.3s ease-in-out'}
            color={'white'}
            _hover={{
              opacity: '0.9',
              boxShadow: 'md',
              transform: 'scale(1.005)'
            }}
            onClick={() => {
              setStep(3);
            }}>
            Submit
          </Button>
        </Box>
      </Flex>
    </Flex>

  );
}
