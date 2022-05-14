package api

import (
	"io/ioutil"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

type LsResultItem struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Time  int64  `json:"time"`
	Mode  string `json:"mode"`
	Size  int64  `json:"size"`
}

func LsHandler(c *kid.Ctx) error {
	var err error

	path := c.FormValue("path")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		return err
	}

	allItems, err := ioutil.ReadDir(workingPath)
	if err != nil {
		return err
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

	return c.Json(common.JsonMap(c, result))
}
