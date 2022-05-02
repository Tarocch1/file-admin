package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func CatHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")
	target := r.FormValue("target")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, target)

	fileBytes, err := os.ReadFile(targetPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, map[string]interface{}{
		"content": string(fileBytes),
	})
}
