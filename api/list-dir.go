package api

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Tarocch1/file-admin/common"
)

type ListDirResultItem struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Time  int64  `json:"time"`
	Size  int64  `json:"size"`
}

func ListDirHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// 获取 path 参数
	path := r.FormValue("path")

	// 计算出目标路径
	targetPath, err := common.GetTargetPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	// 判断目标路径是否存在
	if common.PathNotExist(targetPath) {
		ErrorHandler(w, http.StatusNotFound, errors.New("path is invalid"))
		return
	}
	// 判断目标路径是否是目录
	if !common.PathIsDir(targetPath) {
		ErrorHandler(w, http.StatusBadRequest, errors.New("path is not dir"))
		return
	}

	// 读取文件
	allItems, err := ioutil.ReadDir(targetPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	// 分别保存文件子目录
	var result, dirs, files []ListDirResultItem
	for _, itemInfo := range allItems {
		if itemInfo.IsDir() {
			dirs = append(dirs, ListDirResultItem{
				Name:  itemInfo.Name(),
				IsDir: itemInfo.IsDir(),
				Time:  itemInfo.ModTime().Unix(),
				Size:  itemInfo.Size(),
			})
		} else {
			files = append(files, ListDirResultItem{
				Name:  itemInfo.Name(),
				IsDir: itemInfo.IsDir(),
				Time:  itemInfo.ModTime().Unix(),
				Size:  itemInfo.Size(),
			})
		}
	}
	// 子目录排在文件前面
	result = append(result, dirs...)
	result = append(result, files...)

	JsonHandler(w, result)
}
