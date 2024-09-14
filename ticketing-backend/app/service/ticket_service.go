package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"ticketing-api/app/constant"
	"ticketing-api/app/domain/dao"
	"ticketing-api/app/pkg"
	"ticketing-api/app/repository"
)

type TicketService interface {
	GetTickets(c *gin.Context)
	GetTicketDetails(c *gin.Context)
	CreateTicket(c *gin.Context)
	UpdateTicket(c *gin.Context)
}

type TicketServiceImpl struct {
	ticketRepository repository.TicketRepository
}

func (t TicketServiceImpl) GetTickets(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")

	data, err := t.ticketRepository.GetAllTickets()
	if err != nil {
		log.Error("Happened Error when find all ticket data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TicketServiceImpl) GetTicketDetails(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")
	ticketID, _ := strconv.Atoi(c.Param("ticketID"))

	data, err := t.ticketRepository.GetTicketById(ticketID)
	if err != nil {
		log.Error("Happened Error when find all ticket data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TicketServiceImpl) CreateTicket(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data ticket")
	var request dao.Ticket
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := t.ticketRepository.CreateTicket(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TicketServiceImpl) UpdateTicket(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update ticket data by id")
	ticketID, _ := strconv.Atoi(c.Param("ticketID"))

	var request dao.Ticket
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := t.ticketRepository.GetTicketById(ticketID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Status = request.Status
	data.Priority = request.Priority
	data.SupporterUserID = request.SupporterUserID
	t.ticketRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func TicketServiceInit(ticketRepository repository.TicketRepository) *TicketServiceImpl {
	return &TicketServiceImpl{
		ticketRepository: ticketRepository,
	}
}
