package forcesslheroku

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var testCases = []struct {
	goEnv     string
	proto     string
	expectLoc string
}{
	{goEnv: "production", proto: "http",
		expectLoc: "https://example.com/test"},
	{goEnv: "production", proto: "https"},
	{goEnv: "test", proto: "http"},
	{goEnv: "test", proto: "https"},
}

func TestForceSsl(t *testing.T) {
	noopHandler := func(w http.ResponseWriter, r *http.Request) {}
	forceSsl := ForceSsl(http.HandlerFunc(noopHandler))

	for _, tt := range testCases {
		getenv = func(key string) string {
			switch key {
			case goEnviron:
				return tt.goEnv
			default:
				return ""
			}
		}

		t.Run(tt.goEnv+"_"+tt.proto, func(t *testing.T) {
			req := httptest.NewRequest("", "/test", nil)
			req.Header.Add(xForwardedProtoHeader, tt.proto)

			res := httptest.NewRecorder()
			forceSsl.ServeHTTP(res, req)

			if location := res.Header().Get("Location"); location != tt.expectLoc {
				t.Errorf("expected Location header '%s', got '%s'",
					tt.expectLoc, location)
			}
		})
	}
}

