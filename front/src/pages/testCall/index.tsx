import Layout from "../../components/layout";
import textModal from "../../components/textModal";

import * as React from "react";
import { Flex, Heading, Button, Text, Input } from "@chakra-ui/react";
import { useContractRead } from "wagmi";

export default function Home() {
  const [daoAddress, setdaoAddress] = React.useState("");
  const [cids, setCids] = React.useState<string[]>([]);
  const [imgCids, setImgCids] = React.useState<string[]>([]);
  const [textCids, setTextCids] = React.useState<string[]>([]);
  const [showModal, setShowModal] = React.useState(false);

  const hideModal = () => setShowModal(false);

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

  data = [
    "QmTkuHzN6ZvJGbuk3QAfpva8nD1shqApq9zyoVCpnFHqfi",
    "QmRfB6rmGfmGM2aTEgBWd6GmgyX74GMc7ihYkCAayq1Bs9",
    "Qme9qHPBxrMTN3r6TeEQuucLZS3WvKYQosNjTnWtzPsZQg",
    "QmPNfuJH8AoHgDxheDvxrWmD99NgHcMHJ8iBzpeDByk71Z",
    "QmRfB6rmGfmGM2aTEgBWd6GmgyX74GMc7ihYkCAayq1Bs9",
  ];

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

  const consoleStuff = async () => {
    await sortAndFormatCids();
    console.log("cids: ", cids);
    console.log("cids img: ", imgCids);
    console.log("cids text: ", textCids);
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

        <Button onClick={consoleStuff}>Test</Button>
        <Flex direction="column" p={12} rounded={6} alignItems="start">
          <Text mb={4}>Text Cids:</Text>
          {textCids.map((cid, index) => (
            <React.Fragment key={index}>
              {cid}
              <br />
            </React.Fragment>
          ))}
        </Flex>
        <Flex direction="column" p={12} rounded={6} alignItems="start">
          <Text mb={4}>Image Cids:</Text>
          {imgCids.map((cid, index) => (
            <Flex mb={3} key={index} direction="row" alignItems="center">
              <Text>{cid}</Text>
              <Button ml={4}>Show</Button>
            </Flex>
          ))}
        </Flex>
      </Flex>
    </Layout>
  );
}
