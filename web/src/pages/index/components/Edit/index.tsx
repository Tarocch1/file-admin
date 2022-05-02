import { useState, useEffect, useCallback } from 'react'

import { cat as catApi } from '@/api'
import InputModal from '@/components/InputModal'
import Monaco from './Monaco'
import './utils/useMonacoWorker'

export interface IEditProp {
  paths: string[]
  target: string
  onOk: (value: Blob) => void
  onCancel: () => void
}

export default function Edit({ paths, target, onOk, onCancel }: IEditProp) {
  const [show, setShow] = useState<boolean>(false)
  const [value, setValue] = useState<string>('')

  const cat = useCallback(async () => {
    if (!target) {
      setValue('')
      setShow(false)
      return
    }
    const res = await catApi(paths.join('/'), target)
    if (res) {
      setValue(res.content)
    } else {
      setValue('')
    }
    setShow(true)
  }, [paths, target])

  useEffect(() => {
    cat()
  }, [cat])

  return (
    <InputModal
      visible={show}
      title="Edit"
      width="80%"
      value={value}
      centered
      component={({ value, onChange }) => (
        <Monaco
          value={value}
          target={target}
          onChange={(value) => onChange(value)}
        />
      )}
      onOk={() => onOk(new Blob([value]))}
      onCancel={onCancel}
      onChange={(value) => setValue(value)}
    />
  )
}
