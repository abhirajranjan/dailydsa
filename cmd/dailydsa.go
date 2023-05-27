package main

import (
	"log"

	"github.com/abhirajranjan/dailydsa/internal/auth"
	"github.com/abhirajranjan/dailydsa/internal/config"
	"github.com/abhirajranjan/dailydsa/internal/database"
	"github.com/abhirajranjan/dailydsa/internal/server"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	db := database.CreateDatabaseBridge(cfg.Database)
	auth.InitAuth(db, cfg.Auth)
	a := server.Serve(db)
	a.Run("localhost:1212")
}
