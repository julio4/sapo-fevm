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
  Heading
} from "@chakra-ui/react";

import { useJobContext, File } from '@/components/Context/JobContext';

const NoFileFound = () => {
  return (
    <Flex direction="column" justifyContent='center' h='full' p={12} rounded={6} alignItems="center">
        <Heading as="h1" size="2xl" mb={2} color={useColorModeValue('gray.300', 'gray.700')}>
          No files found
        </Heading>
        <Heading as="h2" size="md" color={useColorModeValue('gray.300', 'gray.700')}>
          Use import buttons to add your files
        </Heading>
    </Flex>
  )
}

const FilesTable = ({ files }: { files: File[] }) => {
  const { job, jobRequest, setJobRequest, setStep } = useJobContext();
  return (
    <TableContainer>
      <Table size='sm'>
        <Thead>
          <Tr>
            <Th>CID</Th>
            <Th>File type</Th>
            <Th>Size</Th>
          </Tr>
        </Thead>
        <Tbody>
          {
            files.map(file => (
              <Box
                key={file.cid}
                transition={'all 0.2s ease-in-out'}
                as='tr'
                borderRadius="md"
                _hover={{
                  cursor: 'pointer',
                  bg: useColorModeValue('green.100', 'green.900'),
                }}
                onClick={() => {
                  if (!job) return;
                  setJobRequest({
                    ...jobRequest,
                    job: job,
                    usrInput: file
                  });
                  setStep(2);
                }}>
                <Td _hover={{ transform: 'scale(1.01)' }} transition={'all 0.2s ease-in-out'}>{file.cid}</Td>
                <Td _hover={{ transform: 'scale(1.01)' }} transition={'all 0.2s ease-in-out'}>{file.type}</Td>
                <Td _hover={{ transform: 'scale(1.01)' }} transition={'all 0.2s ease-in-out'}>{file.size}</Td>
              </Box>
            ))
          }
        </Tbody>
      </Table>
    </TableContainer>
  )
}

export default function SelectCidInput() {
  const { allFiles } = useJobContext();

  return (
    <Flex direction='column' h='full' w='full' p={8}>
      <Box alignContent='center' justifyContent='center'>
        <Text align='center' fontSize="2xl" fontFamily="monospace" fontWeight="bold">
          Select input file
        </Text>
        
      </Box>
      <Box alignContent='center' justifyContent='center' mb={4}>

        <Text>Import CID:</Text>
        <Button onClick={() => alert("Unimplemented")} mr={2}>
          from IPFS CID
        </Button>
        <Button onClick={() => alert("Unimplemented")}>
          from DataDao
        </Button>
      </Box>

      <Box
        h="full" w="full"
        bg={useColorModeValue('whiteAlpha.700', 'blackAlpha.700')}
        backdropFilter='auto'
        backdropBlur='2px'
        border="1px"
        borderColor={useColorModeValue('gray.200', 'gray.700')}
        borderRadius="lg">

        {
          allFiles.length === 0 ? (
            <NoFileFound />
          ) : (
            <FilesTable files={allFiles} />
          )
        }

      </Box>
    </Flex>

  );
}
