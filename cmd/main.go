package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"scoreboard/config"
	"scoreboard/internal/scoreboard"
	"scoreboard/internal/scoreboard/db"
)

func main() {
	config.InitConfig()
	conn, err := pgx.Connect(context.Background(), "postgres://"+config.DB_USER+":"+config.DB_PASSWORD+"@"+config.DB_HOST+":"+config.DB_PORT+"/"+config.DB_NAME)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	scoreboardService := scoreboard.NewScoreboardService(db.New(conn))
	scoreboardHandler := scoreboard.NewScoreboardHandler(scoreboardService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			scoreboardHandler.ListHandler(w)
		case http.MethodPost:
			scoreboardHandler.CreateHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Unable to start server: %v\n", err)
	}
}
