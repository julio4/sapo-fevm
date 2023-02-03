import {
  Box,
  chakra,
  Flex,
  Stack,
  Text,
  useColorModeValue,
  VisuallyHidden,
} from '@chakra-ui/react';
import { FaGithub, FaEthereum } from 'react-icons/fa';

import { ReactNode } from 'react';

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
      bgGradient={useColorModeValue('linear(to-r, teal.50, green.50)', 'linear(to-r, teal.900, green.900)')}
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
          <Text
            color={useColorModeValue('gray.600', 'gray.100')}
            fontSize='sm' pl='20px'>© 2023 KS. All rights reserved</Text>
          <Stack direction={'row'} spacing={6} pr="20px">
            <SocialButton label={'Github'} href={'#'}>
              <FaGithub />
            </SocialButton>
            <SocialButton label={'EthGlobal'} href={'#'}>
              <FaEthereum />
            </SocialButton>
          </Stack>
        </Flex>
      </Box>
    </Box>
  );
}
 