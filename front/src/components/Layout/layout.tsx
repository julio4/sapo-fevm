import Navbar from './navbar'
import Footer from './footer'
import Connect from './connect'
import { ReactNode } from 'react';

import { Flex, useColorModeValue } from '@chakra-ui/react'

import { useAccount } from 'wagmi'

export default function Layout({ children }: { children: ReactNode }) {
  const { isConnected } = useAccount()
  return (
    <Flex flexDir='column' h='100vh'
    bg={useColorModeValue('blackAlpha.50', 'blackAlpha.600')}
    >
      <Navbar />
      {isConnected ?
        <Flex grow={1}>{children}</Flex>
        : <Connect />
      }
      <Footer />
    </Flex>
  )
}
