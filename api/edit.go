package api

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func EditHandler(c *kid.Ctx) error {
	var err error

	path := c.FormValue("path")
	target := c.FormValue("target")
	fileHeader, err := c.FormFile("content")
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(workingPath, target)

	err = os.WriteFile(targetPath, fileBytes, 0644)
	if err != nil {
		return err
	}

	return c.Json(common.JsonMap(c, nil))
}
