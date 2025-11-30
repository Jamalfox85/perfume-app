package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type apikeyTransport struct {
    apiKey string
    rt     http.RoundTripper
}

var JWKS *keyfunc.JWKS

func InitJWKS() {
    jwksURL := "https://mkqhgroqfifpjkhfnoyt.supabase.co/auth/v1/.well-known/jwks.json"

	client := &http.Client{
        Timeout: 10 * time.Second,
        Transport: &apikeyTransport{
            apiKey: os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
            rt:     http.DefaultTransport,
        },
    }

    var err error
    JWKS, err = keyfunc.Get(jwksURL, keyfunc.Options{
        RefreshInterval:  time.Hour,
		RefreshRateLimit:   time.Minute * 5,      
    	RefreshTimeout:     time.Second * 10,
        RefreshErrorHandler: func(err error) {
            log.Println("JWKS refresh error:", err)
        },
		RefreshUnknownKID:  true,
		Client: client,
    })

    if err != nil {
        log.Fatalf("Failed to load JWKS: %v", err)
    }

}

func (t *apikeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    req.Header.Add("apikey", t.apiKey)
    return t.rt.RoundTrip(req)
}


func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "missing Authorization header",
            })
            return
        }

        if !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "invalid Authorization header",
            })
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        // Parse & validate JWT
        token, err := jwt.Parse(tokenStr, JWKS.Keyfunc)
        if err != nil || !token.Valid {
			fmt.Println("Token parsing error:", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "invalid token",
            })
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "invalid claims",
            })
            return
        }

        // Attach user ID to the context for handlers
        if sub, ok := claims["sub"].(string); ok {
            c.Set("userId", sub)
        }

        c.Next()
    }
}
