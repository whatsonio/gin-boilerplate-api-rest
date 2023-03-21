package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FORMATER LES MESSAGES DE RETOUR
func ShouldBindJSON(c *gin.Context, obj any) {

	validationErr := c.ShouldBindJSON(obj)

	//TODO : Format la validation des champs
	if validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Json data not valid", "data": validationErr})
		panic(nil)
	}

}

func ResponseJSON(c *gin.Context, obj any, err error) {

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err})
		panic(nil)
	}

	c.JSON(http.StatusOK, obj)
	c.Abort()

}
