import { useState, useEffect, useCallback } from 'react'
import { Input } from 'antd'

import { cat as catApi } from '@/api'
import InputModal from '@/components/InputModal'

export interface IEditProp {
  visible: boolean
  paths: string[]
  target: string
  onOk: (value: string) => void
  onCancel: () => void
}

export default function Edit({
  visible,
  paths,
  target,
  onOk,
  onCancel,
}: IEditProp) {
  const [value, setValue] = useState<string>('')

  const cat = useCallback(async () => {
    if (!target) {
      setValue('')
      return
    }
    const res = await catApi(paths.join(','), target)
    if (res) {
      setValue(res.content)
    } else {
      setValue('')
    }
  }, [paths, target])

  useEffect(() => {
    cat()
  }, [cat])

  return (
    <InputModal
      visible={visible}
      title="Edit"
      width="80%"
      value={value}
      centered
      component={({ value, onChange }) => (
        <Input.TextArea
          rows={25}
          value={value}
          autoFocus
          onChange={(e) => onChange(e.target.value)}
        />
      )}
      onOk={() => onOk(value)}
      onCancel={onCancel}
      onChange={(value) => setValue(value)}
    />
  )
}
