package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ticketing-api/app/domain/dao"
)

type UserRepository interface {
	GetUsersList(roleId *int) ([]dao.User, error)
	FindUserById(id int) (dao.User, error)
	Save(user *dao.User) (dao.User, error)
	DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) GetUsersList(roleId *int) ([]dao.User, error) {
	var users []dao.User

	query := u.db.Preload("Role")

	if roleId != nil {
		query = query.Where("role_id = ?", *roleId)
	}

	// Execute the query
	err := query.Find(&users).Error
	if err != nil {
		log.Error("Got an error finding users. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
	user := dao.User{
		ID: id,
	}
	err := u.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when saving user. Error: ", err)
		return dao.User{}, err
	}

	retrievedUser, err := u.FindUserById(user.ID)
	if err != nil {
		log.Error("Got an error when retrieving user by ID after saving. Error: ", err)
		return dao.User{}, err
	}
	return retrievedUser, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&dao.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})

	if err := db.AutoMigrate(&dao.Role{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Check if roles exist, if not, insert test data
	var count int64
	db.Model(&dao.Role{}).Count(&count)
	if count == 0 {
		roles := []dao.Role{
			{Role: "Supporter"},
			{Role: "Sachbearbeiter"},
		}

		if err := db.Create(&roles).Error; err != nil {
			log.Fatal("failed to insert roles:", err)
		} else {
			log.Println("Test roles inserted successfully")
		}
	} else {
		log.Println("Roles already exist, skipping seeding")
	}

	return &UserRepositoryImpl{
		db: db,
	}
}
