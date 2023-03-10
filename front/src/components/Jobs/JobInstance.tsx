import { useEffect, useState } from "react";

import { JobProvider } from "@/components/Context/JobContext";
import Layout from "@/components/Layout/layout";
import {
  Text,
  Flex,
  Box,
  useColorModeValue,
  Divider,
  Button,
  Badge,
} from "@chakra-ui/react";

import { CheckCircleIcon, TimeIcon, WarningIcon } from "@chakra-ui/icons";
import { useContractRead, useAccount } from "wagmi";
import { useDisclosure, defineStyleConfig } from "@chakra-ui/react";
import JobResult from "@/components/Jobs/JobResult";
import JobSummary from "./JobSummary";
import AbiSapoJob from "@/constants/AbiSapoJob.json";

const statusText: string[] = ["Pending", "Completed", "Failed", "Loading..."];

const statusColor: string[] = ["blue", "green", "red", "orange"];

const JobInstance = ({
  onSelect,
  jobAddress,
  selected,
  id,
}: {
  onSelect: (select: JobSummary) => void;
  jobAddress: `0x${string}`;
  selected: `0x${string}` | null;
  id: number;
}) => {
  const [outputUrl, setOutputUrl] = useState<string>("");
  const [exitCode, setExitCode] = useState<string>("");
  const [cidOutputs, setCidOutputs] = useState<string>("");
  const [cidStderr, setCidStderr] = useState<string>("");
  const [cidStdout, setCidStdout] = useState<string>("");
  const [stderrContent, setStderrContent] = useState<string>("");
  const [stdoutContent, setStdoutContent] = useState<string>("");
  const [outputsNamesAndCids, setOutputsNamesAndCids] = useState<string[]>([]);
  const getResultAbi = AbiSapoJob.find((v) => v.name === "getResult");
  const getResultCidAbi = AbiSapoJob.find((v) => v.name === "getResultCid");
  const getStatusAbi = AbiSapoJob.find((v) => v.name === "getStatus");

  // TODO Bacalhau
  let [job, setJob] = useState<JobResult>({
    id: id,
    address: jobAddress,
    exitCode: NaN,
    outputs: [""],
    stderr: "",
    stdout: "",
  });

  let {
    data: jobId,
    isLoading: jobIdIsLoading,
    isError: jobIdIsError,
  } = useContractRead({
    address: jobAddress,
    abi: [getResultAbi],
    functionName: "getResult",
  });

  let { data: cidResult } = useContractRead({
    address: jobAddress,
    abi: [getResultCidAbi],
    functionName: "getResultCid",
  });

  let res = useContractRead({
    address: jobAddress,
    abi: [getStatusAbi],
    functionName: "getStatus",
  });

  let { data: status, isLoading: statusIsLoading } = res;

  const unselectedColor = useColorModeValue("whiteAlpha.700", "transparent");
  const selectedColor = useColorModeValue("teal.50", "teal.800");

  async function getFileFromGateway(cid: string) {
    try {
      const request = await fetch(`https://ipfs.io/api/v0/ls/${cid}`);

      const responseJson = await request.json();
      const filesCid = responseJson.Objects[0].Links.map(
        (link: { Hash: any }) => link.Hash
      );
      // Mettre les filescid dans les states correspondants
      setCidOutputs(filesCid[1]);
      setCidStderr(filesCid[2]);
      setCidStdout(filesCid[3]);

      // List files in the outputs directory
      const request2 = await fetch(`https://ipfs.io/api/v0/ls/${filesCid[1]}`);
      const responseJson2 = await request2.json();
      const namesOutputsAndCids = responseJson2.Objects[0].Links.map(
        (link: { Name: any; Hash: any }) => {
          return { name: link.Name, cid: link.Hash };
        }
      );
      setOutputsNamesAndCids(namesOutputsAndCids);

      // Handling images
      const imageExtensions = [".jpg", ".jpeg", ".png", ".gif", ".bmp"];
      const imagelist = namesOutputsAndCids.filter(
        (filename: { name: string; cid: any }) => {
          const extension = filename.name.substr(
            filename.name.lastIndexOf(".")
          );
          if (imageExtensions.includes(extension)) {
            return [imageExtensions.includes(extension), filename.cid];
          }
        }
      );

      if (imagelist.length === 1) {
        const response = await fetch(
          `https://ipfs.io/ipfs/${cid}/outputs/${imagelist[0].name}`
        );
        const formattedUrl = response.url.replace(
          /http.*\/ipfs\//,
          "https://ipfs.io/ipfs/"
        );
        setOutputUrl(formattedUrl);
      } else {
        console.log("");
      }

      // Handling stdout content and stderr content and exit code
      const requestStderr = await fetch(`https://ipfs.io/ipfs/${cid}/stderr`);
      const textStderr = await requestStderr.text();
      setStderrContent(textStderr);

      const requestStdout = await fetch(`https://ipfs.io/ipfs/${cid}/stdout`);
      const textStdout = await requestStdout.text();
      setStdoutContent(textStdout);

      const requestExitCode = await fetch(
        `https://ipfs.io/ipfs/${cid}/exitCode`
      );
      const textExitCode = await requestExitCode.text();
      setExitCode(textExitCode);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    if (cidResult) {
      getFileFromGateway(cidResult);
    }
  }, [cidResult]);

  return (
    <Box
      key={job.id}
      transition="all 0.2s ease-in"
      bg={selected && selected === jobAddress ? selectedColor : unselectedColor}
      _hover={{
        bg: useColorModeValue("blackAlpha.100", "blackAlpha.400"),
        cursor: "pointer",
      }}
      onClick={() => {
        onSelect({
          ...job,
          jobId: jobId ? jobId : null,
          status: status ? status : null,
          outputUrl: outputUrl ? outputUrl : null,
          exitCode: exitCode ? exitCode : null,
          cidOutputs: cidOutputs ? cidOutputs : null,
          cidStderr: cidStderr ? cidStderr : null,
          cidStdout: cidStdout ? cidStdout : null,
          stderr: stderrContent ? stderrContent : null,
          stdout: stdoutContent ? stdoutContent : null,
          outputsNamesAndCids: outputsNamesAndCids ? outputsNamesAndCids : null,
          cidResult: cidResult ? cidResult : null,
        });
      }}
    >
      <Flex px={4} py={2} justifyContent="space-between" paddingX="5%">
        <Text>{job.id}</Text>

        <Text width={!selected ? "35%" : "50%"}>
          {selected && jobId && jobId.length > 22
            ? jobId.substring(0, 22) + "..."
            : jobId}
        </Text>
        <Badge
          colorScheme={statusColor[statusIsLoading ? 3 : status ? status : 0]}
          style={{ display: "flex", alignItems: "center" }}
          borderRadius="8px"
        >
          <Text width="6vw" align="center">
            {statusText[statusIsLoading ? 3 : status ? status : 0]}
          </Text>
        </Badge>
      </Flex>
      <Divider />
    </Box>
  );
};

export default JobInstance;
