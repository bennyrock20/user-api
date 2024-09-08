package cmd

import (
	"fmt"
	"taxi-service/middlewares"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Generate a JWT token for authentication",
	Long: `This command generates a JSON Web Token (JWT) that can be used for authentication purposes.
The token includes a user ID and a short expiration time (5 minutes) for security.
Example usage:
	taxi-service jwt`,
	Run: func(cmd *cobra.Command, args []string) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": "12345",
			"exp":     time.Now().Add(time.Minute * 5).Unix(), // Token expiration
		})

		tokenString, _ := token.SignedString(middlewares.JWTSecret)
		// Output the signed token
		fmt.Println("Generated JWT token:", tokenString)
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}
