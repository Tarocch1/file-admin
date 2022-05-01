import { Fragment, useState, useMemo, useRef } from 'react'
import { Space, Button, Modal, Input } from 'antd'
import { FolderAddOutlined, CloudUploadOutlined } from '@ant-design/icons'

export interface IActionProp {
  onMkdir: (dir: string) => void
  onUpload: (file: File) => void
}

export default function Action({ onMkdir, onUpload }: IActionProp) {
  const [mkdirValue, setMkdirValue] = useState<string>('')
  const [showMkdir, setShowMkdir] = useState<boolean>(false)

  const mkdirDisabled = useMemo(() => mkdirValue === '', [mkdirValue])

  const uploadRef = useRef<HTMLInputElement>(null)

  async function mkdir() {
    onMkdir(mkdirValue)
    closeMkdir()
  }

  function closeMkdir() {
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
        <Button
          type="primary"
          icon={<CloudUploadOutlined />}
          onClick={() => {
            uploadRef.current?.click()
          }}
        >
          Upload File
        </Button>
      </Space>
      <Modal
        visible={showMkdir}
        title="New Folder"
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
      <input
        ref={uploadRef}
        type="file"
        multiple={false}
        style={{ display: 'none' }}
        onChange={(e) => {
          onUpload(e.target.files![0])
        }}
      />
    </Fragment>
  )
}
