package service

import (
	"example-rest-api/app/constant"
	"example-rest-api/app/domain/dao"
	"example-rest-api/app/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
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
