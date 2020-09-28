package v1

import (
	"main/models"
	"main/pkg/e"
	"main/pkg/setting"
	"main/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  valid.Errors,
			"data": make(map[string]string),
		})
		return
	}

	code = e.SUCCESS
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许1或0")
	}

	tagid := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagid = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagid
		valid.Min(tagid, 1, "tag_id").Message("标签ID最小大于1")
	}

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  valid.Errors,
			"data": make(map[string]string),
		})
		return
	}

	code = e.SUCCESS
	data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	tagid := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdby := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagid, 1, "tag_id").Message("标签ID最小大于1")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdby, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许1或0")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  valid.Errors,
			"data": make(map[string]string),
		})
		return
	}

	if models.ExistTagByID(tagid) {
		code = e.SUCCESS
		data := make(map[string]interface{})
		data["tag_id"] = tagid
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdby
		data["state"] = state

		models.AddArticle(data)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagid := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdby := c.Query("created_by")

	valid := validation.Validation{}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许1或0")
	}

	valid.Min(id, 1, "id").Message("ID最小大于1")
	valid.Min(tagid, 1, "tag_id").Message("标签ID最小大于1")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdby, "created_by").Message("创建人不能为空")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  valid.Errors,
			"data": make(map[string]string),
		})
		return
	}

	if models.ExistArticleByID(id) {
		if models.ExistTagByID(tagid) {
			code = e.SUCCESS

			data := make(map[string]interface{})
			if tagid > 0 {
				data["tag_id"] = tagid
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}
			if createdby != "" {
				data["created_by"] = createdby
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditArticle(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DelArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID最小大于1")
	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  valid.Errors,
			"data": make(map[string]string),
		})
		return
	}
	if models.ExistArticleByID(id) {
		code = e.SUCCESS
		models.DelArticle(id)
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
