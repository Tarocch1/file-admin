import { useState, useEffect, useCallback, Fragment } from 'react'
import { Breadcrumb, Table } from 'antd'
import { ColumnsType } from 'antd/es/table'
import {
  HomeOutlined,
  FolderOpenTwoTone,
  FileOutlined,
} from '@ant-design/icons'
import dayjs from 'dayjs'

import { listDir } from '@/api'
import { formatSize } from '@/utils'
import { ListDirResultItem } from '@/types'

export default function App() {
  const [loading, setLoading] = useState<boolean>(false)
  const [paths, setPaths] = useState<string[]>([])
  const [data, setData] = useState<ListDirResultItem[]>([])

  useEffect(() => {
    setLoading(true)
    listDir(paths.join('/')).then((res) => {
      setData(res)
      setLoading(false)
    })
  }, [paths])

  const onClickName = useCallback(
    (row: ListDirResultItem) => {
      if (row.isDir) {
        setPaths([...paths, row.name])
      } else {
        // TODO download
      }
    },
    [paths]
  )

  const columns: ColumnsType<ListDirResultItem> = [
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
      render: (name: string, record) => {
        return <a onClick={() => onClickName(record)}>{name}</a>
      },
    },
    {
      key: 'name',
      title: 'Time',
      dataIndex: 'time',
      render: (time: number) => {
        return dayjs(time * 1000).format('YYYY-MM-DD HH:mm:ss')
      },
    },
    {
      key: 'name',
      title: 'Size',
      dataIndex: 'size',
      render: formatSize,
    },
  ]

  return (
    <Fragment>
      <div style={{ marginBottom: 16 }}>
        <Breadcrumb>
          <Breadcrumb.Item href="#" onClick={() => setPaths([])}>
            <HomeOutlined />
          </Breadcrumb.Item>
          {paths.map((path, i) => (
            <Breadcrumb.Item
              key={`${path}${i}`}
              href="#"
              onClick={() => setPaths(paths.slice(0, i + 1))}
            >
              {path}
            </Breadcrumb.Item>
          ))}
        </Breadcrumb>
      </div>
      <div>
        <Table
          bordered
          rowKey="name"
          size="small"
          pagination={false}
          loading={loading}
          dataSource={data}
          columns={columns}
        />
      </div>
    </Fragment>
  )
}
