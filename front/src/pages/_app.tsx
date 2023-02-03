import type { AppProps } from 'next/app'

// Style
import { ChakraProvider, useColorModeValue } from '@chakra-ui/react'
import '@rainbow-me/rainbowkit/styles.css';

// Rainbow
import { RainbowKitProvider, getDefaultWallets, lightTheme, darkTheme } from '@rainbow-me/rainbowkit';
import { Chain, configureChains, createClient, WagmiConfig } from 'wagmi';

const filecoinHyperspace: Chain = {
  id: 3141,
  name: 'Filecoin Hyperspace',
  network: 'filecoin-hyperspace',
  nativeCurrency: {
    decimals: 18,
    name: 'testnet filecoin',
    symbol: 'tFIL',
  },
  rpcUrls: {
    default: {
      http: ['https://api.hyperspace.node.glif.io/rpc/v1'],
      webSocket: ['wss://wss.hyperspace.node.glif.io/apigw/lotus/rpc/v0']
    },
    public: {
      http: ['https://api.hyperspace.node.glif.io/rpc/v1'],
      webSocket: ['wss://wss.hyperspace.node.glif.io/apigw/lotus/rpc/v0']
    },
  },
  blockExplorers: {
    default: { name: 'Filfox', url: 'https://hyperspace.filfox.info/en' },
    gilf: { name: 'Glif', url: 'https://explorer.glif.io/?network=hyperspace' },
  },
}

// should work in next version
// import { filecoinHyperspace } from 'wagmi/chains';
import { publicProvider } from 'wagmi/providers/public';

const { chains, provider, webSocketProvider } = configureChains(
  [
    filecoinHyperspace
  ],
  [
    publicProvider(),
  ]
);

const { connectors } = getDefaultWallets({
  appName: 'Sapo',
  chains,
});

const wagmiClient = createClient({
  autoConnect: true,
  connectors,
  provider,
  webSocketProvider,
});

function App({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider>
      <WagmiConfig client={wagmiClient}>
        <RainbowKitProvider
          chains={chains}
          //theme={useColorModeValue(lightTheme, darkTheme)}
          >
          <Component {...pageProps} />
        </RainbowKitProvider>
      </WagmiConfig>
    </ChakraProvider>
  )
}

export default App
