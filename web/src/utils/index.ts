/**
 * 格式化文件大小
 * @param {number} size
 * @returns {string}
 */
export function formatSize(size: number) {
  const units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB']
  let _size = size
  let index = 0

  while (_size >= 1024 && index < 5) {
    _size /= 1024
    index++
  }

  return `${_size.toFixed(2)}${units[index]}`
}
