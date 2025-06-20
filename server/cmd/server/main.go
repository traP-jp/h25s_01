package main

import (
	"log"

	"backend/cmd/server/server"
	"backend/pkg/config"
	"backend/pkg/database"
)

func main() {
	db, err := database.Setup(config.MySQL())
	if err != nil {
		log.Fatal("Failed to setup database:", err)
	}
	defer db.Close()

	s := server.Inject(db)

	log.Printf("Server starting on %s", config.AppAddr())
	if err := s.Router.Start(config.AppAddr()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
