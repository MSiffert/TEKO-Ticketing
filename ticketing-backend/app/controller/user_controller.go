package controller

import (
	"github.com/gin-gonic/gin"
	"ticketing-api/app/service"
)

type UserController interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u UserControllerImpl) GetUsers(c *gin.Context) {
	u.svc.GetUsers(c)
}

func (u UserControllerImpl) CreateUser(c *gin.Context) {
	u.svc.CreateUser(c)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
