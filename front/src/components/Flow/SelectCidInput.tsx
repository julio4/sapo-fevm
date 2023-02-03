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
  Button
} from "@chakra-ui/react";

import { useJobContext, File } from '@/components/Context/JobContext';

const AddFile = () => {
  return (
    <></>
  )
}

const FilesTable = ({ files }: { files: File[] }) => {
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
                  bg: useColorModeValue('whiteAlpha.600', 'blackAlpha.300')
                }}
                onClick={() => console.log('clicked cid ')}>
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
      <Box alignContent='center' justifyContent='center' mb={4}>
        <Text align='center' fontSize="2xl" fontFamily="monospace" fontWeight="bold">
          Select input file
        </Text>
        <Text>Import CID:</Text>
        <Button>
          from IPFS CID
        </Button>
        <Button>
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
            <AddFile />
          ) : (
            <FilesTable files={allFiles} />
          )
        }

      </Box>
    </Flex>

  );
}
