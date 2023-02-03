import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Heading,
  Text,
  Stack,
  Box,
  StackDivider,
  Input,
  Button,
  Flex,
} from "@chakra-ui/react";

import TextModal from "./textModal";
import ImageModal from "./imageModal";
import { useContractRead } from "wagmi";

import * as React from "react";

export default function DaoNavBar() {
  const [daoAddress, setdaoAddress] = React.useState("");
  const [cids, setCids] = React.useState<string[]>([]);
  const [imgCids, setImgCids] = React.useState<string[]>([]);
  const [textCids, setTextCids] = React.useState<string[]>([]);

  let { data, isError, isLoading } = useContractRead({
    address: daoAddress,
    abi: [
      {
        name: "getCidStored",
        type: "function",
        stateMutability: "view",
        inputs: [],
        outputs: [
          {
            internalType: "string[]",
            name: "",
            type: "string[]",
          },
        ],
      },
    ],
    functionName: "getCidStored",
  });

  const sortAndFormatCids = async () => {
    let newImgCids: string[] = [];
    let newTextCids: string[] = [];
    for (const cid of cids) {
      const requestURL =
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/" + cid;
      const cidUriResponse = await (await fetch(requestURL)).json();
      if (cidUriResponse.image) {
        console.log("its an image");
        newImgCids = newImgCids.concat(cid);
      } else if (cidUriResponse.text) {
        console.log("its a text");

        newTextCids = newTextCids.concat(cid);
      }
    }
    setImgCids(newImgCids);
    setTextCids(newTextCids);
  };

  React.useEffect(() => {
    if (data) {
      setCids([...data]);
    }
    sortAndFormatCids;
  }, [data]);

  const discoverDAO = async () => {
    await sortAndFormatCids();
  };
  return (
    <Card overflow="hidden">
      <CardHeader>
        <Heading size="md">See your DataDAO</Heading>
      </CardHeader>

      <CardBody>
        <Stack divider={<StackDivider />} spacing="4">
          <Box>
            <Heading size="xs" mb={2} textTransform="uppercase">
              Enter Address
            </Heading>
            <Input
              mb={3}
              onChange={(e) => setdaoAddress(e.target.value)}
              placeholder="Enter DAO Adress"
            />
            <Button onClick={discoverDAO}>See CIDs</Button>
          </Box>
          <Box overflow="hidden">
            <Heading size="s" mb={3} textTransform="uppercase">
              Text Cids
            </Heading>
            <Flex direction="column" p={0} rounded={6} alignItems="start">
              {textCids.map((cid, index) => (
                <React.Fragment key={index}>
                  <Text fontSize="sm">{cid} </Text>
                </React.Fragment>
              ))}
            </Flex>
          </Box>
          <Box overflow="hidden">
            <Heading size="s" textTransform="uppercase" mb={3}>
              Image CIDs
            </Heading>
            {/* <Text pt="2" fontSize="sm">
              See a detailed analysis of all your business clients.
            </Text> */}
            <Flex direction="column" p={0} rounded={6} alignItems="start">
              {imgCids.map((cid, index) => (
                <React.Fragment key={index}>
                  <Text fontSize="sm">{cid} </Text>
                </React.Fragment>
              ))}
            </Flex>
          </Box>
        </Stack>
      </CardBody>
    </Card>
  );
}
