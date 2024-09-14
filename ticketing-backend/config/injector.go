//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"ticketing-api/app/controller"
	"ticketing-api/app/repository"
	"ticketing-api/app/service"
)

var db = wire.NewSet(ConnectToDB)

// Existing user-related sets
var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

// Define sets for Ticket and TicketMessage components
var ticketRepoSet = wire.NewSet(repository.TicketRepositoryInit,
	wire.Bind(new(repository.TicketRepository), new(*repository.TicketRepositoryImpl)),
)

var ticketServiceSet = wire.NewSet(service.TicketServiceInit,
	wire.Bind(new(service.TicketService), new(*service.TicketServiceImpl)),
)

var ticketCtrlSet = wire.NewSet(controller.TicketControllerInit,
	wire.Bind(new(controller.TicketController), new(*controller.TicketControllerImpl)),
)

var ticketMessageRepoSet = wire.NewSet(repository.TicketMessageRepositoryInit,
	wire.Bind(new(repository.TicketMessageRepository), new(*repository.TicketMessageRepositoryImpl)),
)

var ticketMessageServiceSet = wire.NewSet(service.TicketMessageServiceInit,
	wire.Bind(new(service.TicketMsgService), new(*service.TicketMsgServiceImpl)),
)

var ticketMessageCtrlSet = wire.NewSet(controller.TicketMessageControllerInit,
	wire.Bind(new(controller.TicketMessageController), new(*controller.TicketMessageControllerImpl)),
)

var ticketPriorityServiceSet = wire.NewSet(service.TicketPriorityServiceInit,
	wire.Bind(new(service.TicketPriorityService), new(*service.TicketPriorityServiceImpl)),
)

var ticketPriorityCtrlSet = wire.NewSet(controller.TicketPriorityControllerInit,
	wire.Bind(new(controller.TicketPriorityController), new(*controller.TicketPriorityControllerImpl)),
)

var ticketStatusServiceSet = wire.NewSet(service.TicketStatusServiceInit,
	wire.Bind(new(service.TicketStatusService), new(*service.TicketStatusServiceImpl)),
)

var ticketStatusCtrlSet = wire.NewSet(controller.TicketStatusControllerInit,
	wire.Bind(new(controller.TicketStatusController), new(*controller.TicketStatusControllerImpl)),
)

// WireInit function to initialize everything using Wire
func WireInit() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		userCtrlSet,
		userServiceSet,
		userRepoSet,
		ticketRepoSet,
		ticketServiceSet,
		ticketCtrlSet,
		ticketMessageRepoSet,
		ticketMessageServiceSet,
		ticketMessageCtrlSet,
		ticketPriorityServiceSet,
		ticketPriorityCtrlSet,
		ticketStatusServiceSet,
		ticketStatusCtrlSet,
	)
	return nil
}
