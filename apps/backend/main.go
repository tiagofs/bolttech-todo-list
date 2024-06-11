package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tiagofs/bolttech-todo-list/api/config"
	"github.com/tiagofs/bolttech-todo-list/api/database"
	"github.com/tiagofs/bolttech-todo-list/api/routes"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	fmt.Print(cfg)

	pool, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	app := gin.Default()

	routes.SetupApiRoutes(app, pool)

	// Mote this to config/env file.
	host := "localhost"
	port := "8080"

	go func() {
		if err := app.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
			log.Panic(err)
		}
	}()
}
