package main

import (
	"log"

	"user-service/internal/config"
	"user-service/internal/database"
	"user-service/internal/router"
)

func main() {
	cfg := config.Load()

	db := database.New(cfg)

	// Автоматическая миграция — только для локальной разработки.
	// В продакшене используйте SQL-миграции (migrations/001_create_users.sql).
	// db.AutoMigrate(&user.User{})

	r := router.New(db, cfg)

	log.Printf("server listening on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}