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
import * as React from "react";

export default function OcrCard() {
  const [cid, setCid] = React.useState("");
  const [cidName, setCidName] = React.useState("");
  const [cidDescription, setCidDescription] = React.useState("");
  const [cidImage, setCidImage] = React.useState("");
  const [cidText, setCidText] = React.useState("");
  const [showCid, setShowCid] = React.useState(false);

  const handleClick = async () => {};

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
            placeholder="Enter CID Adress to process"
          />
        </CardBody>

        <CardFooter pt={0} display={"flex"} justifyContent={"center"}>
          <Button variant="solid" size="sm" colorScheme="blue">
            Process
          </Button>
        </CardFooter>
        {}
      </Stack>
    </Card>
  );
}
