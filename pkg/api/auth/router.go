package auth

import (
	"database/sql"
	"net/http"
	"upload-service/configs"
	"upload-service/pkg/middlewares"
)

func AuthRouter(db *sql.DB, appConfig configs.AppConfig) *http.ServeMux {
	controller := AuthController(db, appConfig)
	router := http.NewServeMux()

	router.HandleFunc("POST /login", middlewares.JsonMiddleware(controller.Login))
	router.HandleFunc("POST /register", middlewares.JsonMiddleware(controller.Register))

	return router
}
