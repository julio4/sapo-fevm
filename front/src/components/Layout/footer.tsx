import {
  Box,
  Tooltip,
  chakra,
  Flex,
  Stack,
  Text,
  useColorModeValue,
  VisuallyHidden,
} from '@chakra-ui/react';
import { FaGithub, FaEthereum } from 'react-icons/fa';

import { ReactNode } from 'react';

import { useJobContext } from '../Context/JobContext';
import { useIpfsContext } from '@/services/ipfs';

const SocialButton = ({
  children,
  label,
  href,
}: {
  children: ReactNode;
  label: string;
  href: string;
}) => {
  return (
    <chakra.button
      bg={useColorModeValue('blackAlpha.100', 'whiteAlpha.100')}
      rounded={'full'}
      w={8}
      h={8}
      cursor={'pointer'}
      as={'a'}
      href={href}
      target={'_blank'}
      display={'inline-flex'}
      alignItems={'center'}
      justifyContent={'center'}
      transition={'background 0.3s ease'}
      _hover={{
        bg: useColorModeValue('blackAlpha.200', 'whiteAlpha.200'),
      }}>
      <VisuallyHidden>{label}</VisuallyHidden>
      {children}
    </chakra.button>
  );
};

export default function Footer() {
 const { id, version, isOnline } = useIpfsContext();
  return (
    <Box
      bgGradient={useColorModeValue('linear(to-r, teal.100, green.100)', 'linear(to-r, teal.900, green.900)')}
      color={useColorModeValue('gray.700', 'gray.200')}
      width="100%"
      bottom="0"
      zIndex={10}
      >
      <Box
        borderTopWidth={1}
        borderStyle={'solid'}
        borderColor={useColorModeValue('gray.200', 'gray.900')}>
        <Flex
          minH={'60px'}
          direction={{ base: 'column', md: 'row' }}
          justify={{ base: 'center', md: 'space-between' }}
          align={{ base: 'center', md: 'center' }}>


          {/* Debug */}
          <Tooltip label={isOnline 
            ? `Connected to local ipfs node ${id}` 
            : 'Error, can\'t create local ipfs node'}>
            <Flex gap={3} alignItems='center'>
              <Box
                as="div"
                h="10px"
                w="10px"
                ml={4}
                position="relative"
                bgColor={isOnline ? 'green.500' : 'red.500'}
                borderRadius="50%"
                _before={{
                  content: "''",
                  position: 'relative',
                  display: 'block',
                  width: '300%',
                  height: '300%',
                  boxSizing: 'border-box',
                  marginLeft: '-100%',
                  marginTop: '-100%',
                  borderRadius: '50%',
                }}
              />            
              <Text fontSize='sm'>v{version}</Text>
            </Flex>
          </Tooltip>

          <Text
            color={useColorModeValue('gray.600', 'gray.100')}
            fontSize='sm' pl='20px'>© 2023 KS. All rights reserved</Text>

          <Stack direction={'row'} spacing={6} pr="20px">
            <SocialButton label={'Github'} href={'https://github.com/julio4/sapo-fevm/'}>
              <FaGithub />
            </SocialButton>
            <SocialButton label={'EthGlobal'} href={'https://ethglobal.com/showcase/sapo-wtdhn'}>
              <FaEthereum />
            </SocialButton>
          </Stack>
        </Flex>
      </Box>
    </Box>
  );
}
 