import { useCallback } from 'react'
import { Breadcrumb, Table, Button, Popconfirm } from 'antd'
import { ColumnsType } from 'antd/es/table'
import {
  HomeOutlined,
  FolderOpenTwoTone,
  FileOutlined,
  DeleteOutlined,
} from '@ant-design/icons'
import dayjs from 'dayjs'

import { formatSize } from '@/utils'
import { LsResultItem } from '@/types'

export interface IDataTableProp {
  loading: boolean
  paths: string[]
  data: LsResultItem[]
  onSetPaths: (paths: string[]) => void
  onClickName: (item: LsResultItem) => void
  onRm: (item: LsResultItem) => void
}

export default function DataTable({
  loading,
  paths,
  data,
  onSetPaths,
  onClickName,
  onRm,
}: IDataTableProp) {
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
          <div>
            <Popconfirm
              title={`Are you sure to rm '${item.name}'?`}
              arrowPointAtCenter
              onConfirm={() => onRm(item)}
            >
              <Button size="small" icon={<DeleteOutlined />} danger />
            </Popconfirm>
          </div>
        )
      },
      width: 150,
    },
  ]

  return (
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
  )
}
