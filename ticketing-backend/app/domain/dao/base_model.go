package dao

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

type User struct {
	ID       int    `gorm:"column:id; primary_key; auto_increment" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password;->:false" json:"-"`
	RoleID   int    `gorm:"column:role_id;not null" json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}

type Role struct {
	ID   int    `gorm:"column:id; primary_key; auto_increment" json:"id"`
	Role string `gorm:"column:role" json:"role"`
	BaseModel
}

type Ticket struct {
	ID              int             `gorm:"column:id; primary_key; auto_increment" json:"id"`
	Title           string          `gorm:"column:title" json:"title"`
	Description     string          `gorm:"column:description" json:"description"`
	Status          TicketStatus    `gorm:"column:status" json:"status"`
	Priority        TicketPriority  `gorm:"column:priority" json:"priority"`
	CreatorUserID   int             `gorm:"column:creator_user_id;not null" json:"creator_user_id"`
	CreatorUser     User            `gorm:"foreignKey:CreatorUserID;references:ID" json:"creator_user"`
	SupporterUserID *int            `gorm:"column:supporter_user_id;null" json:"supporter_user_id"` // Ensure nullable
	SupporterUser   User            `gorm:"foreignKey:SupporterUserID;references:ID" json:"supporter_user"`
	TicketMessages  []TicketMessage `gorm:"foreignKey:TicketID" json:"ticket_messages"` // This defines the relationship
	BaseModel
}

type TicketMessage struct {
	ID            int    `gorm:"column:id; primary_key; auto_increment" json:"id"`
	Text          string `gorm:"column:text" json:"text"`
	TicketID      int    `gorm:"column:ticket_id" json:"ticket_id"`
	CreatorUserID int    `gorm:"column:creator_user_id;not null" json:"creator_user_id"`
	CreatorUser   User   `gorm:"foreignKey:CreatorUserID;references:ID" json:"creator_user"`
	BaseModel
}
