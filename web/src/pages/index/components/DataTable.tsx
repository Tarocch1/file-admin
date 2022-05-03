import { Fragment, useState, useCallback } from 'react'
import { Breadcrumb, Table, Space, Button, Popconfirm } from 'antd'
import { ColumnsType } from 'antd/es/table'
import {
  HomeOutlined,
  FolderOpenTwoTone,
  FileOutlined,
  DeleteOutlined,
  FontSizeOutlined,
  EditOutlined,
} from '@ant-design/icons'
import dayjs from 'dayjs'

import { formatSize } from '@/utils'
import { LsResultItem } from '@/types'
import InputModal from '@/components/InputModal'
import Edit from './Edit'

export interface IDataTableProp {
  loading: boolean
  paths: string[]
  data: LsResultItem[]
  onSetPaths: (paths: string[]) => void
  onClickName: (item: LsResultItem) => void
  onMv: (target: string, to: string) => void
  onRm: (target: string) => void
  onEdit: (target: string, value: Blob) => void
}

export default function DataTable({
  loading,
  paths,
  data,
  onSetPaths,
  onClickName,
  onMv,
  onRm,
  onEdit,
}: IDataTableProp) {
  const [showMv, setShowMv] = useState<boolean>(false)
  const [mvTarget, setMvTarget] = useState<string>('')
  const [mvValue, setMvValue] = useState<string>('')
  const [editTarget, setEditTarget] = useState<string>('')

  function mv() {
    onMv(mvTarget, mvValue)
    closeMv()
  }

  function closeMv() {
    setShowMv(false)
    setMvValue('')
    setMvTarget('')
  }

  function edit(value: Blob) {
    onEdit(editTarget, value)
    setEditTarget('')
  }

  function closeEdit() {
    setEditTarget('')
  }

  const renderTitle = useCallback(() => {
    return (
      <Breadcrumb>
        <Breadcrumb.Item href="#" onClick={() => onSetPaths([])}>
          <HomeOutlined />
        </Breadcrumb.Item>
        {paths.map((path, i) => (
          <Breadcrumb.Item
            key={`${path}${i}`}
            href="#"
            onClick={() => onSetPaths(paths.slice(0, i + 1))}
          >
            {path}
          </Breadcrumb.Item>
        ))}
      </Breadcrumb>
    )
  }, [paths])

  const columns: ColumnsType<LsResultItem> = [
    {
      key: 'name',
      title: '',
      dataIndex: 'isDir',
      render: (isDir: boolean) => {
        return isDir ? (
          <FolderOpenTwoTone twoToneColor="#d48806" />
        ) : (
          <FileOutlined />
        )
      },
      width: 10,
    },
    {
      key: 'name',
      title: 'Name',
      dataIndex: 'name',
      render: (name: string, item) => {
        return <a onClick={() => onClickName(item)}>{name}</a>
      },
    },
    {
      key: 'name',
      title: 'Time',
      dataIndex: 'time',
      render: (time: number) => {
        return dayjs(time * 1000).format('YYYY-MM-DD HH:mm:ss')
      },
      width: 200,
    },
    {
      key: 'name',
      title: 'Mode',
      dataIndex: 'mode',
      width: 150,
    },
    {
      key: 'name',
      title: 'Size',
      dataIndex: 'size',
      render: formatSize,
      width: 150,
    },
    {
      key: 'name',
      title: 'Operation',
      render: (item: LsResultItem) => {
        return (
          <Space size="small">
            <Button
              size="small"
              icon={<FontSizeOutlined />}
              title="Rename"
              onClick={() => {
                setMvTarget(item.name)
                setMvValue(item.name)
                setShowMv(true)
              }}
            />
            {!item.isDir && (
              <Button
                size="small"
                icon={<EditOutlined />}
                title="Edit"
                onClick={() => {
                  setEditTarget(item.name)
                }}
              />
            )}
            <Popconfirm
              title={`Are you sure to rm '${item.name}'?`}
              placement="topRight"
              arrowPointAtCenter
              onConfirm={() => onRm(item.name)}
            >
              <Button
                size="small"
                icon={<DeleteOutlined />}
                danger
                title="Delete"
              />
            </Popconfirm>
          </Space>
        )
      },
      width: 150,
    },
  ]

  return (
    <Fragment>
      <Table
        bordered
        rowKey="name"
        size="small"
        title={renderTitle}
        pagination={false}
        loading={loading}
        dataSource={data}
        columns={columns}
      />
      <InputModal
        visible={showMv}
        title="Rename"
        placeholder="New Name"
        value={mvValue}
        onOk={mv}
        onCancel={closeMv}
        onChange={(value) => setMvValue(value)}
      />
      <Edit
        paths={paths}
        target={editTarget}
        onOk={edit}
        onCancel={closeEdit}
      />
    </Fragment>
  )
}
