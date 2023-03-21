package guard

import (
	"app/config"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type Claims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
	jwt.RegisteredClaims
}

func ExtractBearerToken(headerToken string) (string, error) {

	if headerToken == "" {
		return "", errors.New("no authorization header found")
	}

	jwtToken := strings.Split(headerToken, " ")

	if len(jwtToken) != 2 {
		return "", errors.New("no authorization header mal formated")
	}

	return jwtToken[1], nil
}

func ParseToken(jwtToken string) (*Claims, error) {

	claims := &Claims{}

	config := config.GetConfig()

	f, err := ioutil.ReadFile(config.Auth.PublicPemPath)

	if err != nil {
		return claims, err
	}

	_, err = jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {

		return jwt.ParseRSAPublicKeyFromPEM([]byte(f))

	})

	if err != nil {
		err = errors.New("token invalid")
		return claims, err
	}

	return claims, err
}

func JwtCheck() gin.HandlerFunc {

	return func(c *gin.Context) {

		jwt, err := ExtractBearerToken(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
				Message: err.Error(),
			})
			panic(nil)
		}

		parsedJwt, err := ParseToken(jwt)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
				Message: err.Error(),
			})
			panic(nil)
		}

		c.Set("user", parsedJwt)

		c.Next()

	}

}
