package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")
	target := r.FormValue("target")
	content := r.FormValue("content")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)

	err = os.WriteFile(targetPath, []byte(content), 0644)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, map[string]interface{}{})
}
