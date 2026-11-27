package httpapi

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"yumee/booking/internal/booking"

	"github.com/gin-gonic/gin"
)

type HouseRequest struct {
	Type     string `json:"type" binding:"required"`
	Capacity int    `json:"capacity" binding:"required,min=1"`
}

type HouseService interface {
	GetHouses(ctx context.Context) ([]booking.House, error)
	AddHouse(ctx context.Context, houseTypeRaw string, capacity int) (*booking.House, error)
	RemoveHouse(ctx context.Context, houseID int) error
}

// maybe change to private later
type HouseHandler struct {
	houseService HouseService
}

func NewHouseHandler(service HouseService) HouseHandler {
	return HouseHandler{houseService: service}
}

func (h *HouseHandler) GetHouses(c *gin.Context) {
	houses, err := h.houseService.GetHouses(c.Request.Context())
	if err != nil {
		log.Default().Println("Error trying to fetch houses")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get houses",
		})
		return
	}

	c.JSON(http.StatusOK, houses)
}

func (h *HouseHandler) AddHouse(c *gin.Context) {
	var req HouseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error trying to parse json body creating a new house")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newHouse, err := h.houseService.AddHouse(c.Request.Context(), req.Type, req.Capacity)
	if err != nil {
		if errors.Is(err, booking.InvalidHouseData) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Provided data is not valid"})
		} else {
			log.Printf("create house: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get house"})
		}
		return
	}

	c.JSON(http.StatusCreated, newHouse)
}

func (h *HouseHandler) RemoveHouse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be > 0",
		})
		return
	}

	err = h.houseService.RemoveHouse(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, booking.HouseDoesNotExist) {
			log.Println("remove house: house not found")
			c.JSON(http.StatusNotFound, gin.H{
				"error": "House not found",
			})
		} else {
			log.Printf("remove house: %w", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not remove house",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
