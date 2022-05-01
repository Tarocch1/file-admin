import { request } from '@/utils/request'
import { ListDirResultItem } from '@/types'

export async function listDir(path: string) {
  const data = new FormData()
  data.append('path', path)

  const res = await request('/api/list-dir', {
    method: 'POST',
    body: data,
  })

  if (res) {
    return res as ListDirResultItem[]
  }
  return []
}
