package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticketing-api/app/constant"
	"ticketing-api/app/domain/dao"
	"ticketing-api/app/pkg"
)

type TicketPriorityService interface {
	GetTicketPriorities(c *gin.Context)
}

type TicketPriorityServiceImpl struct {
}

func (t TicketPriorityServiceImpl) GetTicketPriorities(c *gin.Context) {
	defer pkg.PanicHandler(c)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, dao.GetPriorityMap()))
}

func TicketPriorityServiceInit() *TicketPriorityServiceImpl {
	return &TicketPriorityServiceImpl{}
}
