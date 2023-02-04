import { useJobContext, Job } from '@/components/Context/JobContext';

import {
    Flex,
    Box,
    SimpleGrid,
    Text,
    Badge,
    useColorModeValue
} from '@chakra-ui/react'

const JobCard = ({ job, onClick }: { job: Job, onClick: () => void }) => (
    <Box
        alignItems="center"
        justifyContent="center"
        transition={'all 0.2s ease-in-out'}
        _hover={{
            cursor: 'pointer',
            transform: 'scale(1.01)',
        }}
        onClick={onClick}
        >
        <Box
            bg={useColorModeValue('white', 'gray.800')}
            maxW="sm"
            borderWidth="1px"
            rounded="lg"
            shadow="lg"
            position="relative">

            {/* {data.isNew && (
          <Circle
            size="10px"
            position="absolute"
            top={2}
            right={2}
            bg="red.200"
          />
        )} */}

            <Box
                bg={useColorModeValue(
                    (job.category?.color || "gray") + '.100', 
                    (job.category?.color || "gray") + '.900')}
                height="100px"
                roundedTop="lg"
            />

            <Box p="6">
                <Flex alignItems="baseline">
                    {job.category && (
                        <Badge rounded="full" px="2" fontSize="0.8em" 
                            colorScheme={job.category.color || "gray"}>
                            {job.category.name}
                        </Badge>
                    )}
                </Flex>
                <Flex mt="1" justifyContent="space-between" alignContent="center">
                    <Box
                        fontSize="2xl"
                        fontWeight="semibold"
                        as="h4"
                        lineHeight="tight"
                        isTruncated>
                        {job.name}
                    </Box>
                </Flex>

            </Box>
        </Box>
    </Box>
)

export default function SelectJob() {
    const { category, allJobs, setJob, setStep } = useJobContext();
    return (
        <Flex direction='column' h='full' w='full' p={8}>
            <Box alignContent='center' justifyContent='center' mb={4}>
                <Text align='center' fontSize="2xl" fontFamily="monospace" fontWeight="bold">
                    Select a job
                </Text>
            </Box>

            <SimpleGrid columns={{xl: 4, lg: 3, md: 2}} spacing={8}>
                {allJobs.filter((job) => category === null || job.category === category).map((job) => {
                    return (
                        <JobCard key={job.name} job={job} onClick={() => {
                            setJob(job);
                            setStep(1);
                        }}/>
                    )
                })}
            </SimpleGrid>
        </Flex>
    );
}
