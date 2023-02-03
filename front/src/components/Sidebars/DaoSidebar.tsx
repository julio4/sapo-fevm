import React, { ReactNode } from 'react';

import {
  Box,
  Flex,
  Icon,
  useColorModeValue,
  Link,
  Text,
  useDisclosure,
  FlexProps,
} from '@chakra-ui/react';

import { AiFillCaretRight, AiFillCaretLeft } from "react-icons/ai";

import { IconType } from 'react-icons';

export default function DaoSidebar() {
  const { isOpen, onToggle } = useDisclosure();
  return (
    <>
        <Box
            bg={useColorModeValue('whiteAlpha.800', 'blackAlpha.600')}
            backdropFilter='auto'
            backdropBlur='2px'
            borderRight="1px"
            borderLeft="1px"
            borderLeftColor={useColorModeValue('gray.200', 'gray.700')}
            h="full"
            dir='column'
            transition={'all 0.2s ease-in-out'}
            //transform={isOpen ? 'translateX(0)' : 'translateX(-100%)'}
            opacity={isOpen ? '1' : '0'}
            w={isOpen ? '80' : '0'}
            >

            <Flex h="20" alignItems="center" mx="8" justifyContent="space-between">
                <Text fontSize="2xl" fontFamily="monospace" fontWeight="bold">
                Dao source
                </Text>
            </Flex>

        </Box>
        <Box 
        position='absolute'
        right={isOpen ? '80' : '0'}
        top="45%"
        _hover={{
          color: 'teal.400',
        }}
        cursor={'pointer'}
        transition={'all 0.2s ease-in-out'}
        transform={isOpen ? 'translateX(100%)' : 'translateX(0)'}
        onClick={onToggle}>
            {isOpen ? <AiFillCaretRight /> : <AiFillCaretLeft />}
        </Box>
    </>
  );
}

interface NavItemProps extends FlexProps {
  icon: IconType;
  children: ReactNode;
}

const NavItem = ({ icon, children, ...rest }: NavItemProps) => {
  return (
    <Link href="#" style={{ textDecoration: 'none' }} _focus={{ boxShadow: 'none' }}>
      <Flex
        align="center"
        p="4"
        mx="4"
        my="1"
        borderRadius="lg"
        role="group"
        cursor="pointer"
        transition={'transform 0.1s ease-in-out'}
        _hover={{
          bgGradient: 'linear(to-r, teal.600, green.400)',
          bgClip: 'text',
          transform: 'scale(1.01)'
        }}
        {...rest}>
        {icon && (
          <Icon
            mr="4"
            fontSize="16"
            _groupHover={{
              color: 'teal.400',
            }}
            as={icon}
          />
        )}
        {children}
      </Flex>
    </Link>
  );
};