package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamalfox85/perfume-app/backend/api/handlers"
	"github.com/jamalfox85/perfume-app/backend/api/middleware"
)

func (app *Application) Routes() http.Handler {

    router := gin.Default()

    // Global Middlewares
    router.Use(CORSMiddleware())

    public := router.Group("/")
    {
        public.GET("/health", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Service is running!"})
        })

		// Profile Routes
        public.GET("/profile/check-email/:email", handlers.CheckProfileExists(app.Profiles))

		// Perfume Routes
		public.GET("/perfumes", handlers.GetPerfumes(app.Perfumes))
    }
	
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
		// Profile Routes
        protected.GET("/profile/:id", handlers.GetProfile(app.Profiles))
    }

    return router
}
