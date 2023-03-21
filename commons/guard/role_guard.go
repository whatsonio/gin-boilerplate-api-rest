package guard

import (
	"app/commons/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RolesCheck(r string) gin.HandlerFunc {

	return func(c *gin.Context) {

		user := c.Keys["user"].(*Claims)

		userRoles := strings.Split(user.Roles, ";")

		roles := strings.Split(strings.ReplaceAll(r, " ", ""), ",")

		isPermit := false

		for _, s := range roles {

			if lib.InArray(s, userRoles) > -1 {
				isPermit = true
			}

		}

		if !isPermit {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
				Message: "not authorized",
			})
			panic(nil)
		}

		c.Next()
	}

}
