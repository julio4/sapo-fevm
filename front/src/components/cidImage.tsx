import { useEffect, useState } from "react";
import Image from "next/image";
import React from "react";
import { Flex, Heading, Button, Text, Input } from "@chakra-ui/react";

export default function cidImage({ cid }) {
  cid = "QmdsqWZGCrkpjhCgCepEtFePmkvxkXda7P7w3z7Y13dxvH";
  const [imageURI, setImageURI] = useState("");
  const [cidName, setCidName] = useState("");
  const [cidDescription, setCidDescription] = useState("");

  async function updateUI() {
    console.log(`The cidURI is ${cid}`);
    // We are going to cheat a little here...
    if (cid) {
      // IPFS Gateway: A server that will return IPFS files from a "normal" URL.
      const requestURL =
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/" + cid;
      console.log("request url: ", requestURL);
      const cidURIResponse = await (await fetch(requestURL)).json();
      const imageURI = cidURIResponse.image;
      const imageURIURL = imageURI.replace(
        "ipfs://",
        "https://olive-absolute-silverfish-298.mypinata.cloud/ipfs/"
      );
      setImageURI(imageURIURL);
      setCidName(cidURIResponse.name);
      setCidDescription(cidURIResponse.description);
    }
  }

  return (
    <div>
      <Button onClick={updateUI}>Click to load</Button>
      {imageURI ? (
        <div className="container-img">
          <Image
            loader={() => imageURI}
            src={imageURI}
            width={200}
            height={200}
            alt="Image of the CID"
          />
        </div>
      ) : (
        <div></div>
      )}
    </div>
  );
}
