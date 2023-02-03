import Layout from '../../components/layout'
import Sidebar from '../../components/Sidebars/JobsSidebar'
import FileSelector from '../../components/Selectors/FileSelector'
import DaoSelector from '../../components/Sidebars/DaoSidebar'

export default function Home() {
  return (
    <Layout>
      <Sidebar />
      <FileSelector />
      <DaoSelector />
    </Layout>
  )
}
