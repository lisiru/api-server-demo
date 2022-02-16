package api_server_demo

import (
	"api-server-demo/cache/redis"
	"api-server-demo/controller/v1/user"
	"api-server-demo/pkg/response"
	"api-server-demo/store/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}
func installController(g *gin.Engine) *gin.Engine {
	g.GET("test", func(context *gin.Context) {
		res := make(map[string]string)
		res["user"] = "lisr"
		response.WriteResponse(context, nil, res)

	})
	g.GET("testTime", func(context *gin.Context) {
		time.Sleep(10 * time.Second)
		context.String(http.StatusOK, "success")
	})

	// 获取mysql的
	storeInstance, _ := mysql.GetMySQLFactoryOr(nil)
	cacheInstance, _ := redis.NewRedisFactoryOr(nil)
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeInstance, cacheInstance)
			userv1.GET(":name", userController.GetUser)
		}
	}
	return g
}
