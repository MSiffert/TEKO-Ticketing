package config

import (
	"example-rest-api/app/controller"
	"example-rest-api/app/repository"
	"example-rest-api/app/service"
	"log"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController

	ticketRepo repository.TicketRepository
	ticketSvc  service.TicketService
	TicketCtrl controller.TicketController

	ticketMsgRepo repository.TicketMessageRepository
	ticketMsgSvc  service.TicketMsgService
	TicketMsgCtrl controller.TicketMessageController

	ticketPrioritySvc  service.TicketPriorityService
	TicketPriorityCtrl controller.TicketPriorityController

	ticketStatusSvc  service.TicketStatusService
	TicketStatusCtrl controller.TicketStatusController
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,

	ticketRepo repository.TicketRepository,
	ticketService service.TicketService,
	ticketCtrl controller.TicketController,

	ticketMessageRepo repository.TicketMessageRepository,
	ticketMessageService service.TicketMsgService,
	ticketMessageCtrl controller.TicketMessageController,

	ticketPriorityService service.TicketPriorityService,
	ticketPriorityCtrl controller.TicketPriorityController,

	ticketStatusService service.TicketStatusService,
	ticketStatusCtrl controller.TicketStatusController) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,

		ticketRepo: ticketRepo,
		ticketSvc:  ticketService,
		TicketCtrl: ticketCtrl,

		ticketMsgRepo: ticketMessageRepo,
		ticketMsgSvc:  ticketMessageService,
		TicketMsgCtrl: ticketMessageCtrl,

		ticketPrioritySvc:  ticketPriorityService,
		TicketPriorityCtrl: ticketPriorityCtrl,

		ticketStatusSvc:  ticketStatusService,
		TicketStatusCtrl: ticketStatusCtrl,
	}
}

// Init initializes and returns an Initialization struct
func Init() *Initialization {
	log.Println("Connecting to the database...")
	db := ConnectToDB() // Detailed logging handled inside ConnectToDB
	log.Println("Database connection established.")

	// Initialize repositories
	log.Println("Initializing repositories...")
	userRepo := repository.UserRepositoryInit(db)
	log.Println("User repository initialized.")
	ticketRepo := repository.TicketRepositoryInit(db)
	log.Println("Ticket repository initialized.")
	ticketMsgRepo := repository.TicketMessageRepositoryInit(db)
	log.Println("Ticket message repository initialized.")

	// Initialize services
	log.Println("Initializing services...")
	userSvc := service.UserServiceInit(userRepo)
	log.Println("User service initialized.")
	ticketSvc := service.TicketServiceInit(ticketRepo)
	log.Println("Ticket service initialized.")
	ticketMsgSvc := service.TicketMessageServiceInit(ticketMsgRepo)
	log.Println("Ticket message service initialized.")
	ticketPrioritySvc := service.TicketPriorityServiceInit()
	log.Println("Ticket priority service initialized.")
	ticketStatusSvc := service.TicketStatusServiceInit()
	log.Println("Ticket status service initialized.")

	// Initialize controllers
	log.Println("Initializing controllers...")
	userCtrl := controller.UserControllerInit(userSvc)
	log.Println("User controller initialized.")
	ticketCtrl := controller.TicketControllerInit(ticketSvc)
	log.Println("Ticket controller initialized.")
	ticketMsgCtrl := controller.TicketMessageControllerInit(ticketMsgSvc)
	log.Println("Ticket message controller initialized.")
	ticketPriorityCtrl := controller.TicketPriorityControllerInit(ticketPrioritySvc)
	log.Println("Ticket priority controller initialized.")
	ticketStatusCtrl := controller.TicketStatusControllerInit(ticketStatusSvc)
	log.Println("Ticket Status controller initialized.")

	// Return the Initialization struct with all dependencies injected
	log.Println("All dependencies have been initialized successfully. Returning the Initialization struct.")
	return NewInitialization(
		userRepo,
		userSvc,
		userCtrl,

		ticketRepo,
		ticketSvc,
		ticketCtrl,

		ticketMsgRepo,
		ticketMsgSvc,
		ticketMsgCtrl,

		ticketPrioritySvc,
		ticketPriorityCtrl,

		ticketStatusSvc,
		ticketStatusCtrl,
	)
}
