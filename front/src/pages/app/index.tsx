import Layout from '../../components/layout'
import Sidebar from '../../components/Sidebars/JobsSidebar'
import FileSelector from '../../components/Selectors/FileSelector'

export default function Home() {
  return (
    <Layout>
      <Sidebar />
      <FileSelector />
    </Layout>
  )
}
