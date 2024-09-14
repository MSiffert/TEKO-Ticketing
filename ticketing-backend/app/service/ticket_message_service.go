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

type TicketMsgService interface {
	CreateTicketMessage(c *gin.Context)
	DeleteTicketMessage(c *gin.Context)
}

type TicketMsgServiceImpl struct {
	ticketMessageRepository repository.TicketMessageRepository
}

func (t TicketMsgServiceImpl) CreateTicketMessage(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data ticket")
	var request dao.TicketMessage
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := t.ticketMessageRepository.CreateTicketMessage(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (t TicketMsgServiceImpl) DeleteTicketMessage(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	ticketMessageId, _ := strconv.Atoi(c.Param("ticketMessageID"))

	err := t.ticketMessageRepository.DeleteTicketMessage(ticketMessageId)
	if err != nil {
		log.Error("Happened Error when try delete data ticket message from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func TicketMessageServiceInit(ticketMessageRepository repository.TicketMessageRepository) *TicketMsgServiceImpl {
	return &TicketMsgServiceImpl{
		ticketMessageRepository: ticketMessageRepository,
	}
}
