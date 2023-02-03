import {
  Box,
  useColorModeValue,
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  Td
} from "@chakra-ui/react";

interface File {
  cid: string,
  type: string,
  size: string
}

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
  let files = [
    {
      cid: "QmWXShtJXt6Mw3FH7hVCQvR56xPcaEtSj4YFSGjp2QxA4v",
      type: "img",
      size: "4 MB"
    },
    {
      cid: "QmU7gJi6Bz3jrvbuVfB7zzXStLJrTHf6vWh8ZqkCsTGoRC",
      type: "json",
      size: "154 KB"
    },
    {
      cid: "QmWXShtJXt6Mw3FH7hVCQvR56xPcaEtSj4YFSGjp2QxB4v",
      type: "txt",
      size: "234 KB"
    },
  ]

  //files = []
  return (
    <Box
      flexGrow={1}
      p={6}>
      <Box
        h="full" w="full"
        bg={useColorModeValue('whiteAlpha.700', 'blackAlpha.700')}
        backdropFilter='auto'
        backdropBlur='2px'
        border="1px"
        borderColor={useColorModeValue('gray.200', 'gray.700')}
        borderRadius="lg">

        {
          files.length === 0 ? (
            <AddFile />
          ) : (
            <FilesTable files={files} />
          )
        }

      </Box>
    </Box >
  );
}
