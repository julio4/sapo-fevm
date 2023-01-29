import Navbar from './navbar'
import Footer from './footer'
import { ReactNode } from 'react';

import { Flex } from '@chakra-ui/react'

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <>
      <Navbar />
      <Flex height="90vh">{children}</Flex>
      <Footer />
    </>
  )
}
