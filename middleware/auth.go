package middleware

import (
	"net/http"
	"strings"

	"github.com/Tarocch1/file-admin/common"
)

func AuthMiddleware(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if common.FlagAuth == "" {
			fn(w, r)
			return
		}
		basicAuthPrefix := "Basic "
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			token := auth[len(basicAuthPrefix):]
			if token == common.FlagAuth {
				fn(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", "Basic")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
