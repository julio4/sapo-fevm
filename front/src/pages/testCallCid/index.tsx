import Layout from "../../components/layout";
import DaoNavBar from "../../components/daoNavBar";
import ZoomCard from "../../components/zoomCard";
import OcrCard from "../../components/OcrCard";
import { Card, CardBody } from "@chakra-ui/react";

import * as React from "react";
import { Flex, Heading, Button, Text, Input } from "@chakra-ui/react";
import { useContractRead } from "wagmi";

export default function Home() {
  return (
    <Layout>
      <Flex w="25%">
        <DaoNavBar />
      </Flex>
      <Flex w="50%" direction="column" p={12} rounded={6} alignItems="center">
        <Heading mb={6}>Jobs</Heading>
        <ZoomCard />
        <OcrCard />
      </Flex>
    </Layout>
  );
}
