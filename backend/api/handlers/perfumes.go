package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamalfox85/perfume-app/backend/data"
)

type PerfumeFinder interface {
	GetAllPerfumes(context.Context, map[string]string) ([]data.Perfume, error)
}

func GetPerfumes(perfumes PerfumeFinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		filters := map[string]string{
			"q":      			c.Query("q"),
			"year":         	c.Query("year"),
			"concentration": 	c.Query("concentration"),
			"gender":       	c.Query("gender"),
			"category":     	c.Query("category"),
			"longevity":    	c.Query("longevity"),
		}

		perfumes, err := perfumes.GetAllPerfumes(ctx, filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"perfumes": perfumes})
	}
}