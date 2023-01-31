import { Text, Button, Card, CardBody, Flex } from "@chakra-ui/react";
import * as React from "react";

export default function TextModal({ cid }) {
  const [showModal, setShowModal] = React.useState(false);
  const hideModal = () => setShowModal(false);

  const [imageURI, setImageURI] = React.useState("");
  const [cidName, setCidName] = React.useState("");
  const [cidDescription, setCidDescription] = React.useState("");
  const [cidText, setCidText] = React.useState("");

  const handleShow = () => {
    updateUI();
    setShowModal(!showModal);
  };

  async function updateUI() {
    if (cid) {
      const requestURL =
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/" + cid;
      console.log("request url: ", requestURL);
      const cidURIResponse = await (await fetch(requestURL)).json();
      const text = cidURIResponse.text;

      setCidName(cidURIResponse.name);
      setCidDescription(cidURIResponse.description);
      setCidText(text);
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
              <Text>Content: {cidText}</Text>
            </Flex>
          ) : (
            <div></div>
          )}
        </CardBody>
      </Card>
    </>
  );
}
