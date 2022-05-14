package common

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Tarocch1/kid"
)

type LoggerLevel string

const (
	LoggerLevelInfo  LoggerLevel = "info"
	LoggerLevelError LoggerLevel = "error"
)

const maxExtraLength = 32 << 10 // 32 KB

type Logger struct {
	module string
}

func NewLogger(module string) *Logger {
	return &Logger{
		module: module,
	}
}

func (l *Logger) log(data string) {
	fmt.Println(data)
}

func (l *Logger) FormatMessage(
	c *kid.Ctx,
	message string,
	level LoggerLevel,
	extra map[string]interface{},
	err error,
) string {
	var extraStr string
	if extra != nil {
		extraBytes, _err := json.Marshal(extra)
		if _err != nil {
			logger.Error(c, "logger formatMessage error", nil, _err)
		} else if len(extraBytes) > maxExtraLength {
			extraStr = "too long to show"
		} else {
			extraStr = string(extraBytes)
		}
	}

	var requestId string
	if c != nil {
		requestId = c.Get("X-Request-ID").(string)
	}
	var errStr string
	if err != nil {
		errStr = err.Error()
	}

	units := []string{
		time.Now().Format(time.RFC3339),
		fmt.Sprintf("| %s", level),
		If(requestId != "", fmt.Sprintf("| %s", requestId), ""),
		If(l.module != "", "| ["+l.module+"]", ""),
		message,
		If(extraStr != "", "-Extra "+extraStr, ""),
		If(errStr != "", "-Error "+errStr, ""),
	}
	units = SliceFilter(units, func(item string) bool {
		return item != ""
	})
	return strings.ReplaceAll((strings.Join(units, " ")), "\n", "")
}

func (l *Logger) Info(c *kid.Ctx, message string, extra map[string]interface{}) {
	l.log(l.FormatMessage(c, message, LoggerLevelInfo, extra, nil))
}

func (l *Logger) Error(c *kid.Ctx, message string, extra map[string]interface{}, err error) {
	l.log(l.FormatMessage(c, message, LoggerLevelInfo, extra, err))
}

var logger = NewLogger("Logger")
