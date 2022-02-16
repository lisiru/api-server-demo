package user

import (
	"api-server-demo/cache"
	"api-server-demo/pkg/logger"
	"api-server-demo/pkg/response"
	"api-server-demo/service"
	"api-server-demo/store"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.Service
}

func NewUserController(store store.Factory,cache cache.Factory) *UserController  {
	return &UserController{
		service: service.NewService(store,cache),
	}
}

func (us *UserController) GetUser(c *gin.Context)  {
	logger.L(c).Info("get user function called.")
	type userTest struct {
		Name string `json:"name"`
	}
	response.WriteResponse(c,nil,userTest{
		Name: "lisiru",
	})
	return

	user,err:=us.service.Users().Get(c,c.Param("name"))
	if err != nil {
		response.WriteResponse(c,err,nil)
		return
	}
	response.WriteResponse(c,err,user)
}