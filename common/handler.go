package common

import (
	"fmt"
	"time"

	"github.com/Tarocch1/kid"
)

var responseLogger = kid.NewLogger("HTTP Response")
var errorLogger = kid.NewLogger("HTTP Error")

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
		"requestId": c.Get(kid.CtxRequestId),
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
		"requestId": c.Get(kid.CtxRequestId),
		"timestamp": time.Now().Unix(),
	}
}
