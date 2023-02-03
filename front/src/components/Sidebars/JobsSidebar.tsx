import React, { ReactNode } from 'react';

import { Category, useJobContext } from 'src/components/Context/JobContext';

import {
  Box,
  Flex,
  Icon,
  useColorModeValue,
  Text,
  useDisclosure,
} from '@chakra-ui/react';

import {
  FaImages,
  FaCogs,
  FaHdd,
  FaRegObjectGroup,
  FaPen,
} from 'react-icons/fa';
import { AiFillCaretRight, AiFillCaretLeft } from "react-icons/ai";

const ListCatogory: Array<Category> = [
  {
    id: 1,
    name: 'Images processing',
    icon: FaImages
  },
  {
    id: 2,
    name: 'Data generation',
    icon: FaCogs
  },
  {
    id: 3,
    name: 'Artifical Intelligence',
    icon: FaHdd
  },
  {
    id: 4,
    name: 'Simulation',
    icon: FaRegObjectGroup
  },
  {
    id: 5,
    name: 'Custom',
    icon: FaPen
  }
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

          {ListCatogory.map((cat) => (
            <NavItem cat={cat} />
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

const NavItem = ({ cat }: { cat: Category}) => {
  const { setStep, setCategory, category } = useJobContext();
  return (
    <Box onClick={() => {
      setCategory(cat)
      }}>
      <Flex
        align="center"
        p="4"
        mx="4"
        my="1"
        role="group"
        cursor="pointer"
        transition={'transform 0.1s ease-in-out'}
        border='2px solid transparent'
        borderRadius="xl"
        _hover={{
          bgGradient: 'linear(to-r, teal.600, green.400)',
          bgClip: 'text',
          transform: 'scale(1.01)'
        }}
        {...category?.id === cat.id ? {
          bgGradient: 'linear(to-r, teal.600, green.400)',
          bgClip: 'text',
          borderColor: "teal.400"
        }:{}}
        >
        <Icon
            mr="4"
            fontSize="16"
            transition={'transform 0.1s ease-in-out'}
            color={category?.id === cat.id ? 'teal.400' : 'black.500'}
            _groupHover={{
              color: 'teal.400',
            }}
            as={cat.icon}
        />
        {cat.name}
      </Flex>
    </Box>
  );
};