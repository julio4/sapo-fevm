import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Stack,
  Button,
  Heading,
  Text,
  Flex,
  Input,
} from "@chakra-ui/react";

import Image from "next/image";
import ImageModal from "./imageModal";
import * as React from "react";

export default function ZoomCard() {
  const [cid, setCid] = React.useState("");
  const [cidName, setCidName] = React.useState("");
  const [cidDescription, setCidDescription] = React.useState("");
  const [cidImage, setCidImage] = React.useState("");
  const [cidText, setCidText] = React.useState("");
  const [showCid, setShowCid] = React.useState(false);

  async function updateUI() {
    if (cid) {
      const requestURL =
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/" + cid;
      console.log("request url: ", requestURL);
      const cidURIResponse = await (await fetch(requestURL)).json();
      const imageURI = cidURIResponse.image;
      const imageURIURL = imageURI.replace(
        "ipfs://",
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/"
      );

      setCidName(cidURIResponse.name);
      setCidDescription(cidURIResponse.description);
      setCidImage(imageURIURL);
    }
  }

  const handleClick = async () => {
    const requestURL =
      "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/" + cid;
    const cidUriResponse = await (await fetch(requestURL)).json();
    if (cidUriResponse.image) {
      const imageURI = cidUriResponse.image;
      const imageURIURL = imageURI.replace(
        "ipfs://",
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/"
      );
      setCidImage(imageURIURL);
      setCidText("");
    } else if (cidUriResponse.text) {
      setCidText(cidUriResponse.text);
      setCidImage("");
    }
    setCidName(cidUriResponse.name);
    setCidDescription(cidUriResponse.description);
    setShowCid(true);
  };

  return (
    <Card
      direction={{ base: "column", sm: "row" }}
      overflowX="hidden"
      variant="outline"
      mb={6}
      width={460}
    >
      <Stack>
        <CardBody>
          <Heading size="md">Inspect a CID</Heading>

          <Text py="2">
            Check the content of a CID, weither it is an image or a text
          </Text>
          <Input
            onChange={(e) => setCid(e.target.value)}
            placeholder="Enter CID Adress"
          />
        </CardBody>

        <CardFooter pt={0} display={"flex"} justifyContent={"center"}>
          <Button
            variant="solid"
            size="sm"
            colorScheme="blue"
            onClick={handleClick}
          >
            Inspect
          </Button>
        </CardFooter>
        {showCid ? (
          cidImage ? (
            <Flex alignItems={"center"} flexDirection={"column"} p={5}>
              <Text>Title: {cidName}</Text>
              <Text>Description: {cidDescription}</Text>
              <Image
                loader={() => cidImage}
                src={cidImage}
                width={200}
                height={200}
                alt="Image of the CID"
                style={{ borderRadius: "10%" }}
              />
            </Flex>
          ) : (
            <div>
              <Flex alignItems={"start"} flexDirection={"column"} p={5}>
                <Text mb={1}>Title: {cidName}</Text>
                <Text mb={2}>Description: {cidDescription}</Text>
                <Text>
                  Content: <br /> {cidText}
                </Text>
              </Flex>
            </div>
          )
        ) : (
          <div></div>
        )}
      </Stack>
    </Card>
  );
}
