package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/go-backend/common/rest"
	"github.com/GabrielAchumba/go-backend/user-module/dtos"
	"github.com/GabrielAchumba/go-backend/user-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type UserController interface {
	Login(ctx *gin.Context) *rest.Response
	CreateUser(ctx *gin.Context) *rest.Response
	GetUsers(ctx *gin.Context) *rest.Response
	GetUser(ctx *gin.Context) *rest.Response
	UpdateUser(ctx *gin.Context) *rest.Response
	DeleteUser(ctx *gin.Context) *rest.Response
	DeleteUser2(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	userService services.UserService
}

func New(userService services.UserService) UserController {
	return &controllerImpl{
		userService: userService,
	}
}

func (ctrl *controllerImpl) Login(ctx *gin.Context) *rest.Response {
	var model dtos.LoginDTO
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.LoginUser(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) CreateUser(ctx *gin.Context) *rest.Response {
	var model dtos.PersonalProfieDTO
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.CreateUser(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetUsers(ctx *gin.Context) *rest.Response {

	if m, er := ctrl.userService.GetUsers(); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetUser(ctx *gin.Context) *rest.Response {

	lastname := ctx.Param("lastname")

	if m, er := ctrl.userService.GetUser(lastname); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) UpdateUser(ctx *gin.Context) *rest.Response {
	var model dtos.PersonalProfieDTO

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.UpdateUser(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteUser(ctx *gin.Context) *rest.Response {

	lastname := ctx.Param("lastname")

	if m, er := ctrl.userService.DeleteUser(lastname); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteUser2(ctx *gin.Context) *rest.Response {

	lastname := ctx.Param("lastname")

	if m, er := ctrl.userService.DeletUser2(lastname); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
