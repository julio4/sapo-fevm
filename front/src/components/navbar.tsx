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
  useColorMode,
  Center,
} from '@chakra-ui/react';
import { MoonIcon, SunIcon } from '@chakra-ui/icons';
import { FaWallet } from 'react-icons/fa';
import { useAccount, useDisconnect } from 'wagmi'

export default function Nav() {
  const { colorMode, toggleColorMode } = useColorMode();

  // Wallet
  const { address, isConnected } = useAccount()
  const { disconnect } = useDisconnect()

  return (
    <>
      <Box
        bgGradient={useColorModeValue('linear(to-r, teal.50, green.50)', 'linear(to-r, teal.900, green.900)')}
        px={4}>

        <Flex h={16} alignItems={'center'} justifyContent={'space-between'}>
          <Box>
            <Heading as='h1' size='lg'>🐸 Sapo</Heading>
          </Box>

          <Flex alignItems={'center'}>
            <Stack direction={'row'} spacing={3}>
              <Button onClick={toggleColorMode}>
                {colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
              </Button>

              <Menu>
                {isConnected ? 
                  <Box>
                    <MenuButton
                      as={Button}
                      color="green.400"
                      cursor={'pointer'}>
                        <FaWallet />
                    </MenuButton>
                    <MenuList
                      alignItems={'center'}
                      bg={useColorModeValue('whiteAlpha.200', 'darkAlpha.200')} 
                      backdropFilter='auto'
                      backdropBlur='2px'
                      >
                      <Center>
                        <Text>Wallet</Text>
                      </Center>
                      <Center>
                        <Text fontSize='md' as='b'>{
                        address?.slice(0,4) + "..." + address?.slice(address.length - 4, address.length)
                        }</Text>
                      </Center>
                      <MenuDivider />
                      <MenuItem
                        onClick={() => disconnect()}
                        bg='inherit'
                        >Disconnect</MenuItem>
                    </MenuList>
                  </Box>
                  : null
                }
              </Menu>
            </Stack>
          </Flex>
        </Flex>
      </Box>
    </>
  );
}
