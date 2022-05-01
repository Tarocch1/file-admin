package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func RmHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")
	target := r.FormValue("target")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)

	err = os.RemoveAll(targetPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, make(map[string]interface{}))
}
