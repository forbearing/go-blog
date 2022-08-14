package routers

import (
	"strconv"

	apiv1 "github.com/forbearing/go-blog/api/v1"
	"github.com/forbearing/go-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(setting.AppMode)
	r := gin.Default()

	v1 := r.Group("api/v1")
	{

		// 用户模块的路由接口
		v1.POST("user/add", apiv1.AddUser)
		v1.GET("user/:id", apiv1.GetUserInfo)
		v1.GET("users", apiv1.GetUsers)
		v1.PUT("user/:id", apiv1.EditUser)
		v1.DELETE("user/:id", apiv1.DeleteUser)

		// 分类模块的路由接口
		v1.POST("category/add", apiv1.AddCategory)
		v1.GET("category", apiv1.GetCategory)
		v1.PUT("category/:id", apiv1.EditCategory)
		v1.DELETE("category/:id", apiv1.DeleteCategory)

		// 文章模块的路由接口
		v1.POST("article/add", apiv1.AddArticle)
		v1.GET("article", apiv1.GetArticle)
		v1.GET("article/info/:id", apiv1.GetArticleInfo)
		v1.PUT("article/:id", apiv1.EditArticle)
		v1.DELETE("article/:id", apiv1.DeleteArticle)
	}

	r.Run(":" + strconv.Itoa(setting.HttpPort))
}
