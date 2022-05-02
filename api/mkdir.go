package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func MkdirHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")
	target := r.FormValue("target")

	// 计算出工作路径
	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)

	err = os.MkdirAll(targetPath, 0755)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, map[string]interface{}{})
}
