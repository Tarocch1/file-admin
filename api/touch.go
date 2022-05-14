package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func TouchHandler(c *kid.Ctx) error {
	var err error

	path := c.FormValue("path")
	target := c.FormValue("target")

	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(workingPath, target)
	if !common.PathNotExist(target) {
		return kid.NewError(http.StatusConflict, "target already exists")
	}

	err = os.WriteFile(targetPath, []byte{}, 0644)
	if err != nil {
		return err
	}

	return c.Json(common.JsonMap(c, nil))
}
