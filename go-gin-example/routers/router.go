package routers

import (
	jwt "Backend-Go/go-gin-example/middleware"
	"Backend-Go/go-gin-example/pkg/setting"
	v1 "Backend-Go/go-gin-example/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 新的路由实例
	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger())

	// 使用中间件
	r.Use(gin.Recovery())

	// 设置gin框架的运行模式
	gin.SetMode(setting.RunMode)

	//api := r.Group("/api/v1")
	//{
	//	// 获取标签列表
	//	api.GET("/tags", v1.GetTags)
	//	// 新增标签
	//	api.POST("/tags", v1.AddTag)
	//	// 修改标签
	//	api.PUT("/tags/:id", v1.EditTag)
	//	// 删除标签
	//	api.DELETE("/tags/:id", v1.DeleteTag)
	//}

	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	r.GET("/auth", v1.GetAuth)

	return r
}
