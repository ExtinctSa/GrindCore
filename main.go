package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ExtinctSa/final_project/internal/database"
	handler "github.com/ExtinctSa/final_project/internal/handlers"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbUrl := os.Getenv("DB_URL")
	Secret_Key := os.Getenv("SK")
	platform := os.Getenv("PLATFORM")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Database error: ", err)
	}
	dbQueries := database.New(db)

	apiCfg := &handler.ApiConfig{
		Platform:  platform,
		Sk:        Secret_Key,
		DBQueries: dbQueries,
	}

	mux := http.NewServeMux()

	http.StripPrefix("/app", http.FileServer(http.Dir(".")))

	//User Handlers
	mux.HandleFunc("POST /api/users", apiCfg.CreateUserHandler)
	mux.HandleFunc("POST /api/login", apiCfg.UserLogin)

	server := &http.Server{
		Addr:    ":9999",
		Handler: mux,
	}

	server.ListenAndServe()
}
