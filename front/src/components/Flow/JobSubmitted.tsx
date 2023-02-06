import {
  Box,
  Flex,
  Text,
  useColorModeValue,
  Button,
  Spinner
} from "@chakra-ui/react";
import { useRouter } from 'next/router'

export default function JobSubmitted() {
  const router = useRouter()

  return (
    <Flex direction='column' h='full' w='full' p={8}>
      <Box alignContent='center' justifyContent='center'>
        <Text align='center' fontSize="2xl" fontFamily="system-ui" fontWeight="bold">
          Job Submitted
        </Text>
      </Box>

      <Flex alignContent='center' justifyContent='center' h='full'>
        <Flex direction='column' alignContent='center' justifyContent='center'>
          <Spinner
            alignSelf='center'
            thickness='4px'
            speed='0.65s'
            emptyColor='gray.200'
            color='green.500'
            size='xl'
          />
          <Text color={useColorModeValue('gray.600', 'gray.400')}
            my={4} align='center' fontSize="md" fontFamily="system-ui" fontWeight="bold">
            Your job has been submitted. Job can take up to several minutes to be completed...
          </Text>
          <Button onClick={() => router.push('/app/jobsHistory')}>Go to job history</Button>
        </Flex>
      </Flex>
    </Flex>
  );
}
