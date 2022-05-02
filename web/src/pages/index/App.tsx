import { useState, useEffect, useCallback } from 'react'

import {
  ls as lsApi,
  mkdir as mkdirApi,
  touch as touchApi,
  edit as editApt,
  mv as mvApi,
  rm as rmApi,
  upload as uploadApi,
} from '@/api'
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

  const touch = useCallback(
    async (target: string) => {
      const res = await touchApi(paths.join('/'), target)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const mkdir = useCallback(
    async (target: string) => {
      const res = await mkdirApi(paths.join('/'), target)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const edit = useCallback(
    async (target: string, value: string) => {
      const res = await editApt(paths.join('/'), target, value)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const mv = useCallback(
    async (target: string, to: string) => {
      const res = await mvApi(paths.join('/'), target, to)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const rm = useCallback(
    async (target: string) => {
      const res = await rmApi(paths.join('/'), target)
      if (res) {
        ls()
      }
    },
    [paths, ls]
  )

  const upload = useCallback(
    async (file: File) => {
      const res = await uploadApi(paths.join('/'), file)
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
        window.open(
          `/api/download?target=${[...paths, item.name].join('/')}`,
          '_blank',
          'noopener=yes,noreferrer=yes'
        )
      }
    },
    [paths]
  )

  useEffect(() => {
    ls()
  }, [ls])

  return (
    <div>
      <Action onMkdir={mkdir} onTouch={touch} onUpload={upload} />
      <DataTable
        loading={loading}
        paths={paths}
        data={data}
        onSetPaths={setPaths}
        onClickName={onClickName}
        onEdit={edit}
        onMv={mv}
        onRm={rm}
      />
    </div>
  )
}
