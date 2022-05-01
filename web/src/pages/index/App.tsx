import { useState, useEffect, useCallback } from 'react'

import { ls as lsApi, mkdir as mkdirApi } from '@/api'
import { LsResultItem } from '@/types'
import Action from './components/Action'
import DataTable from './components/DataTable'

export default function App() {
  const [loading, setLoading] = useState<boolean>(false)
  const [paths, setPaths] = useState<string[]>([])
  const [data, setData] = useState<LsResultItem[]>([])

  const ls = useCallback(async () => {
    setLoading(true)
    const res = await lsApi(paths.join('/'))
    setData(res)
    setLoading(false)
  }, [paths])

  const mkdir = useCallback(
    async (dir: string) => {
      const res = await mkdirApi(paths.join('/'), dir)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const onClickName = useCallback(
    (row: LsResultItem) => {
      if (row.isDir) {
        setPaths([...paths, row.name])
      } else {
        // TODO download
      }
    },
    [paths]
  )

  useEffect(() => {
    ls()
  }, [ls])

  return (
    <div>
      <Action onMkdir={mkdir} />
      <DataTable
        loading={loading}
        paths={paths}
        data={data}
        onSetPaths={setPaths}
        onClickName={onClickName}
      />
    </div>
  )
}
