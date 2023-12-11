package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Set token expiration to one month
const TokenExpireDuration = time.Hour * 24 * 30
const UserTokenExpirationSec = 900

// Constrant variable for jwt, works as a key
var secret = []byte("Here is tiktok")

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

// IssueToken returns a jwt for given user
func IssueToken(username string) (string, error) {
	c := MyClaims{
		username,
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Token
		tokenString := c.GetHeader("Authorization")
		// 验证 Token
		claims, err := ParseToken(tokenString)
		if err != nil {
			// Token 验证失败
			// 而不是直接返回，设置一个标记表示未认证
			c.Set("isAuthenticated", false)
		} else {
			// Token 验证成功，将用户信息添加到上下文中
			c.Set("isAuthenticated", true)
			c.Set("username", claims.Username)
		}

		// 继续执行后续的处理函数
		c.Next()
	}
}

func jwtFunc(token *jwt.Token) (interface{}, error) {
	return secret, nil
}
