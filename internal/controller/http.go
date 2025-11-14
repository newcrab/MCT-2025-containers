package controller

import (
	"fmt"
	"log"

	"github.com/newcrab/MCT-2025-containers.git/internal/config"
	pg "github.com/newcrab/MCT-2025-containers.git/internal/postgres"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	config *config.Config
	pg     *pg.PostgresDB
}

func NewHandler(cfg *config.Config, pg *pg.PostgresDB) *Handler {
	return &Handler{
		config: cfg,
		pg:     pg,
	}
}

func (h *Handler) PingHandler(c *gin.Context) {
	err := h.pg.SaveVisit(c.ClientIP())
	if err != nil {
		log.Printf("error while saving visit: %v", err)
		return
	}

	c.String(200, "pong")
}

func (h *Handler) VisitsHandler(c *gin.Context) {
	count, err := h.pg.GetVisitsCount(c.ClientIP())
	if err != nil {
		log.Printf("error getting visits num: %v", err)
		return
	}

	c.String(200, fmt.Sprintf("%d", count))
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
	router.GET("/ping", h.PingHandler)
	router.GET("/visits", h.VisitsHandler)
}
