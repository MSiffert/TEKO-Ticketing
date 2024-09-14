// repository/ticket_message_repository.go

package repository

import (
	"example-rest-api/app/domain/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TicketMessageRepository interface {
	CreateTicketMessage(ticketMessage *dao.TicketMessage) (dao.TicketMessage, error)
	DeleteTicketMessage(id int) error
}

type TicketMessageRepositoryImpl struct {
	db *gorm.DB
}

func (t TicketMessageRepositoryImpl) CreateTicketMessage(ticketMessage *dao.TicketMessage) (dao.TicketMessage, error) {
	err := t.db.Save(ticketMessage).Error
	if err != nil {
		log.Error("Got an error when saving ticket message. Error: ", err)
		return dao.TicketMessage{}, err
	}
	return *ticketMessage, nil
}

func (t TicketMessageRepositoryImpl) DeleteTicketMessage(id int) error {
	err := t.db.Delete(&dao.TicketMessage{}, id).Error
	if err != nil {
		log.Error("Got an error when deleting ticket message. Error: ", err)
		return err
	}
	return nil
}

func TicketMessageRepositoryInit(db *gorm.DB) *TicketMessageRepositoryImpl {
	db.AutoMigrate(&dao.TicketMessage{})
	return &TicketMessageRepositoryImpl{
		db: db,
	}
}

func (u TicketMessageRepositoryImpl) Save(ticketMessage *dao.TicketMessage) (dao.TicketMessage, error) {
	var err = u.db.Save(ticketMessage).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.TicketMessage{}, err
	}
	return *ticketMessage, nil
}
