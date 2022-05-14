package api

import (
	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func DownloadHandler(c *kid.Ctx) error {
	target := c.FormValue("target")

	return c.SendFile(target, true, &kid.FileSystem{
		Root: common.RootDir,
	})
}
