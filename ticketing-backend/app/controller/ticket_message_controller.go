package controller

import (
	"example-rest-api/app/service"
	"github.com/gin-gonic/gin"
)


type TicketMessageController interface {
	CreateTicketMessage(c *gin.Context)
	DeleteTicketMessage(c *gin.Context)
}

type TicketMessageControllerImpl struct {
	svc service.TicketMsgService
}

func (u TicketMessageControllerImpl) CreateTicketMessage(c *gin.Context) {
	u.svc.CreateTicketMessage(c)
}

func (u TicketMessageControllerImpl) DeleteTicketMessage(c *gin.Context) {
	u.svc.DeleteTicketMessage(c)
}


func TicketMessageControllerInit(userService service.TicketMsgService) *TicketMessageControllerImpl {
	return &TicketMessageControllerImpl{
		svc: userService,
	}
}