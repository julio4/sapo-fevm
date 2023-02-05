import { JobProvider } from '@/components/Context/JobContext'
import Layout from '@/components/Layout/layout'
import Sidebar from '@/components/Sidebars/JobsSidebar'
import JobRequestFlow from '@/components/Flow/JobRequestFlow'

import { IpfsProvider } from '@/services/ipfs'

export default function Home() {
  return (
    <IpfsProvider>
      <JobProvider>
        <Layout>
          <Sidebar />
          <JobRequestFlow />
        </Layout>
      </JobProvider>
    </IpfsProvider>
  )
}