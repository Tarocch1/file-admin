import { Fragment, useState, useRef } from 'react'
import { Space, Button } from 'antd'
import { FolderAddOutlined, CloudUploadOutlined } from '@ant-design/icons'

import InputModal from '@/components/InputModal'

export interface IActionProp {
  onMkdir: (dir: string) => void
  onUpload: (file: File) => void
}

export default function Action({ onMkdir, onUpload }: IActionProp) {
  const [showMkdir, setShowMkdir] = useState<boolean>(false)
  const [mkdirValue, setMkdirValue] = useState<string>('')

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
      <InputModal
        visible={showMkdir}
        title="New Folder"
        placeholder="Folder Path"
        value={mkdirValue}
        onOk={mkdir}
        onCancel={closeMkdir}
        onChange={(value) => setMkdirValue(value)}
      />
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
