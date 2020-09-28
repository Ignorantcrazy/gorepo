package util

import (
	"main/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//GetPage return pagination data count
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
