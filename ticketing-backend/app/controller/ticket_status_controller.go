package controller

import (
	"github.com/gin-gonic/gin"
	"ticketing-api/app/service"
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
