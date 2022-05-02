import { useMemo } from 'react'
import { Modal, Input } from 'antd'

export interface IInputModalSubComponentProp {
  value: string
  onChange: (value: string) => void
}

export interface IInputModalProp {
  visible: boolean
  title: string
  width?: string | number
  centered?: boolean
  placeholder?: string
  value: string
  component?: (props: IInputModalSubComponentProp) => JSX.Element
  onOk: () => void
  onCancel: () => void
  onChange: (value: string) => void
}

export default function InputModal({
  visible,
  title,
  width,
  centered,
  placeholder,
  value,
  component,
  onOk,
  onCancel,
  onChange,
}: IInputModalProp) {
  const disabled = useMemo(() => value === '', [value])

  return (
    <Modal
      visible={visible}
      title={title}
      width={width}
      centered={centered}
      okButtonProps={{
        disabled,
      }}
      onOk={onOk}
      onCancel={onCancel}
      destroyOnClose
    >
      {component &&
        component({
          value,
          onChange: (value) => onChange(value),
        })}
      {!component && (
        <Input
          value={value}
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          autoFocus
        />
      )}
    </Modal>
  )
}
