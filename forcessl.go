package forcesslheroku

import (
	"net/http"
	"os"
)

var getenv = os.Getenv

const (
	xForwardedProtoHeader = "x-forwarded-proto"
	goEnviron             = "GO_ENV"
)

func ForceSsl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getenv(goEnviron) == "production" {
			if r.Header.Get(xForwardedProtoHeader) != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
