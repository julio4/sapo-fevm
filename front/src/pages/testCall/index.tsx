import Layout from "../../components/layout";
import TextModal from "../../components/textModal";
import ImageModal from "../../components/imageModal";
import { Card, CardBody } from "@chakra-ui/react";

import * as React from "react";
import { Flex, Heading, Button, Text, Input } from "@chakra-ui/react";
import { useContractRead } from "wagmi";

export default function Home() {
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
    <Layout>
      <Flex direction="column" p={12} rounded={6} alignItems="start">
        <Heading mb={6}>Testing calling the function</Heading>
        <Text mb={3}>Test Dao:</Text>
        <Input
          mb={3}
          onChange={(e) => setdaoAddress(e.target.value)}
          placeholder="Enter DAO Adress"
        />

        <Button onClick={discoverDAO}>Check the DAO</Button>
        <Flex direction="column" p={12} rounded={6} alignItems="start">
          <Text mb={4}>Text Cids:</Text>
          {textCids.map((cid, index) => (
            <React.Fragment key={index}>
              <TextModal cid={cid} />
              <br />
            </React.Fragment>
          ))}
        </Flex>
        <Flex direction="column" p={12} rounded={6} alignItems="start">
          <Text mb={4}>Image Cids:</Text>
          {imgCids.map((cid, index) => (
            <Flex mb={3} key={index} direction="row" alignItems="center">
              <ImageModal cid={cid} />
            </Flex>
          ))}
        </Flex>
      </Flex>
    </Layout>
  );
}
