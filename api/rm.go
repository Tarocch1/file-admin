package api

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func RmHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// 获取 path 参数
	path := r.FormValue("path")
	target := r.FormValue("target")

	// 计算出工作路径
	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)
	if common.PathNotExist(targetPath) {
		ErrorHandler(w, http.StatusNotFound, errors.New("target not exists"))
		return
	}

	err = os.RemoveAll(targetPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, make(map[string]interface{}))
}
