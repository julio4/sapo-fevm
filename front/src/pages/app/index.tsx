import { JobProvider } from '@/components/Context/JobContext'
import Layout from '@/components/layout'
import Sidebar from '@/components/Sidebars/JobsSidebar'
import JobRequestFlow from '@/components/Flow/JobRequestFlow'

export default function Home() {
  return (
    <JobProvider>
      <Layout>
        <Sidebar />
        <JobRequestFlow />
      </Layout>
    </JobProvider>
  )
}
