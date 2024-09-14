package router

import (
	"example-rest-api/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	api := router.Group("/api")
	{
		user := api.Group("/users")
		user.GET("", init.UserCtrl.GetUsers)
		user.POST("", init.UserCtrl.CreateUser)

		ticketpriority := api.Group("/ticketpriorities")
		ticketpriority.GET("", init.TicketPriorityCtrl.GetTicketPriorities)

		ticketstatus := api.Group("/ticketstatus")
		ticketstatus.GET("", init.TicketStatusCtrl.GetTicketStatus)

		ticket := api.Group("/tickets")
		ticket.GET("", init.TicketCtrl.GetAllTickets)
		ticket.GET(":ticketID", init.TicketCtrl.GetTicketDetails)
		ticket.POST("", init.TicketCtrl.CreateTicket)
		ticket.PUT("/:ticketID", init.TicketCtrl.UpdateTicket)

		ticketMessage := api.Group("/ticketmessages")
		ticketMessage.POST("", init.TicketMsgCtrl.CreateTicketMessage)
		ticketMessage.DELETE("/:ticketMessageID", init.TicketMsgCtrl.DeleteTicketMessage)
	}

	return router
}
