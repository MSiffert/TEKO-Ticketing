package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticketing-api/app/constant"
	"ticketing-api/app/domain/dao"
	"ticketing-api/app/pkg"
)

type TicketStatusService interface {
	GetTicketStatus(c *gin.Context)
}

type TicketStatusServiceImpl struct {
}

func (t TicketStatusServiceImpl) GetTicketStatus(c *gin.Context) {
	defer pkg.PanicHandler(c)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, dao.GetStatusMap()))
}

func TicketStatusServiceInit() *TicketStatusServiceImpl {
	return &TicketStatusServiceImpl{}
}
