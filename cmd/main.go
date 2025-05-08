package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"scoreboard/config"
	db "scoreboard/db/queries"
	"scoreboard/internal/scoreboard"
)

func main() {
	config.InitConfig()
	conn, err := pgx.Connect(context.Background(), "postgres://"+config.DB_USER+":"+config.DB_PASSWORD+"@"+config.DB_HOST+":"+config.DB_PORT+"/"+config.DB_NAME)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	scoreboardService := scoreboard.NewScoreboardService(db.New(conn))
	scoreboardHandler := scoreboard.NewScoreboardHandler(scoreboardService)

	// 設定路由
	http.HandleFunc("/", scoreboardHandler.ListHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Unable to start server: %v\n", err)
	}
}
