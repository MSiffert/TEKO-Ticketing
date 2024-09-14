package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"ticketing-api/app/service"
)

type TicketController interface {
	GetAllTickets(c *gin.Context)
	GetTicketDetails(c *gin.Context)
	CreateTicket(c *gin.Context)
	UpdateTicket(c *gin.Context)
}

type TicketControllerImpl struct {
	svc service.TicketService
}

func (u TicketControllerImpl) GetAllTickets(c *gin.Context) {
	log.Println("Executing get all tickets.")
	u.svc.GetTickets(c)
}

func (u TicketControllerImpl) GetTicketDetails(c *gin.Context) {
	log.Println("Executing get ticket details.")
	u.svc.GetTicketDetails(c)
}

func (u TicketControllerImpl) CreateTicket(c *gin.Context) {
	log.Println("Executing create ticket.")
	u.svc.CreateTicket(c)
}

func (u TicketControllerImpl) UpdateTicket(c *gin.Context) {
	log.Println("Executing update ticket status.")
	u.svc.UpdateTicket(c)
}

func TicketControllerInit(ticketService service.TicketService) *TicketControllerImpl {
	return &TicketControllerImpl{
		svc: ticketService,
	}
}
