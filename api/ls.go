package api

import (
	"io/ioutil"
	"net/http"

	"github.com/Tarocch1/file-admin/common"
)

type LsResultItem struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Time  int64  `json:"time"`
	Mode  string `json:"mode"`
	Size  int64  `json:"size"`
}

func LsHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	path := r.FormValue("path")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	allItems, err := ioutil.ReadDir(workingPath)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	var result, dirs, files []LsResultItem
	for _, itemInfo := range allItems {
		if itemInfo.IsDir() {
			dirs = append(dirs, LsResultItem{
				Name:  itemInfo.Name(),
				IsDir: itemInfo.IsDir(),
				Time:  itemInfo.ModTime().Unix(),
				Mode:  itemInfo.Mode().String(),
				Size:  itemInfo.Size(),
			})
		} else {
			files = append(files, LsResultItem{
				Name:  itemInfo.Name(),
				IsDir: itemInfo.IsDir(),
				Time:  itemInfo.ModTime().Unix(),
				Mode:  itemInfo.Mode().String(),
				Size:  itemInfo.Size(),
			})
		}
	}
	result = append(result, dirs...)
	result = append(result, files...)

	JsonHandler(w, result)
}
