package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileFinder interface {
	CheckProfileExists(context.Context, string) (bool, error)
}

func GetProfile(profiles ProfileFinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for getting a profile by ID would go here
	}
}

func CheckProfileExists(profiles ProfileFinder) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		email := c.Param("email")

		exists, err := profiles.CheckProfileExists(ctx, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Exists:", exists)
		c.JSON(http.StatusOK, gin.H{"exists": exists})
	}
}