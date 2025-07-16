package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"upload-service/configs"
	"upload-service/pkg/api/auth"
	"upload-service/pkg/api/documents"
	"upload-service/pkg/database"
	"upload-service/pkg/middlewares"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Warning: loading .env failed: %v", err)
		}
	}

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	appConfig := configs.GetAppConfig()

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/documents/", http.StripPrefix("/documents", documents.DocumentRouter(db)))
	mainRouter.Handle("/auth/", http.StripPrefix("/auth", auth.AuthRouter(db, appConfig)))

	rootRouter := middlewares.CORSMiddleware(mainRouter)

	fmt.Printf("Server starting on port %s\n", appConfig.Port)
	log.Fatal(http.ListenAndServe(":"+appConfig.Port, rootRouter))
}
