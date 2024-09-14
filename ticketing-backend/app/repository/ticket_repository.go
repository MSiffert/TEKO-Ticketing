// repository/ticket_repository.go

package repository

import (
	"example-rest-api/app/domain/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TicketRepository interface {
	GetAllTickets() ([]dao.Ticket, error)
	GetTicketById(id int) (dao.Ticket, error)
	CreateTicket(ticket *dao.Ticket) (dao.Ticket, error)
	UpdateTicketStatus(id int) error
	Save(ticket *dao.Ticket) (dao.Ticket, error)
}

type TicketRepositoryImpl struct {
	db *gorm.DB
}

func (t TicketRepositoryImpl) GetAllTickets() ([]dao.Ticket, error) {
	var tickets []dao.Ticket

	err := t.db.Preload("TicketMessages").Find(&tickets).Error
	if err != nil {
		log.Error("Got an error finding all tickets. Error: ", err)
		return nil, err
	}

	return tickets, nil
}

func (t TicketRepositoryImpl) GetTicketById(id int) (dao.Ticket, error) {
	var ticket dao.Ticket

	err := t.db.Preload("TicketMessages.CreatorUser").First(&ticket, id).Error
	if err != nil {
		log.Error("Got an error when finding ticket by id. Error: ", err)
		return dao.Ticket{}, err
	}

	return ticket, nil
}

func (t TicketRepositoryImpl) CreateTicket(ticket *dao.Ticket) (dao.Ticket, error) {
	err := t.db.Create(ticket).Error
	if err != nil {
		log.Error("Got an error when creating ticket. Error: ", err)
		return dao.Ticket{}, err
	}
	return *ticket, nil
}

func (t TicketRepositoryImpl) UpdateTicketStatus(id int) error {
	// Example: This method might need to be implemented based on your logic
	// For simplicity, assuming we're updating a status field in the Ticket model
	err := t.db.Model(&dao.Ticket{}).Where("id = ?", id).Update("status", "updated_status").Error
	if err != nil {
		log.Error("Got an error when updating ticket status. Error: ", err)
		return err
	}
	return nil
}

func (t TicketRepositoryImpl) Save(ticket *dao.Ticket) (dao.Ticket, error) {
	err := t.db.Save(ticket).Error
	if err != nil {
		log.Error("Got an error when saving ticket. Error: ", err)
		return dao.Ticket{}, err
	}
	return *ticket, nil
}

func TicketRepositoryInit(db *gorm.DB) *TicketRepositoryImpl {
	db.AutoMigrate(&dao.Ticket{})
	return &TicketRepositoryImpl{db: db}
}
