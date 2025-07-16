package documents

import (
	"database/sql"
	"net/http"
	"upload-service/pkg/middlewares"
)

func DocumentRouter(db *sql.DB) *http.ServeMux {
	controller := DocumentController(db)
	router := http.NewServeMux()

	router.HandleFunc("GET /", middlewares.JsonMiddleware(controller.Index))
	router.HandleFunc("POST /", middlewares.JsonMiddleware(middlewares.AuthMiddleware(controller.Store)))
	router.HandleFunc("DELETE /{id}", middlewares.JsonMiddleware(middlewares.AuthMiddleware(controller.Delete)))

	return router
}
