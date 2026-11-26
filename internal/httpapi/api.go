// Package contains a router to create a HTTP server
package httpapi

import (
	"github.com/gin-gonic/gin"
)

type house struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

// Registers a route for checking info about houses
func NewRouter(service HouseService) *gin.Engine {
	r := gin.Default()
	handler := NewHouseHandler(service)
	r.GET("/houses", handler.GetHouses)
	r.POST("/houses", handler.AddHouse)
	r.DELETE("/houses/:id", handler.RemoveHouse)

	return r
}
