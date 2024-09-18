package ipusher

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key") // Replace with a secure key

// Generate a JWT token for a user
func generateJWT(userID int) (string, error) {
	// TODO: Pusher Authentication
	// Create the APP_KEY and APP_SECRET for the TT Backend with a long expiration date
	// The backend will call iPusher to generate a new token for the Logged-In User using the UserID Or Any Subject. The expiration will be short time.
	// - - iPusher will generate a JWT token using a JWT_SECRET_KEY
	// - - So, without the JWT_SECRET_KEY, no one can generate the JWT token for the front-end
	// - - Only registered backends will be able to request a new token from the iPusher server using the APP_KEY and APP_SECRET
	// The backend will then share the iPusher token with the front-end during login
	// The front-end will use the token for the Socket Connection
	// Finally, the iPusher will validate the JWT token using the iPuser's JWT_SECRET_KEY.

	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", userID),             // Store userID in the token
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expiration time
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
