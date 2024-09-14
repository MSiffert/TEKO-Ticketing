package controller

import (
	"github.com/gin-gonic/gin"
	"ticketing-api/app/service"
)

type TicketPriorityController interface {
	GetTicketPriorities(c *gin.Context)
}

type TicketPriorityControllerImpl struct {
	svc service.TicketPriorityService
}

func (u TicketPriorityControllerImpl) GetTicketPriorities(c *gin.Context) {
	u.svc.GetTicketPriorities(c)
}

func TicketPriorityControllerInit(ticketPriorityService service.TicketPriorityService) *TicketPriorityControllerImpl {
	return &TicketPriorityControllerImpl{
		svc: ticketPriorityService,
	}
}
