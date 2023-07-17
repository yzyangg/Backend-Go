package v1

import (
	"Backend-Go/go-gin-example/models"
	"Backend-Go/go-gin-example/pkg/err"
	"Backend-Go/go-gin-example/pkg/setting"
	"Backend-Go/go-gin-example/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
)

// Controller层

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := err.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		// 判断文章是否存在
		if models.ExistArticleById(id) {
			// 获取文章
			data = models.GetArticle(id)
			code = err.SUCCESS
		} else {
			// 返回错误信息
			for _, error := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", error.Key, error.Message)
			}
		}
	}
	util.Success(c, code, err.GetMsg(code), data)
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	params := make(map[string]interface{})
	data := make(map[string]interface{})

	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		params["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var TagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		TagId = com.StrTo(arg).MustInt()
		params["tag_id"] = TagId
		valid.Min(TagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	code := err.INVALID_PARAMS
	if !valid.HasErrors() {
		code = err.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, params)
		data["total"] = models.GetArticleTotal(params)
	}
	util.Success(c, code, err.GetMsg(code), data)
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := err.INVALID_PARAMS

	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = err.SUCCESS

		} else {
			code = err.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	util.Success(c, code, err.GetMsg(code), make(map[string]interface{}))
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	// 验证参数
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := err.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
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
				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = err.SUCCESS
			} else {
				code = err.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = err.ERROR_NOT_EXIST_ARTICLE
		}

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	util.Success(c, code, err.GetMsg(code), make(map[string]interface{}))
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	log.Printf("id: %d", id)

	valid := validation.Validation{}
	valid.Min(id, 0, "id").Message("ID必须大于0")

	code := err.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			models.DeleteArticle(id)
			code = err.SUCCESS
		} else {
			code = err.ERROR_NOT_EXIST_ARTICLE
		}
	}
	util.Success(c, code, err.GetMsg(code), make(map[string]interface{}))
}
