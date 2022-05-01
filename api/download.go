package api

import (
	"net/http"

	"github.com/Tarocch1/file-admin/common"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	target := r.FormValue("target")

	targetPath, err := common.GetWorkingPath(target)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+target+"\"")

	http.ServeFile(w, r, targetPath)
}