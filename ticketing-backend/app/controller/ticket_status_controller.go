package controller

import (
	"example-rest-api/app/service"
	"github.com/gin-gonic/gin"
)

type TicketStatusController interface {
	GetTicketStatus(c *gin.Context)
}

type TicketStatusControllerImpl struct {
	svc service.TicketStatusService
}

func (u TicketStatusControllerImpl) GetTicketStatus(c *gin.Context) {
	u.svc.GetTicketStatus(c)
}

func TicketStatusControllerInit(ticketStatusService service.TicketStatusService) *TicketStatusControllerImpl {
	return &TicketStatusControllerImpl{
		svc: ticketStatusService,
	}
}
