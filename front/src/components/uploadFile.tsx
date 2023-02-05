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
  Progress,
  Input,
  Checkbox,
  CheckboxGroup,
} from "@chakra-ui/react";

import * as React from "react";
import lighthouse from "@lighthouse-web3/sdk";
const lighthouseApi = "YOURAPIKEY";

export default function UploadCard() {
  const [cid, setCid] = React.useState("");
  const [progressBar, setProgressbar] = React.useState(0);
  const [loaded, setLoaded] = React.useState(false);
  const [cidLoaded, setCidLoaded] = React.useState("");
  const [privacyMod, setPrivacyMod] = React.useState(false);

  const progressCallback = (progressData) => {
    let percentageDone =
      100 - (progressData?.total / progressData?.uploaded)?.toFixed(2);
    setProgressbar(percentageDone);
  };

  const deploy = async (e) => {
    // Push file to lighthouse node
    const output = await lighthouse.upload(e, lighthouseApi, progressCallback);
    setCidLoaded(output.data.Hash);
    console.log("File Status:", output);
    /*
      output:
        {
          Name: "filename.txt",
          Size: 88000,
          Hash: "QmWNmn2gr4ZihNPqaC5oTeePsHvFtkWNpjY3cD6Fd5am1w"
        }
      Note: Hash in response is CID.
    */

    console.log(
      "Visit at https://gateway.lighthouse.storage/ipfs/" + output.data.Hash
    );
  };

  function handleDeploy(e) {
    if (privacyMod) {
      console.log("ITS PRIVACY MOD");
    } else {
      deploy(e);
    }
  }

  return (
    <Card mt={6}>
      <CardBody>
        <Heading size="md" mb={4}>
          Upload File
        </Heading>
        <Checkbox mb={3} onChange={(e) => setPrivacyMod(!privacyMod)}>
          Privacy Mode
        </Checkbox>

        <Input onChange={(e) => handleDeploy(e)} type="file" />
        {progressBar > 0 && loaded ? (
          <Progress hasStripe value={progressBar} />
        ) : (
          <div></div>
        )}
        {cidLoaded ? <Text>Cid: {cidLoaded}</Text> : <div></div>}
      </CardBody>
      <Button onClick={() => console.log(privacyMod)}>Click</Button>
    </Card>
  );
}
