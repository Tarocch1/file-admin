package api

import (
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func MkdirHandler(c *kid.Ctx) error {
	var err error

	path := c.FormValue("path")
	target := c.FormValue("target")

	// 计算出工作路径
	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(workingPath, target)

	err = os.MkdirAll(targetPath, 0755)
	if err != nil {
		return err
	}

	return c.Json(common.JsonMap(c, nil))
}
