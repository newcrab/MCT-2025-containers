package main

import (
	"log"

	"github.com/newcrab/MCT-2025-containers.git/internal/config"
	"github.com/newcrab/MCT-2025-containers.git/internal/controller"
	pg "github.com/newcrab/MCT-2025-containers.git/internal/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	pgStore := pg.NewPostgresDB(cfg)

	router := gin.Default()
	handler := controller.NewHandler(cfg, pgStore)
	handler.SetupRoutes(router)

	log.Printf("Server started at port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
