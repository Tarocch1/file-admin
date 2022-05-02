import { Fragment, useState, useRef } from 'react'
import { Space, Button } from 'antd'
import {
  FolderAddOutlined,
  FileAddOutlined,
  CloudUploadOutlined,
} from '@ant-design/icons'

import InputModal from '@/components/InputModal'

export interface IActionProp {
  onMkdir: (target: string) => void
  onTouch: (target: string) => void
  onUpload: (file: File) => void
}

export default function Action({ onMkdir, onTouch, onUpload }: IActionProp) {
  const [showModal, setShowModal] = useState<boolean>(false)
  const [modalValue, setModalValue] = useState<string>('')
  const [modalMode, setModalMode] = useState<'mkdir' | 'touch'>('mkdir')

  const uploadRef = useRef<HTMLInputElement>(null)

  function modalOk() {
    if (modalMode === 'mkdir') {
      onMkdir(modalValue)
    } else {
      onTouch(modalValue)
    }
    closeModal()
  }

  function closeModal() {
    setShowModal(false)
    setModalValue('')
  }

  return (
    <Fragment>
      <Space style={{ marginBottom: 16 }}>
        <Button
          type="primary"
          icon={<FolderAddOutlined />}
          onClick={() => {
            setModalMode('mkdir')
            setShowModal(true)
          }}
        >
          New Folder
        </Button>
        <Button
          type="primary"
          icon={<FileAddOutlined />}
          onClick={() => {
            setModalMode('touch')
            setShowModal(true)
          }}
        >
          New File
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
        visible={showModal}
        title={modalMode === 'mkdir' ? 'New Folder' : 'New File'}
        placeholder={modalMode === 'mkdir' ? 'Folder Name' : 'File Name'}
        value={modalValue}
        onOk={modalOk}
        onCancel={closeModal}
        onChange={(value) => setModalValue(value)}
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
