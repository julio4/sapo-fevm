import {
  Box,
  Flex,
  Text,
  useColorModeValue,
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  Td,
  Button,
  Heading,
  Input,
} from "@chakra-ui/react";

import { useJobContext, Job } from "@/components/Context/JobContext";
import { useEffect, useState } from "react";

export default function SelectFromDaoInput() {
  const { allFiles, setAllFiles } = useJobContext();
  const [cidInput, setCidInput] = useState("");
  const [cid, setCid] = useState<string>("");

  const storeCidType = async () => {
    let newAllFiles: File[] = [];
    let file = "https://ipfs.io/ipfs/" + cid;

    try {
      const fetchPromise = fetch(file, { method: "HEAD" });
      const timeoutPromise = new Promise((resolve, reject) => {
        setTimeout(() => {
          reject(new Error(`ERROR: ${cid}`));
        }, 2000);
      });
      const result = await Promise.race([fetchPromise, timeoutPromise]);
      if (result.ok) {
        let size = result.headers.get("content-length");
        size = size || null;
        let type = result.headers.get("content-type");
        type = type || "undefined";
        newAllFiles.push({
          cid,
          type,
          size,
        });
      } else {
        throw new Error(`ERROR: ${cid}`);
      }
    } catch (error) {
      newAllFiles.push({
        cid: `ERROR: ${cid}`,
        type: null,
        size: null,
      });
    }
    setAllFiles(newAllFiles);
  };

  useEffect(() => {
    if (cid) {
      storeCidType();
    }
  }, [cid, setAllFiles]);

  return (
    <Flex pt={3}>
      <Flex h="full" w="33%" pr={3}>
        <Input
          onChange={(e) => setCidInput(e.target.value)}
          placeholder="Enter CID"
        ></Input>
      </Flex>
      <Button
        bgGradient={useColorModeValue(
          "linear(to-r, teal.100, green.100)",
          "linear(to-r, teal.900, green.900)"
        )}
        onClick={() => {
          if (cidInput == cid) {
            storeCidType();
          }

          setCid(cidInput);
        }}
      >
        Load
      </Button>
    </Flex>
  );
}
