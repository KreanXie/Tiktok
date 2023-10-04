package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Set token expiration to one month
const TokenExpireDuration = time.Hour * 24 * 30
const UserTokenExpirationSec = 900

// Constrant variable for jwt, works as a key
var secret = []byte("Here is tiktok")

type MyClaims struct {
	Account string
	jwt.StandardClaims
}

// IssueToken returns a jwt for given user
func IssueToken(account string) (string, error) {
	c := MyClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Krean",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

// Parse given token, return claims if caller needs
func ParseToken(tokenString string) (*MyClaims, error) {
	c := new(MyClaims)
	if token, err := jwt.ParseWithClaims(tokenString, c, jwtFunc); err != nil {
		return nil, err
	} else {
		if _, ok := token.Claims.(*MyClaims); !ok {
			return nil, errors.New("unable to assert token'claims as *MyClaims")
		} else if !token.Valid {
			return nil, errors.New("invalid token")
		} else {
			return c, nil
		}
	}
}

func jwtFunc(token *jwt.Token) (interface{}, error) {
	return secret, nil
}
