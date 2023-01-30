import {
  Heading,
  Flex,
} from '@chakra-ui/react';
import { ConnectButton } from '@rainbow-me/rainbowkit';

export default function Connect() {
  return (
    <Flex height="75vh" justifyContent="center" alignItems="center">
      <ConnectButton />
    </Flex>
  );
}
 