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

import { useContractRead } from "wagmi";

import * as React from "react";

export default function DaoNavBar() {
  const [daoAddress, setdaoAddress] = React.useState("");
  const [cids, setCids] = React.useState<string[]>([]);
  const [rawCids, setRawCids] = React.useState<string[]>([]);
  const [cidType, setCidType] = React.useState<string[][]>([]);

  let { data, isError, isLoading } = useContractRead({
    address: daoAddress,
    abi: [
      {
        inputs: [],
        name: "getAllCids",
        outputs: [
          {
            internalType: "bytes[]",
            name: "",
            type: "bytes[]",
          },
        ],
        stateMutability: "view",
        type: "function",
      },
    ],
    functionName: "getAllCids",
  });

  function hex2a(hex) {
    var hex = hex.toString(); //force conversion
    var str = "";
    for (var i = 0; i < hex.length; i += 2)
      str += String.fromCharCode(parseInt(hex.substr(i, 2), 16));
    return str.substring(1);
  }

  const convertToCids = () => {
    let newCids: string[] = [];
    rawCids.forEach((raw) => {
      newCids.push(hex2a(raw));
    });
    setCids(newCids);
  };

  const getCidType = async () => {
    let newTypeCids: string[][] = [];

    cids.forEach(async (cid) => {
      let file = "https://ipfs.io/ipfs/" + cid;
      var req = await fetch(file, { method: "HEAD" });
      let type = req.headers.get("content-type");
      type = type || "undefined";
      newTypeCids.push([cid, type]);
    });
    setCidType(newTypeCids);
  };

  const sortAndFormatCids = async () => {
    console.log("Raw CIDS: ", rawCids);
    convertToCids();
    console.log("Clean CIDS: ", cids);
    await getCidType();
    console.log("Cid types:", cidType);
  };

  React.useEffect(() => {
    if (data) {
      setRawCids(data);
    }
  }, [data]);

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
            <Button onClick={sortAndFormatCids}>See CIDs</Button>
          </Box>

          <Box overflow="hidden">
            <Heading size="s" mb={3} textTransform="uppercase">
              Raw Cids
            </Heading>
            <Flex direction="column" p={0} rounded={6} alignItems="start">
              {cids.map((cid, index) => (
                <React.Fragment key={index}>
                  <Text fontSize="sm">
                    {index}. {cid}{" "}
                  </Text>
                </React.Fragment>
              ))}
            </Flex>
          </Box>

          <Box overflow="hidden">
            <Heading size="s" mb={3} textTransform="uppercase">
              Type Cids
            </Heading>
          </Box>
        </Stack>
      </CardBody>
    </Card>
  );
}
