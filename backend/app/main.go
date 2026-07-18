package main

import (
	"log/slog"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "gas-thief/app/docs"
)

// @title Gas Thief API
// @version 1.0
// @description API for gas price tracking
// @host localhost:8080
// @BasePath /
func main() {
	settings := LoadSettings()
	db := connectToDB(settings)
	go run(db, settings)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", handleHealth(db))
	mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	slog.Info("starting web server", "port", settings.WebPort)
	if err := http.ListenAndServe(":"+settings.WebPort, middleware(mux)); err != nil {
		slog.Error("web server failed to start", "port", settings.WebPort, "error", err)
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				slog.Warn("web middleware recovery caught smth", "error", err)
				http.Error(w, "there was an internal server error", http.StatusInternalServerError)
			}
		}()

		slog.Info("incoming request", "method", r.Method, "path", r.URL.Path, "remote", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
