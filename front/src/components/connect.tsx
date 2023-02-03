import {
  Heading,
  Flex,
} from '@chakra-ui/react';
import { ConnectButton } from '@rainbow-me/rainbowkit';

export default function Connect() {
  return (
    <Flex grow="1" justifyContent="center" alignItems="center">
      <ConnectButton />
    </Flex>
  );
}
 
