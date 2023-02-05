import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import Image from 'next/image'
import { Text, Flex, Heading, Button, useColorModeValue, Box } from '@chakra-ui/react'

import { FaChevronDown, FaChevronUp } from 'react-icons/fa'
import Cloud from '../../public/cloud.svg'
import Connected from '../../public/connected.svg'
import Growth from '../../public/growth.svg'
import CloudDark from '../../public/cloudDark.svg'
import ConnectedDark from '../../public/connectedDark.svg'
import GrowthDark from '../../public/growthDark.svg'

import ModeToggler from '@/components/Layout/modeToggler'

const HeroItem = ({ page, img, imgDark, heading, desc, left }) => (
  <Flex
    height="100vh"
    justifyContent="center"
    alignItems="center"
    transform={"translateY(-" + (page*100).toString() + "%)"}
    transition="transform 0.5s"
  >
    <Flex direction="row" p={12} rounded={6} alignItems="center" justifyContent="space-evenly" w="full" gap={4}>
      {left ? <Image src={useColorModeValue(img, imgDark)} alt="Connected" width={600} /> : null}
      <Box>
        <Heading as="h1" size="xl" mb={6}>{heading}</Heading>
        <Text color={useColorModeValue("gray.700", "gray.200")}>
          {desc}
        </Text>
      </Box>
      {!left ? <Image src={useColorModeValue(img, imgDark)} alt="Connected" width={600} /> : null}
    </Flex>
  </Flex>
)

export default function HeroPage() {
  const [ page, setPage ] = useState(0)
  const [ thresold, setThresold ] = useState(0)
  const router = useRouter()

  const items = [
    {
      img: Cloud,
      imgDark: CloudDark,
      left: false,
      heading: 'Secure, scalable, and reliable computing resources for everyone',
      desc: 'Sapo and Bacalhau bring you a versatile and affordable answer to your cloud computing needs. Any DataDao can now leverage the power of computation on their own data stored on Filecoin while enjoying enhanced security, privacy, and verification through decentralization with Sapo.'
    },
    {
      img: Connected,
      imgDark: ConnectedDark,
      left: true,
      heading: 'A new Decentralized Economy',
      desc: 'Our platform empowers individuals and organizations to directly connect, harnessing the value of their data and driving the growth of a new decentralized global economy.'
    },
    {
      img: Growth,
      imgDark: GrowthDark,
      left: false,
      heading: 'Your one-stop solution for all your needs',
      desc: 'Unlock the full potential of your data with Sapo. Whether you need a generic job or a custom one, our platform provides a complete and convenient solution for all your cloud computing needs. Experience unparalleled security, privacy, and control with Sapo\'s decentralized approach to cloud computing.'
    }
  ]

  const onScroll = (event: any) => {
    const isDown = parseInt(event.deltaY) > 0;
    setThresold(thresold + 1)
    console.log(thresold)
    if (thresold > 3) {
      setThresold(0)
      setPage(isDown ? (page >= items.length ? page : page + 1) : 0)
    }
  }

  useEffect(() => {
    window.addEventListener('wheel', onScroll);
    return () => {
      window.removeEventListener('scroll', onScroll);
    }
  }, [thresold]);

  return (
    <Box height={"200vh"}
        bgGradient={useColorModeValue("radial(green.100 5%, whiteAlpha.200 95%)","radial(green.800 5%, whiteAlpha.200 95%)")}
        >
      <Flex
        height="100vh"
        justifyContent="center"
        alignItems="center"
        transform={"translateY(-" + (page*100).toString() + "%)"}
        transition="transform 0.5s"
        >
        <Flex direction="column" p={12} rounded={6} alignItems="center">
          <Heading as="h1" size="4xl" mb={8}>Sapo 🐸</Heading>
          <Heading as="h2" textAlign="center" maxW="75%" size="md" mb={6}
            color={useColorModeValue("gray.700","gray.200")}
            >
            Discover the power of decentralized cloud computing with Sapo and experience greater security, privacy, and control over your data and applications.
            </Heading>
          <Button size='lg' onClick={() => router.push('/app')}
            transition="all 0.3s"
            bgGradient={useColorModeValue("linear(to-r, teal.200, green.300)", "linear(to-r, teal.600, green.700)")}
            _hover={{
              bgGradient: useColorModeValue("linear(to-r, teal.200, green.300)", "linear(to-r, teal.500, green.600)"),
              boxShadow: "xl",
              transform: "scale(1.05)"
            }}>
            Enter app
          </Button>
        </Flex>
      </Flex>

      {items.map((item, index) => (
        <HeroItem key={index} page={page} {...item} />
      ))}

      {(page > 0 && page <= items.length) && (
      <Box
        pos={"fixed"}
        top="5%"
        right="49vw"
        transition="all 0.5s"
        _hover={{ 
          cursor: "pointer",
          transform: "scale(1.2)",
          color: useColorModeValue("green.500","green.200")
        }}
        onClick={() => setPage(page - 1)}
      >
        <FaChevronUp size="2vw" />
      </Box>
      )}

      {(page >= 0 && page < items.length) && (
      <Box
        pos={"fixed"}
        bottom="5%"
        right="49vw"
        transition="all 0.5s"
        _hover={{ 
          cursor: "pointer",
          transform: "scale(1.2)",
          color: useColorModeValue("green.500","green.200")
        }}
        onClick={() => setPage(page + 1)}
      >
        <FaChevronDown size="2vw" />
      </Box>
      )}
      <Box
        pos={"fixed"}
        top="0"
        right="0"
        p={4}>
        <ModeToggler />
      </Box>
    </Box>
  )
}
