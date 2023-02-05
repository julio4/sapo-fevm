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

import AbiSapoBridge from "@/constants/AbiSapoBridge.json";
import * as React from "react";
import { useContractWrite, usePrepareContractWrite } from "wagmi";

export default function OcrCard() {
  const [cid, setCid] = React.useState("");
  const [cidName, setCidName] = React.useState("");
  const [cidDescription, setCidDescription] = React.useState("");
  const [cidImage, setCidImage] = React.useState("");
  const [cidText, setCidText] = React.useState("");
  const [showCid, setShowCid] = React.useState(false);

  const handleClick = async () => {};

  const { config } = usePrepareContractWrite({
    address: "0x91c66B42bAC7b0d28Dd350a98DE327d8B07A31ad",
    abi: AbiSapoBridge,
    functionName: "request",
    args: [cid],
  });
  const { data, isLoading, isSuccess, write } = useContractWrite(config);

  return (
    <Card
      direction={{ base: "column", sm: "row" }}
      overflowX="hidden"
      variant="outline"
      width={460}
    >
      <Stack>
        <CardBody>
          <Heading size="md">EasyOCR</Heading>

          <Text py="2">
            Transform your image to text. Cid should have text attribute
          </Text>
          <Input
            onChange={(e) => setCid(e.target.value)}
            placeholder="Enter CID to process"
          />
        </CardBody>

        <CardFooter pt={0} display={"flex"} justifyContent={"center"}>
          <Button
            variant="solid"
            size="sm"
            colorScheme="blue"
            disabled={!write}
            onClick={() => write?.()}
          >
            Process
          </Button>
        </CardFooter>
        {}
      </Stack>
    </Card>
  );
}
