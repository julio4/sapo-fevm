import { useState } from 'react'

import { JobProvider } from '@/components/Context/JobContext'
import Layout from '@/components/Layout/layout'
import { Text, Flex, Box, useColorModeValue, Heading, Divider, Badge, Slide } from '@chakra-ui/react'

type JobResult = {
  id: string
  status: string
  input: {
    cid: string
    size: number
  }
  output: {
    cid: string
    size: number
  }
  createdAt: string
}

const jobsList: JobResult[] = [
  {
    id: '1',
    status: 'success',
    input: {
      cid: 'QmX1',
      size: 100
    },
    output: {
      cid: 'QmX2',
      size: 200
    },
    createdAt: '2021-01-01'
  },
  {
    id: '2',
    status: 'fail',
    input: {
      cid: 'QmX3',
      size: 100
    },
    output: {
      cid: 'QmX4',
      size: 200
    },
    createdAt: '2021-01-01'
  },
]

const JobsList = (
  { jobs, selected, onSelect }:
    {
      jobs: JobResult[],
      onSelect: (job: JobResult) => void
      selected: JobResult | null
    }
) => {
  return (
    <Box
      w='full'
      m={4}
      mr={2}
      py={4}
      bg={useColorModeValue('whiteAlpha.700', 'transparent')}
      backdropFilter='auto'
      backdropBlur='2px'
      border="1px"
      borderColor={useColorModeValue('gray.200', 'gray.700')}
      borderRadius="lg"
    >
      <Box px={4}>
        <Heading size="md" mb={2}>Jobs List</Heading>
      </Box>
      <Divider />

      {jobs.map((job) => (
        <Box
          key={job.id}
          transition='all 0.2s ease-in'
          bg={selected && selected.id === job.id 
            ? useColorModeValue('teal.50', 'teal.800') 
            : useColorModeValue('whiteAlpha.700', 'transparent')
          }
          _hover={{
            bg: useColorModeValue('blackAlpha.100', 'blackAlpha.400'),
            cursor: 'pointer'
          }}
          onClick={() => onSelect(job)}
        >
          <Flex px={4} py={2} justifyContent='space-between'>
            <Text>{job.id}</Text>
            <Badge
              colorScheme={job.status === 'success' ? 'green' : 'red'}
            >{job.status}</Badge>
          </Flex>
          <Divider />
        </Box>
      ))}

    </Box>
  )
}

const JobDetails = ({ job }: { job: JobResult | null }) => {
  return (
    <Box
      w={job ? 'full' : '0'}
      m={4}
      ml={2}
      py={4}
      bg={useColorModeValue('whiteAlpha.700', 'transparent')}
      backdropFilter='auto'
      backdropBlur='2px'
      border="1px"
      borderColor={useColorModeValue('gray.200', 'gray.700')}
      borderRadius="lg"
    >
      <Box px={4}>
        <Heading size="md" mb={2}>Job Detail</Heading>
      </Box>
      <Divider />

      <Box px={4} py={2}>
        <Text>Job ID: {job?.id}</Text>
        <Text>Job Status: {job?.status}</Text>
        <Text>Job Created At: {job?.createdAt}</Text>
        <Text>Job Input CID: {job?.input.cid}</Text>
        <Text>Job Input Size: {job?.input.size}</Text>
        <Text>Job Output CID: {job?.output.cid}</Text>
        <Text>Job Output Size: {job?.output.size}</Text>
      </Box>
    </Box>
  )
}

export default function JobsHistory() {
  const [job, setJob] = useState<JobResult | null>(null)

  return (
    <JobProvider>
      <Layout>
        <Flex w='full'>
          <JobsList 
            jobs={jobsList} 
            onSelect={(selected) => setJob(selected)} 
            selected={job}/>
          {job && <JobDetails job={job} />}
        </Flex>
      </Layout>
    </JobProvider>
  )
}
