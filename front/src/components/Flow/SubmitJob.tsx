import { useState, useEffect } from "react";
import {
  Spinner,
  Box,
  Flex,
  Text,
  useColorModeValue,
  Button,
  Input,
  SimpleGrid,
  FormControl,
  FormLabel,
  FormHelperText,
} from "@chakra-ui/react";

import { JobRequest, useJobContext, } from "../Context/JobContext";

import AbiSapoBridge from "@/constants/AbiSapoBridge.json";
import AddressSapoBridge from "@/constants/AddressSapoBridge.json";
import { useContractWrite, usePrepareContractWrite, useWaitForTransaction } from "wagmi";
import lighthouse from '@lighthouse-web3/sdk';
import { ethers } from "ethers";

const SpecCard = ({ title, value, desc }: { title: string, value: string, desc: string }) => (
  <FormControl>
    <FormLabel>{title}</FormLabel>
    <Input
      placeholder="ERROR"
      value={value}
      bg={useColorModeValue('gray.100', 'gray.800')}
      border={0}
      color={useColorModeValue('gray.500', 'gray.200')}
      _placeholder={{
        color: 'gray.500',
      }}
      _disabled={{
        bg: useColorModeValue('gray.100', 'gray.800'),
        color: useColorModeValue('gray.500', 'gray.200'),
        cursor: 'not-allowed'
      }}
      disabled
    />
    <FormHelperText>{desc}</FormHelperText>
  </FormControl>
)

function byteArrayToHex(byteArray) {
  let hexString = '0x';
  for (const byte of byteArray) {
    hexString += byte.toString(16).padStart(2, '0');
  }
  return hexString;
}

const cidToHex = (cid: string) => {
  const encoder = new TextEncoder();
  const bArr = encoder.encode(cid);
  const bArr1 = bArr.slice(0, 32);
  let bArr2 = bArr.slice(32);

  if (bArr2.length < 32) {
    bArr2 = [...bArr2, ...Array(32 - bArr2.length).fill(0)];
  }

  return [byteArrayToHex(bArr1) , byteArrayToHex(bArr2)];
}

export default function SubmitJob() {
  const { step, setStep, jobRequest, setJobRequest } = useJobContext();

  const [ request, setRequest ] = useState<JobRequest>(jobRequest);

  const { config, error, isError } = usePrepareContractWrite({
    address: AddressSapoBridge.address,
    abi: AbiSapoBridge,
    functionName: "request",
    args: cidToHex(request.input?.cid || ""),
    overrides: {
      gasLimit: ethers.BigNumber.from(1000000000),
      value: ethers.utils.parseEther('0.1'),
    }
  });

  const { data, write } = useContractWrite(config);

  const { isLoading, isSuccess } = useWaitForTransaction({
    hash: data?.hash,
    onSuccess: () => {
      setJobRequest(request);
      setStep(3);
    },
  })

  const handleSubmit = async () => {
    // Wrap Job spec in a IPFS object
    const jobSpec = {
      "image": jobRequest.job?.image,
      "inputsCid": jobRequest.input?.cid,
      "runParams": jobRequest.job?.runParams
    }
    const jobSpecInput = await lighthouse.uploadText(
      JSON.stringify(jobSpec),
      process.env.lightHouseApi
    );

    console.log("wrapped input to pinned ", jobSpecInput)

    // Update request state, trigger write
    setRequest((prev) => ({
      ...prev,
      input: {
        cid: jobSpecInput.data.Hash,
        type: "application/json",
        size: parseInt(jobSpecInput.data.Size.toString())
      }
    }))
  }

  useEffect(() => {
    console.log("trying to confirm tx", request)
    if (request.input?.cid) {
      write?.();
    }
  }, [request])

  const descriptionText = useColorModeValue('gray.600', 'gray.400')

  return (
    <Flex direction='column' h='full' w='full' p={8}>
      <Box alignContent='center' justifyContent='center' mb={4}>
        <Text align='center' fontSize="2xl" fontFamily="system-ui" fontWeight="bold">
          Submit Job
        </Text>
      </Box>

      <Flex
        as={'form'}
        flexDir='column'
        w="full"
        bg={useColorModeValue('whiteAlpha.700', 'transparent')}
        backdropFilter='auto'
        backdropBlur='2px'
        border="1px"
        borderColor={useColorModeValue('gray.200', 'gray.700')}
        borderRadius="lg"
      >
        <SimpleGrid columns={{xl: 2, lg: 1}} spacing={8} p={4}>

          <SpecCard 
            title="Job name" 
            value={jobRequest.job?.name || (jobRequest.custom && "Custom") || ""} 
            desc={(jobRequest.custom && "Your custom job") || jobRequest.job?.desc} />

          <SpecCard 
            title="Docker image" 
            value={jobRequest.job?.image || ""} 
            desc="The docker image used for the job execution" />

          <SpecCard 
            title="Input file" 
            value={jobRequest.usrInput?.cid || ""} 
            desc="Source file from IPFS/Filecoin CID" />

          <SpecCard 
            title="Approximate price" 
            value='~0.1 FIL'
            desc="If the job fail, you'll be refunded atleast 50%" />
        </SimpleGrid>

        <Box p={4}>
          <Button
            fontFamily={'heading'}
            mt={8}
            w={'full'}
            bgGradient={useColorModeValue(
              'linear(to-r, green.300, teal.300)',
              'linear(to-r, green.500, teal.500)')}
            transition={'all 0.3s ease-in-out'}
            color={'white'}
            _hover={{
              opacity: '0.9',
              boxShadow: 'md',
              transform: 'scale(1.005)'
            }}
            disabled={!write}
            onClick={handleSubmit}>
            Submit
          </Button>
        </Box>

        {isLoading &&
          <Flex my={8} direction='column' alignContent='center' justifyContent='center'>
            <Spinner
              alignSelf='center'
              thickness='4px'
              speed='0.65s'
              emptyColor='gray.200'
              color='green.500'
              size='xl'
            />
            <Text color={descriptionText}
              mt={4} align='center' fontSize="xl" fontFamily="system-ui" fontWeight="bold">
              Waiting for transaction confirmation...
            </Text>
          </Flex>
        }

      </Flex>
    </Flex>
  );
}
