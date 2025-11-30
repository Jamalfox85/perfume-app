package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamalfox85/perfume-app/backend/data"
)

type CabinetFinder interface {
	GetProfileCabinet(context.Context, string) (data.Cabinet, error)
}

func GetProfileCabinet(cabinets CabinetFinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userId := c.Param("userId")

		cabinet, err := cabinets.GetProfileCabinet(ctx, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cabinet": cabinet})
	}
}