import { message } from 'antd'
import nProgress from 'nprogress'

export async function request<T>(
  url: string,
  init: RequestInit,
  showSuccess = false
) {
  try {
    nProgress.start()
    const res = await fetch(url, init)
    if (res.ok) {
      showSuccess && message.success('success')
      return (await res.json()) as T
    } else {
      console.log(res)
      message.error(url)
      return null
    }
  } catch (error) {
    console.error(error)
    message.error(url)
    return null
  } finally {
    nProgress.done()
  }
}
