package documents

import (
	"database/sql"
	"net/http"
	"upload-service/configs"
	"upload-service/pkg/middlewares"
)

func DocumentRouter(db *sql.DB, appConfig configs.AppConfig) *http.ServeMux {
	controller := DocumentController(db)
	router := http.NewServeMux()

	router.HandleFunc("GET /", middlewares.JsonMiddleware(controller.Index))
	router.HandleFunc("POST /", middlewares.JsonMiddleware(middlewares.AuthMiddleware(appConfig)(controller.Store)))
	router.HandleFunc("DELETE /{id}", middlewares.JsonMiddleware(middlewares.AuthMiddleware(appConfig)(controller.Delete)))

	return router
}
