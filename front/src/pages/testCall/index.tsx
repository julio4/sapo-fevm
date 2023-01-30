import Layout from "../../components/layout";
import * as React from "react";
import { Flex, Heading, Button, Text, Input } from "@chakra-ui/react";
import { useContractRead } from "wagmi";

export default function Home() {
  const [daoAddress, setdaoAddress] = React.useState("");
  const [cids, setCids] = React.useState<string[]>([]);

  const { data, isError, isLoading } = useContractRead({
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

  React.useEffect(() => {
    if (data) {
      setCids([...data]);
    }
  }, [data]);

  return (
    <Layout>
      <Flex direction="column" p={12} rounded={6} alignItems="center">
        <Heading mb={6}>Testing calling the function</Heading>
        <Text mb={3}>Test Dao:</Text>
        <Input
          mb={3}
          onChange={(e) => setdaoAddress(e.target.value)}
          placeholder="Enter DAO Adress"
        />

        <Button onClick={() => console.log(cids)}>Test</Button>
        <Text>Cids: {data}</Text>
      </Flex>
    </Layout>
  );
}
