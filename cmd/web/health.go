package main

import (
	"github.com/alexhiggins/go-api/internal/transformer"
	"github.com/gin-gonic/gin"
)

// ShowHealthHandler godoc
// @Summary Show the status of server.
// @Accept */*
// @Produce json
// @Success 200 {object} transformer.ShowHealthResponse
// @Router /health [get]
func (s *server) ShowHealthHandler(c *gin.Context) {
	s.statusOk(c, transformer.ShowHealth("OK"))
}
