package api

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
)

func DownloadHandler(c *kid.Ctx) error {
	var err error

	target := c.FormValue("target")

	targetPath, err := common.GetWorkingPath(target)
	if err != nil {
		return err
	}

	return c.SendFile(targetPath, true, http.FS(&myfs{}))
}

type myfs struct{}

func (f *myfs) Open(path string) (fs.File, error) {
	return os.Open("/" + path)
}
