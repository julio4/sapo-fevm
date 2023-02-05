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
  return (
    <Box
      bgGradient={useColorModeValue('linear(to-r, teal.100, green.100)', 'linear(to-r, teal.800, green.800)')}
      color={useColorModeValue('gray.700', 'gray.200')}
      width="100%"
      bottom="0"
      zIndex={10}
      >
      <Box
        borderTopWidth={1}
        borderStyle={'solid'}
        borderColor={useColorModeValue('gray.200', 'gray.700')}>
        <Flex
          minH={'60px'}
          direction={{ base: 'column', md: 'row' }}
          justify={{ base: 'center', md: 'space-between' }}
          align={{ base: 'center', md: 'center' }}>

          <Text
            cursor='select'
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
 