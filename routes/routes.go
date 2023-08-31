package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/songyanping/gin-simple/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default() //生成了一个WSGI应用程序实例
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSONP(200, "success")
		})

		//用户操作
		v1.GET("user/register", api.UserRegister)
	}

	return r
}
