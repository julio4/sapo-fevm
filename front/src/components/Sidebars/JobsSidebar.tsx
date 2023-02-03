import React, { ReactNode } from 'react';

import {
  IconButton,
  Box,
  CloseButton,
  Flex,
  Icon,
  useColorModeValue,
  Link,
  Drawer,
  DrawerContent,
  Text,
  useDisclosure,
  BoxProps,
  FlexProps,
} from '@chakra-ui/react';

import {
  FaImages,
  FaCogs,
  FaHdd,
  FaRegObjectGroup,
  FaPen,
} from 'react-icons/fa';
import { AiFillCaretRight, AiFillCaretLeft } from "react-icons/ai";

import { IconType } from 'react-icons';

interface LinkItemProps {
  name: string;
  icon: IconType;
}

const LinkItems: Array<LinkItemProps> = [
  { name: 'Images processing', icon: FaImages },
  { name: 'Data generation', icon: FaCogs },
  { name: 'Artifical Intelligence', icon: FaHdd },
  { name: 'Simulation', icon: FaRegObjectGroup },
  { name: 'Custom', icon: FaPen },
];

export default function JobsSidebar() {
  const { isOpen, onToggle } = useDisclosure();
  return (
    <>
      <Box
        flexShrink='0'
        bg={useColorModeValue('whiteAlpha.800', 'blackAlpha.600')}
        backdropFilter='blur(20px)'
        borderRight="1px"
        borderRightColor={useColorModeValue('gray.200', 'gray.700')}
        w={isOpen ? '60' : '0'}
        h="full"
        dir='column'
        transition={'all 0.2s ease-in-out'}
        transform={isOpen ? 'translateX(0)' : 'translateX(-100%)'}
        opacity={isOpen ? '1' : '0'}
      >

        <Box
          transition={'opacity 0.1s ease-in-out'}
          opacity={isOpen ? '1' : '0'}
          visibility={isOpen ? 'visible' : 'hidden'}
        >
          <Flex h="20" alignItems="center" mx="8" justifyContent="space-between">
            <Text fontSize="2xl" fontFamily="monospace" fontWeight="bold">
              Jobs
            </Text>
          </Flex>

          {LinkItems.map((link) => (
            <NavItem key={link.name} icon={link.icon}>
              {link.name}
            </NavItem>
          ))}
        </Box>
      </Box>
      <Box
        position='absolute'
        left={isOpen ? '60' : '0'}
        top="45%"
        _hover={{
          color: 'teal.400',
        }}
        cursor={'pointer'}
        transition={'all 0.2s ease-in-out'}
        transform={isOpen ? 'translateX(-100%)' : 'translateX(0)'}
        onClick={onToggle}>
        {isOpen ? <AiFillCaretLeft /> : <AiFillCaretRight />}
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