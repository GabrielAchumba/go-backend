package usermodule

import (
	"github.com/GabrielAchumba/go-backend/common/rest"
	"github.com/GabrielAchumba/go-backend/user-module/controllers"
	"github.com/GabrielAchumba/go-backend/user-module/services"
	"github.com/gin-gonic/gin"
)

type userModule struct {
	controller controllers.UserController
}

func InjectService(service services.UserService) *userModule {
	module := new(userModule)
	module.controller = controllers.New(service)
	return module
}

func (module *userModule) RegisterRoutes(rg *gin.RouterGroup) {
	moduleRoute := rg.Group("/user")
	serverHttp := rest.ServeHTTP

	moduleRoute.POST("/movielogin", serverHttp(module.controller.Login))
	moduleRoute.POST("/movieregister", serverHttp(module.controller.CreateUser))
	moduleRoute.GET("/getusers", serverHttp(module.controller.GetUsers))
	moduleRoute.GET("/getuser/:lastname", serverHttp(module.controller.GetUser))
	moduleRoute.POST("/update", serverHttp(module.controller.UpdateUser))
	moduleRoute.DELETE("/delete/:lastname", serverHttp(module.controller.DeleteUser))
	moduleRoute.DELETE("/delete2/:lastname", serverHttp(module.controller.DeleteUser2))
}
