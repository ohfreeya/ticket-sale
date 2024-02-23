package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type JWTClaims struct {
	Username string `json:"username"`
	Uid      int    `json:"uid"`
	jwt.StandardClaims
}

// Generate a JWT token by user name and id.
func GenerateToken(username string, uid int) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &JWTClaims{
		Username: username,
		Uid:      uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// Parse the JWT token and return the claims.
func VerifyToken(tokenStr string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token is expired")
		return nil, err
	}
	return 
}
