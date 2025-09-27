package handlers
import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

func SetupServer(s *Handlers) error {
	corsMiddleware := handlers.CORS(
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://localhost:3000",
		}),
	)

	corsRouter := corsMiddleware(s.handlers.Route)
	s.handlers.Route.Use(corsMiddleware)
	s.SetupRoutes()
	s.handlers.HTTP = &http.Server{
		Addr:           s.Config.SERVER.HOST + ":" + s.Config.SERVER.PORT,
		Handler:        handlers.LoggingHandler(os.Stdout, corsRouter),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.handlers.Logger.Info("server started on port ", "", s.Config.SERVER.PORT)
	if err := s.handlers.HTTP.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %v", s.handlers.HTTP.Addr, err)
	}
	return nil
}
