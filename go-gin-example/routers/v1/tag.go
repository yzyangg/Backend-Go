package v1

import (
	"Backend-Go/go-gin-example/models"
	"Backend-Go/go-gin-example/pkg/err"
	"Backend-Go/go-gin-example/pkg/setting"
	"Backend-Go/go-gin-example/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	//c.Query可用于获取?name=test&state=1这类 URL 参数，而c.DefaultQuery则支持设置一个默认值
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name

	}
	var state = -1
	value := c.Query("state")
	if value != "" {
		state = com.StrTo(value).MustInt()
		maps["state"] = state
	}
	code := err.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)

	util.Success(c, http.StatusOK, code, data)
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	createdBy := c.Query("created_by")
	// DefaultQuery支持设置一个默认值
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	//校验
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := err.INVALID_PARAMS
	var success bool
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = err.SUCCESS
			success = models.AddTag(name, state, createdBy)
		} else {
			code = err.ERROR_EXIST_TAG
		}
	}
	util.Success(c, http.StatusOK, code, success)
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state = -1
	arg := c.Query("state")
	if arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := err.INVALID_PARAMS
	if !valid.HasErrors() {
		code = err.SUCCESS

		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name

			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = err.ERROR_NOT_EXIST_TAG
		}
		util.Success(c, http.StatusOK, code, make(map[string]interface{}))
	}
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := err.INVALID_PARAMS
	if !valid.HasErrors() {
		exist := models.ExistTagById(id)
		{
			if exist {
				code = err.SUCCESS
				models.DeleteTag(id)
			} else {
				code = err.ERROR_NOT_EXIST_TAG
			}
		}
	}
	util.Success(c, http.StatusOK, code, make(map[string]interface{}))
}
