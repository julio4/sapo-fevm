import { Box, useColorModeValue } from "@chakra-ui/react";

export default function FileSelector() {
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

          </Box>
    </Box>
  );
}
