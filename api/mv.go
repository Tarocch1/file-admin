package api

import (
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func MvHandler(c *kid.Ctx) error {
	var err error

	path := c.FormValue("path")
	target := c.FormValue("target")
	to := c.FormValue("to")

	// 计算出工作路径
	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(workingPath, target)
	toPath := filepath.Join(workingPath, to)

	err = os.Rename(targetPath, toPath)
	if err != nil {
		return err
	}

	return c.Json(common.JsonMap(c, nil))
}
