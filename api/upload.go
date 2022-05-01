package api

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")
	target := r.FormValue("target")
	file, _, err := r.FormFile("file")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)

	err = ioutil.WriteFile(targetPath, fileBytes, 0755)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, make(map[string]interface{}))
}
