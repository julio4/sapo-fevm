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

import { useJobContext } from "@/components/Context/JobContext";
import { useEffect, useState } from "react";
import { useContractRead } from "wagmi";

export default function SelectFromDaoInput() {
  const { allFiles, setAllFiles } = useJobContext();
  const [daoAddress, setdaoAddress] = useState("");
  const [daoInput, setdaoInput] = useState("");

  const [rawCids, setRawCids] = useState<string[]>([]);
  const [cids, setCids] = useState<string[]>([]);

  let { data, isError, isLoading } = useContractRead({
    address: daoAddress,
    abi: [
      {
        inputs: [],
        name: "getAllCids",
        outputs: [
          {
            internalType: "bytes[]",
            name: "",
            type: "bytes[]",
          },
        ],
        stateMutability: "view",
        type: "function",
      },
    ],
    functionName: "getAllCids",
    onSuccess(data) {
      setRawCids([...data]);
    },
  });

  function hex2a(h: any) {
    var hex: string = h.toString(); //force conversion
    var str = "";
    for (var i = 0; i < hex.length; i += 2)
      str += String.fromCharCode(parseInt(hex.substr(i, 2), 16));
    return str.substring(1);
  }

  const convertToCids = () => {
    let newCids: string[] = [];
    rawCids.forEach((raw) => {
      newCids.push(hex2a(raw));
    });
    setCids(newCids);
  };

  const storeCidType = async () => {
    let newAllFiles: File[] = [];
    for (const index in cids) {
      const cid = cids[index];
      let file = "https://ipfs.io/ipfs/" + cid;

      const fetchPromise = fetch(file, { method: "HEAD" });
      const timeoutPromise = new Promise((resolve, reject) => {
        setTimeout(() => {
          resolve({ cid: `ERROR: ${cid}`, size: null, type: null });
        }, 1000);
      });

      const result = await Promise.race([fetchPromise, timeoutPromise]);
      if (result instanceof Response) {
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
        newAllFiles.push(result);
      }
    }
    setAllFiles(newAllFiles);
  };

  useEffect(() => {
    convertToCids();
  }, [rawCids]);

  useEffect(() => {
    storeCidType();
  }, [cids]);

  return (
    <Flex pt={3}>
      <Flex h="full" w="33%" pr={3}>
        <Input
          onChange={(e) => setdaoInput(e.target.value)}
          placeholder="Enter DAO Address"
        ></Input>
      </Flex>
      <Button
        bgGradient={useColorModeValue(
          "linear(to-r, teal.100, green.100)",
          "linear(to-r, teal.900, green.900)"
        )}
        color={useColorModeValue("gray.700", "gray.200")}
        onClick={() => {
          if (daoInput == daoAddress) {
            setRawCids([...data]);
          }
          setdaoAddress(daoInput);
        }}
      >
        Load
      </Button>
    </Flex>
  );
}
