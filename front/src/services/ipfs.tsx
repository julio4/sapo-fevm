import React, { createContext, useState, useContext, useEffect } from "react";
import { create, IPFS } from 'ipfs-core'

type IpfsContextType = {
  id: string,
  ipfs: IPFS | null,
  version: string,
  isOnline: boolean,
};

const IpfsContext = createContext<IpfsContextType>({
  id: "",
  ipfs: null,
  version: "",
  isOnline: false,
});

const useIpfsContext = () => useContext(IpfsContext);

const IpfsProvider = ({ children }: { children: React.ReactNode }) => {
  const [id, setId] = useState<string>("");
  const [ipfs, setIpfs] = useState<IPFS | null>(null);
  const [version, setVersion] = useState<string>("");
  const [isOnline, setIsOnline] = useState<boolean>(false);

  useEffect(() => {
    const init = async () => {
      if (ipfs) return

      const node = await create({repo: 'sapo' + Math.random()});

      const nodeId = await node.id();
      const nodeVersion = await node.version();
      const nodeIsOnline = node.isOnline();

      setIpfs(node);
      setId(nodeId.id.toString());
      setVersion(nodeVersion.version);
      setIsOnline(nodeIsOnline);
    }

    init()
  }, [ipfs]);

  return (
    <IpfsContext.Provider
      value={{
        id, ipfs, version, isOnline
      }}
    >
      {children}
    </IpfsContext.Provider>
  );
};



export { useIpfsContext, IpfsProvider };