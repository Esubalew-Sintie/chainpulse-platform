package routes

import (
	"log/slog"
	"net/http"

	"github.com/chainpulse/backend/api-gateway/pkg/config"
	"github.com/gorilla/mux"
)

type Routes struct {
	Route  *mux.Router
	Logger slog.Logger
	Config *config.GlobalConfig
	HTTP   *http.Server
}

func NewServer(r *mux.Router, config *config.GlobalConfig, logger slog.Logger) *Routes {
	s := &Routes{
		Route:  r,
		Logger: logger,
		Config: config,
	}
	return s
}
