package common

import (
	"fmt"
	"time"

	"github.com/Tarocch1/kid"
	"github.com/Tarocch1/kid/middlewares/requestid"
)

var responseLogger = NewLogger("HTTP Response")
var errorLogger = NewLogger("HTTP Error")

func JsonMap(c *kid.Ctx, data interface{}) map[string]interface{} {
	message := fmt.Sprintf("%s %s", c.Method(), c.Url().RequestURI())
	extra := map[string]interface{}{
		"data": data,
	}
	responseLogger.Info(c, message, extra)

	return map[string]interface{}{
		"code":      0,
		"message":   "success",
		"data":      data,
		"requestId": c.Get(requestid.HeaderRequestId),
		"timestamp": time.Now().Unix(),
	}
}

func ErrorMap(c *kid.Ctx, status int, err error) map[string]interface{} {
	message := fmt.Sprintf("%s %s", c.Method(), c.Url().RequestURI())
	extra := map[string]interface{}{
		"status": status,
	}
	errorLogger.Error(c, message, extra, err)

	return map[string]interface{}{
		"code":      status,
		"message":   err.Error(),
		"requestId": c.Get(requestid.HeaderRequestId),
		"timestamp": time.Now().Unix(),
	}
}
