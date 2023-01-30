import { useRouter } from 'next/router'
import { Flex, Heading, Button, useColorModeValue } from '@chakra-ui/react'

export default function HeroPage() {
  const router = useRouter()
  return (
    <Flex
      height="100vh"
      justifyContent="center"
      alignItems="center"
      bgGradient={useColorModeValue("radial(green.100 5%, whiteAlpha.200 95%)","radial(green.800 5%, whiteAlpha.200 95%)")} >
      <Flex direction="column" p={12} rounded={6} alignItems="center">
        <Heading as="h1" size="4xl" mb={6}>Sapo 🐸</Heading>
        <Heading as="h2" size="md" mb={6}>Your gateway to decentralized computation</Heading>
        <Button onClick={() => router.push('/app')}>
          Enter app
        </Button>
      </Flex>
    </Flex>
  )
}
