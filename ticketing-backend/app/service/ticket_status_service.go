package service

import (
	"example-rest-api/app/constant"
	"example-rest-api/app/domain/dao"
	"example-rest-api/app/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
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
