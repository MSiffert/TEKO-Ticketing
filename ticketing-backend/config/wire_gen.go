// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"ticketing-api/app/controller"
	"ticketing-api/app/repository"
	"ticketing-api/app/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

// WireInit function to initialize everything using Wire
func WireInit() *Initialization {
	gormDB := ConnectToDB()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := service.UserServiceInit(userRepositoryImpl)
	userControllerImpl := controller.UserControllerInit(userServiceImpl)
	ticketRepositoryImpl := repository.TicketRepositoryInit(gormDB)
	ticketServiceImpl := service.TicketServiceInit(ticketRepositoryImpl)
	ticketControllerImpl := controller.TicketControllerInit(ticketServiceImpl)
	ticketMessageRepositoryImpl := repository.TicketMessageRepositoryInit(gormDB)
	ticketMsgServiceImpl := service.TicketMessageServiceInit(ticketMessageRepositoryImpl)
	ticketMessageControllerImpl := controller.TicketMessageControllerInit(ticketMsgServiceImpl)
	ticketPriorityServiceImpl := service.TicketPriorityServiceInit()
	ticketPriorityControllerImpl := controller.TicketPriorityControllerInit(ticketPriorityServiceImpl)
	ticketStatusServiceImpl := service.TicketStatusServiceInit()
	ticketStatusControllerImpl := controller.TicketStatusControllerInit(ticketStatusServiceImpl)
	initialization := NewInitialization(userRepositoryImpl, userServiceImpl, userControllerImpl, ticketRepositoryImpl, ticketServiceImpl, ticketControllerImpl, ticketMessageRepositoryImpl, ticketMsgServiceImpl, ticketMessageControllerImpl, ticketPriorityServiceImpl, ticketPriorityControllerImpl, ticketStatusServiceImpl, ticketStatusControllerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

// Existing user-related sets
var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var userCtrlSet = wire.NewSet(controller.UserControllerInit, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

// Define sets for Ticket and TicketMessage components
var ticketRepoSet = wire.NewSet(repository.TicketRepositoryInit, wire.Bind(new(repository.TicketRepository), new(*repository.TicketRepositoryImpl)))

var ticketServiceSet = wire.NewSet(service.TicketServiceInit, wire.Bind(new(service.TicketService), new(*service.TicketServiceImpl)))

var ticketCtrlSet = wire.NewSet(controller.TicketControllerInit, wire.Bind(new(controller.TicketController), new(*controller.TicketControllerImpl)))

var ticketMessageRepoSet = wire.NewSet(repository.TicketMessageRepositoryInit, wire.Bind(new(repository.TicketMessageRepository), new(*repository.TicketMessageRepositoryImpl)))

var ticketMessageServiceSet = wire.NewSet(service.TicketMessageServiceInit, wire.Bind(new(service.TicketMsgService), new(*service.TicketMsgServiceImpl)))

var ticketMessageCtrlSet = wire.NewSet(controller.TicketMessageControllerInit, wire.Bind(new(controller.TicketMessageController), new(*controller.TicketMessageControllerImpl)))

var ticketPriorityServiceSet = wire.NewSet(service.TicketPriorityServiceInit, wire.Bind(new(service.TicketPriorityService), new(*service.TicketPriorityServiceImpl)))

var ticketPriorityCtrlSet = wire.NewSet(controller.TicketPriorityControllerInit, wire.Bind(new(controller.TicketPriorityController), new(*controller.TicketPriorityControllerImpl)))

var ticketStatusServiceSet = wire.NewSet(service.TicketStatusServiceInit, wire.Bind(new(service.TicketStatusService), new(*service.TicketStatusServiceImpl)))

var ticketStatusCtrlSet = wire.NewSet(controller.TicketStatusControllerInit, wire.Bind(new(controller.TicketStatusController), new(*controller.TicketStatusControllerImpl)))
