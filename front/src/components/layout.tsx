import Navbar from './navbar'
import Footer from './footer'
import Connect from './connect'
import { ReactNode } from 'react';

import { Flex } from '@chakra-ui/react'

import { useAccount } from 'wagmi'

export default function Layout({ children }: { children: ReactNode }) {
  const { isConnected } = useAccount()
  return (
    <>
      <Navbar />
      {isConnected ?
        <Flex height="90vh">{children}</Flex>
        : <Connect />
      }
      <Footer />
    </>
  )
}
