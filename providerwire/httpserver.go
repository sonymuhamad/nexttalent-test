package providerwire

import (
	"net/http"
	"time"

	"github.com/sonymuhamad/nexttalent-test/config"
	"github.com/sonymuhamad/nexttalent-test/httpserver"
)

func provideHttpServer(cfg config.EnvConfig, h *httpserver.Handler) *http.Server {
	r := httpserver.InitRouter(cfg, h)

	return &http.Server{
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
