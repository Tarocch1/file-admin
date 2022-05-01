import { request } from '@/utils/request'
import { LsResultItem } from '@/types'

export async function ls(path: string) {
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

export async function mkdir(path: string, dir: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('dir', dir)

  const res = await request('/api/mkdir', {
    method: 'POST',
    body: data,
  })

  if (res) {
    return true
  }
  return false
}

export async function rm(path: string, target: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)

  const res = await request('/api/rm', {
    method: 'POST',
    body: data,
  })

  if (res) {
    return true
  }
  return false
}
