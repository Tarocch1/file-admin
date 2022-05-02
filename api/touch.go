package api

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func TouchHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")
	target := r.FormValue("target")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)
	if !common.PathNotExist(target) {
		ErrorHandler(w, http.StatusConflict, errors.New("target already exists"))
		return
	}

	err = os.WriteFile(targetPath, []byte{}, 0644)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, map[string]interface{}{})
}
