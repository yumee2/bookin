package httpapi

import (
	"context"
	"net/http"
	"yumee/booking/internal/booking"

	"github.com/gin-gonic/gin"
)

type HouseService interface {
	GetHouses(ctx context.Context) ([]booking.House, error)
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get houses",
		})
		return
	}

	c.JSON(http.StatusOK, houses)
}
