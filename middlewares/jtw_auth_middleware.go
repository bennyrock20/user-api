package middlewares

import (
	"fmt"
	"net/http"
	"taxi-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = utils.GetEnv("JWT_SECRET_KEY", "")

var JWTSecret = []byte(secretKey)

// Middleware to check JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Optionally pass the user ID or other claims to the context
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("userID", claims["user_id"])
		}

		c.Next()
	}
}

// func main() {
// 	r := gin.Default()

// 	// Protect your routes with JWT middleware
// 	r.GET("/protected", JWTAuthMiddleware(), func(c *gin.Context) {
// 		userID, _ := c.Get("userID")
// 		c.JSON(http.StatusOK, gin.H{"message": "Welcome", "userID": userID})
// 	})

// 	r.Run(":8080")
// }
