package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/2Rahul2/trafficControl/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB     *database.Queries
	DBConn *sql.DB
}

func main() {
	var portString string = "8001"
	fmt.Println("World")
	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Could not open connection with database")
	}

	connection.SetMaxOpenConns(200)
	connection.SetMaxIdleConns(200)
	connection.SetConnMaxLifetime(5 * time.Minute)
	err = connection.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	db := database.New(connection)
	apiCfg := apiConfig{
		DB:     db,
		DBConn: connection,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	V1Router := chi.NewRouter()
	V1Router.Get("/", apiCfg.testNetwork)
	V1Router.Post("/create", apiCfg.handlerUser)
	V1Router.Post("/create-concert", apiCfg.handlerCreateConcert)
	V1Router.Post("/book-ticket", apiCfg.middlewareHandler(apiCfg.handlerBookTickets))
	V1Router.Post("/book-ticket-nogo", apiCfg.middlewareHandler(apiCfg.handlerBookTicketsNoGo))

	V1Router.Post("/book-ticket-uni", apiCfg.middlewareHandler(apiCfg.handlerBookTicketsUni))
	V1Router.Post("/book-ticket-shards", apiCfg.middlewareHandler(apiCfg.handlerBookTicketsWithShardin))

	V1Router.Post("/create-concert-shards", apiCfg.handlerCreateShardsConcert)

	router.Mount("/v1", V1Router)
	serve := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("server starting on port %v", portString)

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
