import { Text, Button, Card, CardBody, Flex } from "@chakra-ui/react";
import * as React from "react";
import Image from "next/image";

export default function ImageModal({ cid }) {
  const [showModal, setShowModal] = React.useState(false);
  const hideModal = () => setShowModal(false);

  const [cidName, setCidName] = React.useState("");
  const [cidDescription, setCidDescription] = React.useState("");
  const [cidImage, setCidImage] = React.useState("");

  const handleShow = () => {
    updateUI();
    console.log("image uri is: ", cidImage);
    setShowModal(!showModal);
  };

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

  return (
    <>
      <Card>
        <CardBody>
          <Text>CID: {cid}</Text>
          <Button ml={4} onClick={handleShow}>
            {showModal ? <Text>Hide</Text> : <Text>Show</Text>}
          </Button>

          {showModal ? (
            <Flex direction="column" p={12} rounded={6} alignItems="start">
              <Text>Title: {cidName}</Text>
              <Text>Description: {cidDescription}</Text>
              <Image
                loader={() => cidImage}
                src={cidImage}
                width={200}
                height={200}
                alt="Image of the CID"
              />
            </Flex>
          ) : (
            <div></div>
          )}
        </CardBody>
      </Card>
    </>
  );
}
