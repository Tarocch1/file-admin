package api

import (
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func CatHandler(c *kid.Ctx) error {
	var err error

	path := c.FormValue("path")
	target := c.FormValue("target")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(workingPath, target)

	fileBytes, err := os.ReadFile(targetPath)
	if err != nil {
		return err
	}

	return c.Json(common.JsonMap(c, map[string]interface{}{
		"content": string(fileBytes),
	}))
}
