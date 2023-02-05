import {
    Box,
    Flex,
    Text,
    useColorModeValue,
    Button,
    Spinner
  } from "@chakra-ui/react";

export default function JobSubmitted() {

    return (
        <Flex direction='column' h='full' w='full' p={8}>
          <Box alignContent='center' justifyContent='center'>
            <Text align='center' fontSize="2xl" fontFamily="monospace" fontWeight="bold">
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
              mt={4} align='center' fontSize="sm" fontFamily="monospace" fontWeight="bold">
                Your job has been submitted. You will be redirected to the job page once it has been processed.
              </Text>
            </Flex>
          </Flex>
        </Flex>
    
      );
    }
    