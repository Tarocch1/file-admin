package middleware

import "net/http"

func Middleware(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return AuthMiddleware(fn)
}
