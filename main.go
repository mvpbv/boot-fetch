package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mvpbv/boot-fetch/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	Default := Boot{
		Art:  wizardFrog,
		Init: "Boot.fetch initializing",
	}
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
		fmt.Println("No PORT environment variable detected, defaulting to " + portString)
	}
	dbUrl := os.Getenv("CONN")
	if dbUrl == "" {
		log.Fatal("No CONN environment variable detected")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go apiCfg.Boot_Fetch(&Default)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/Users", apiCfg.createUser)
	mux.HandleFunc("PUT /api/v1/Users/{id}", apiCfg.updateUser)
	mux.HandleFunc("GET /api/v1/Users/{id}", apiCfg.getUserStats)
	mux.HandleFunc("POST /api/v1/Messages", apiCfg.addMessage)
	mux.HandleFunc("GET /api/v1/ArchWizard", apiCfg.getArchWizard)
	mux.HandleFunc("POST /api/v1/Fix", apiCfg.dbFixxer)
	mux.HandleFunc("POST /api/v1/Duels", apiCfg.addDuel)
	fs := http.FileServer(http.Dir("./static/"))

	mux.Handle("/logs/", fs)
	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: mux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
