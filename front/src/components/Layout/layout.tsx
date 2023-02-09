import { useState, useEffect } from 'react';

import Navbar from './navbar'
import Footer from './footer'
import Connect from './connect'
import { ReactNode } from 'react';

import { Flex, useColorModeValue } from '@chakra-ui/react'

import { useAccount } from 'wagmi'

export default function Layout({ children }: { children: ReactNode }) {
  const [connected, setConnected] = useState(false)
  const { isConnected } = useAccount({
    onConnect: () => {
      setConnected(true)
    },
    onDisconnect: () => {
      setConnected(false)
    }
  })

  useEffect(() => {
    setConnected(isConnected)
  }, [isConnected])

  if (!connected) {
    return (
      <Flex flexDir='column' h='100vh'
      bg={useColorModeValue('blackAlpha.200', 'blackAlpha.600')}
    >
      <Navbar />
      {!connected && <Connect />}
      <Footer />
    </Flex>
    )
  }

  return (
    <Flex flexDir='column' h='100vh'
      bg={useColorModeValue('blackAlpha.200', 'blackAlpha.600')}
    >
      <Navbar />
      <Flex grow={1}>
        {children}
      </Flex>
      <Footer />
    </Flex>
  )
}
