import { useState, useEffect, useCallback } from 'react'

import { ls as lsApi, mkdir as mkdirApi, mv as mvApi, rm as rmApi } from '@/api'
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

  const mv = useCallback(
    async (item: LsResultItem, target: string) => {
      const res = await mvApi(paths.join('/'), item.name, target)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const rm = useCallback(
    async (item: LsResultItem) => {
      const res = await rmApi(paths.join('/'), item.name)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const onClickName = useCallback(
    (item: LsResultItem) => {
      if (item.isDir) {
        setPaths([...paths, item.name])
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
        onMv={mv}
        onRm={rm}
      />
    </div>
  )
}
