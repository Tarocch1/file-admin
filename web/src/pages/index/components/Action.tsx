import { Fragment, useState, useMemo } from 'react'
import { Space, Button, Modal, Input } from 'antd'
import { FolderAddOutlined } from '@ant-design/icons'

export interface IActionProp {
  onMkdir: (dir: string) => Promise<void>
}

export default function Action({ onMkdir }: IActionProp) {
  const [mkdirValue, setMkdirValue] = useState<string>('')
  const [showMkdir, setShowMkdir] = useState<boolean>(false)
  const [mkdirLoading, setMkdirLoading] = useState<boolean>(false)

  const mkdirDisabled = useMemo(() => mkdirValue === '', [mkdirValue])

  async function mkdir() {
    setMkdirLoading(true)
    await onMkdir(mkdirValue)
    closeMkdir()
  }

  function closeMkdir() {
    setMkdirLoading(false)
    setShowMkdir(false)
    setMkdirValue('')
  }

  return (
    <Fragment>
      <Space style={{ marginBottom: 16 }}>
        <Button
          type="primary"
          icon={<FolderAddOutlined />}
          onClick={() => setShowMkdir(true)}
        >
          New Folder
        </Button>
      </Space>
      <Modal
        visible={showMkdir}
        title="New Folder"
        confirmLoading={mkdirLoading}
        okButtonProps={{
          disabled: mkdirDisabled,
        }}
        onOk={mkdir}
        onCancel={closeMkdir}
        destroyOnClose
      >
        <Input
          placeholder="Folder Path"
          value={mkdirValue}
          autoFocus
          onChange={(e) => setMkdirValue(e.target.value)}
        />
      </Modal>
    </Fragment>
  )
}
