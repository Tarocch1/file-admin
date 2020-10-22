# File Server

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/tarocch1/file-server)](https://github.com/Tarocch1/file-server/releases)
[![GitHub All Releases](https://img.shields.io/github/downloads/tarocch1/file-server/total)](https://github.com/Tarocch1/file-server/releases)
[![GitHub](https://img.shields.io/github/license/tarocch1/file-server)](https://github.com/Tarocch1/file-server/blob/master/LICENSE)

一个简易的文件服务器。

## Feature

- 基于 HTTP 协议的文件管理，包括文件列表、上传、更新、删除以及新建文件夹。
- 支持 HTTPS。
- 支持 HTTP Basic 认证。
- 单个二进制文件，方便部署。

## Usage

```bash
$ file-server --help
Usage of file-server:
  -a string
        <username:password> Basic auth user.
  -d string
        Dir path to serve. (default ".")
  -h string
        Host to listen. (default "0.0.0.0")
  -https-cert string
        Path to https cert.
  -https-key string
        Path to https key.
  -p string
        Port to listen. (default "8080")
```
