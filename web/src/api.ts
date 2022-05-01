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

  const res = await request(
    '/api/mkdir',
    {
      method: 'POST',
      body: data,
    },
    true
  )

  if (res) {
    return true
  }
  return false
}

export async function mv(path: string, from: string, target: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('from', from)
  data.append('target', target)

  const res = await request(
    '/api/mv',
    {
      method: 'POST',
      body: data,
    },
    true
  )

  if (res) {
    return true
  }
  return false
}

export async function rm(path: string, target: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)

  const res = await request(
    '/api/rm',
    {
      method: 'POST',
      body: data,
    },
    true
  )

  if (res) {
    return true
  }
  return false
}

export async function upload(path: string, file: File) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', file.name)
  data.append('file', file)

  const res = await request(
    '/api/upload',
    {
      method: 'POST',
      body: data,
    },
    true
  )

  if (res) {
    return true
  }
  return false
}
