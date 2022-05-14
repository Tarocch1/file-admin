import { request } from '@/utils/request'
import { LsResultItem } from '@/types'

export async function ls(path: string) {
  const data = new FormData()
  data.append('path', path)

  const res = await request<LsResultItem[]>('/api/ls', {
    method: 'POST',
    body: data,
  })

  if (res) {
    return res.data
  }
  return []
}

export async function touch(path: string, target: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)

  const res = await request<null>(
    '/api/touch',
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

export async function mkdir(path: string, target: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)

  const res = await request<null>(
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

export async function cat(path: string, target: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)

  const res = await request<{ content: string }>('/api/cat', {
    method: 'POST',
    body: data,
  })

  if (res) {
    return res.data
  }
  return false
}

export async function edit(path: string, target: string, content: Blob) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)
  data.append('content', content)

  const res = await request<null>(
    '/api/edit',
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

export async function mv(path: string, target: string, to: string) {
  const data = new FormData()
  data.append('path', path)
  data.append('target', target)
  data.append('to', to)

  const res = await request<null>(
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

  const res = await request<null>(
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

  const res = await request<null>(
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
