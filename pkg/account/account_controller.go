package account

import (
	"app/commons/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

var accountRepo = new(AccountRepository)

func (ctl AccountController) Create(c *gin.Context) {

	var data AccountCreateInput

	helpers.ShouldBindJSON(c, &data)

	account, err := accountRepo.Create(&data)

	helpers.ResponseJSON(c, account, err)

}

func (ctl AccountController) Find(c *gin.Context) {

	result, err := accountRepo.Find(c.Request.URL.Query())

	helpers.ResponseJSON(c, result, err)

}

func (ctl AccountController) FindOne(c *gin.Context) {

	i, _ := strconv.Atoi(c.Param("id"))

	account, err := accountRepo.FindOneById(uint(i))

	helpers.ResponseJSON(c, account, err)

}

func (ctl AccountController) Update(c *gin.Context) {

	var data AccountUpdateInput

	i, _ := strconv.Atoi(c.Param("id"))

	helpers.ShouldBindJSON(c, &data)

	account, err := accountRepo.Update(uint(i), &data)

	helpers.ResponseJSON(c, account, err)

}

func (ctl AccountController) Delete(c *gin.Context) {

	i, _ := strconv.Atoi(c.Param("id"))

	account, err := accountRepo.Delete(uint(i))

	helpers.ResponseJSON(c, account, err)

}
