import { useState } from 'react';
import {
  Box,
  Text,
  Heading,
  Flex,
  Button,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  MenuDivider,
  useColorModeValue,
  Stack,
  Center,
} from '@chakra-ui/react';

import { FaWallet } from 'react-icons/fa';
import { useAccount, useDisconnect } from 'wagmi'

import Link from 'next/link'
import { useRouter } from 'next/router'
import ModeToggler from './modeToggler';

export default function Nav() {
  const [connected, setConnected] = useState(false)
  const router = useRouter();

  // Wallet
  const { isConnected, address } = useAccount({
    onConnect: () => {
      setConnected(true)
    },
    onDisconnect: () => {
      setConnected(false)
    }
  })

  const { disconnect } = useDisconnect()

  const buttonBg = useColorModeValue('green.50', 'whiteAlpha.200')
  const buttonBgHover = useColorModeValue('whiteAlpha.600', 'whiteAlpha.300')
  const menuBg = useColorModeValue('whiteAlpha.500', 'darkAlpha.200')
  const itemBgHover = useColorModeValue('blackAlpha.50', 'whiteAlpha.300')

  return (
    <>
      <Box
        bgGradient={useColorModeValue('linear(to-r, teal.100, green.100)', 'linear(to-r, teal.800, green.800)')}
        borderBottomWidth={1}
        borderStyle={'solid'}
        borderColor={useColorModeValue('gray.200', 'gray.900')}
        px={4}>

        <Flex h={16} alignItems={'center'} justifyContent={'space-between'}>

          <Box>
            <Link href='/app' replace>
              <Heading as='h1' fontFamily='system-ui' size='xl' ml={2}>🐸 Sapo</Heading>
            </Link>
          </Box>

          <Flex alignItems={'center'}>
            <Stack direction={'row'} spacing={3}>
              <ModeToggler />

              <Menu>
                {connected ?
                  <Box>
                    <MenuButton
                      as={Button}
                      color="green.400"
                      bg={buttonBg}
                      _hover={{
                        bg: buttonBgHover,
                      }}
                      cursor={'pointer'}>
                      <FaWallet />
                    </MenuButton>
                    <MenuList
                      alignItems={'center'}
                      bg={menuBg}
                      backdropFilter='auto'
                      backdropBlur='4px'
                    >
                      <Center>
                        <Text>Wallet</Text>
                      </Center>
                      <Center>
                        <Text fontSize='md' as='b'>{
                          address?.slice(0, 4) + "..." + address?.slice(address.length - 4, address.length)
                        }</Text>
                      </Center>
                      <MenuDivider />
                      <MenuItem
                        _hover={
                          { bg: itemBgHover }
                        }
                        onClick={() => router.push('/app/jobsHistory')}
                        bg='inherit'>
                        Jobs history
                      </MenuItem>
                      <MenuItem
                        _hover={
                          { bg: itemBgHover }
                        }
                        onClick={() => disconnect()}
                        bg='inherit'>
                        Disconnect
                      </MenuItem>
                    </MenuList>
                  </Box>
                  : <></>
                }
              </Menu>
            </Stack>
          </Flex>
        </Flex>
      </Box>
    </>
  );
}
