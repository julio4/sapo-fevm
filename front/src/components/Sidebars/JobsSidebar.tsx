import React, { useEffect } from "react";

import { Category, useJobContext } from "src/components/Context/JobContext";

import { FaBorderAll } from "react-icons/fa";

import {
  Box,
  Flex,
  Icon,
  useColorModeValue,
  Text,
  useDisclosure,
} from "@chakra-ui/react";

import { AiFillCaretRight, AiFillCaretLeft } from "react-icons/ai";

export default function JobsSidebar() {
  const { isOpen, onToggle, onOpen } = useDisclosure();
  const {
    setStep,
    setCategory,
    setJob,
    setAllFiles,
    category,
    allCategories,
    setSeeIncompatibleFiles,
  } = useJobContext();

  useEffect(() => {
    onOpen();
  }, []);

  const allCat = { name: "All", id: 0, icon: FaBorderAll, color: null };

  return (
    <>
      <Box
        flexShrink="0"
        bg={useColorModeValue("whiteAlpha.800", "blackAlpha.600")}
        backdropFilter="blur(20px)"
        borderRight="1px"
        borderRightColor={useColorModeValue("gray.200", "gray.700")}
        w={isOpen ? "60" : "0"}
        h="full"
        dir="column"
        transition={"all 0.2s ease-in-out"}
        transform={isOpen ? "translateX(0)" : "translateX(-100%)"}
        opacity={isOpen ? "1" : "0"}
      >
        <Box
          transition={"opacity 0.1s ease-in-out"}
          opacity={isOpen ? "1" : "0"}
          visibility={isOpen ? "visible" : "hidden"}
        >
          <Flex
            h="20"
            alignItems="center"
            mx="8"
            justifyContent="space-between"
          >
            <Text fontSize="2xl" fontFamily="monospace" fontWeight="bold">
              Category
            </Text>
          </Flex>

          <NavItem
            cat={allCat}
            current={category || allCat}
            onClick={() => {
              setStep(0);
              setCategory(null);
              setJob(null);
              setAllFiles([]);
              setSeeIncompatibleFiles(false);
            }}
          />

          {allCategories.map((cat) => (
            <NavItem
              cat={cat}
              current={category}
              onClick={() => {
                setStep(0);
                setCategory(cat);
                setAllFiles([]);
                setSeeIncompatibleFiles(false);
              }}
            />
          ))}
        </Box>
      </Box>
      <Box
        position="absolute"
        left={isOpen ? "60" : "0"}
        top="45%"
        _hover={{
          color: "teal.400",
        }}
        cursor={"pointer"}
        transition={"all 0.2s ease-in-out"}
        transform={isOpen ? "translateX(-100%)" : "translateX(0)"}
        onClick={onToggle}
      >
        {isOpen ? <AiFillCaretLeft /> : <AiFillCaretRight />}
      </Box>
    </>
  );
}

const NavItem = ({
  cat,
  current,
  onClick,
}: {
  cat: Category;
  current: Category | null;
  onClick: any;
}) => {
  return (
    <Box onClick={onClick}>
      <Flex
        align="center"
        p="4"
        mx="4"
        my="1"
        role="group"
        cursor="pointer"
        transition={"transform 0.1s ease-in-out"}
        border="2px solid transparent"
        borderRadius="xl"
        _hover={{
          bgGradient: useColorModeValue(
            "linear(to-r, teal.600, green.400)",
            "linear(to-r, teal.300, green.100)"
          ),
          bgClip: "text",
          transform: "scale(1.01)",
        }}
        {...(current?.id === cat.id
          ? {
              bgGradient: useColorModeValue(
                "linear(to-r, teal.600, green.400)",
                "linear(to-r, teal.300, green.100)"
              ),
              bgClip: "text",
              borderColor: "teal.400",
            }
          : {})}
      >
        <Icon
          mr="4"
          fontSize="16"
          transition={"transform 0.1s ease-in-out"}
          color={current?.id === cat.id ? "teal.400" : "black.500"}
          _groupHover={{
            color: "teal.400",
          }}
          as={cat.icon}
        />
        {cat.name}
      </Flex>
    </Box>
  );
};
