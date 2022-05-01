import { request } from '@/utils/request'
import { LsResultItem } from '@/types'

export async function listDir(path: string) {
  const data = new FormData()
  data.append('path', path)

  const res = await request('/api/ls', {
    method: 'POST',
    body: data,
  })

  if (res) {
    return res as LsResultItem[]
  }
  return []
}
