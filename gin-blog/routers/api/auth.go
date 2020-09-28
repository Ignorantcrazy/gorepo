package api

import (
	"main/models"
	"main/pkg/e"
	"main/pkg/logging"
	"main/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	pwd := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: pwd}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if !ok {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  valid.Errors,
			"data": make(map[string]string),
		})
		return
	}

	if models.CheckAuth(username, pwd) {
		token, err := util.GenerateToken(username, pwd)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
