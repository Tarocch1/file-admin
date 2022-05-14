package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/Tarocch1/file-admin/api"
	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/kid"
	"github.com/Tarocch1/kid/middlewares/basicauth"
	"github.com/Tarocch1/kid/middlewares/recovery"
	"github.com/Tarocch1/kid/middlewares/requestid"
	"github.com/Tarocch1/kid/middlewares/requestlogger"
)

var requestLogger = common.NewLogger("HTTP Request")

func startServer() error {
	k := kid.New(kid.Config{
		ErrorHandler: errorHandler,
	})

	k.Use(recovery.New())
	k.Use(requestid.New())
	k.Use(requestlogger.New(requestlogger.Config{
		UseFmt: true,
		Formatter: func(c *kid.Ctx) string {
			message := fmt.Sprintf("%s %s", c.Method(), c.Url().RequestURI())
			extra := map[string]interface{}{
				"body":   string(c.Body()),
				"header": c.Header(),
			}
			return requestLogger.FormatMessage(c, message, common.LoggerLevelInfo, extra, nil)
		},
	}))

	if common.AuthUser != "" && common.AuthPass != "" {
		k.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				common.AuthUser: common.AuthPass,
			},
		}))
	}

	apiGroup := k.Group("/api")
	apiGroup.Post("/ls", api.LsHandler)
	apiGroup.Post("/touch", api.TouchHandler)
	apiGroup.Post("/mkdir", api.MkdirHandler)
	apiGroup.Post("/cat", api.CatHandler)
	apiGroup.Post("/edit", api.EditHandler)
	apiGroup.Post("/mv", api.MvHandler)
	apiGroup.Post("/rm", api.RmHandler)
	apiGroup.Get("/download", api.DownloadHandler)
	apiGroup.Post("/upload", api.UploadHandler)

	fsRoot, _ := fs.Sub(static, "static")

	k.Get("/", func(c *kid.Ctx) error {
		return c.SendFile("index.html", false, http.FS(fsRoot))
	})
	k.Get("/*path", func(c *kid.Ctx) error {
		return c.SendFile(c.GetParam("path"), false, http.FS(fsRoot))
	})

	addr := fmt.Sprintf("%s:%s", common.FlagHost, common.FlagPort)

	return k.Listen(addr)
}

func errorHandler(c *kid.Ctx, err error) error {
	status := http.StatusInternalServerError

	if e, ok := err.(*kid.Error); ok {
		status = e.Status
	}

	err = c.Status(status).Json(common.ErrorMap(c, status, err))
	if err != nil {
		status = http.StatusInternalServerError
		return c.Status(status).Json(common.ErrorMap(c, status, err))
	}

	return nil
}
