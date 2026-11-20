// Package contains a router to create a HTTP server
package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type house struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

// Registers a route for checking info about houses
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/houses", func(c *gin.Context) {
		c.JSON(http.StatusOK, []house{house{ID: 1, Name: "Big house", Capacity: 8},
			house{ID: 2, Name: "Small house", Capacity: 4}})
	})

	return r
}
