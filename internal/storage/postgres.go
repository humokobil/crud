package storage

import (
	"errors"
	"otus/crud/v1/internal/entity"

	"gorm.io/gorm"
)

// pgx driver


type psDB struct {
	db      *gorm.DB
}

func New(db *gorm.DB, migration bool) entity.IStorage {
	if migration {
		db.AutoMigrate(entity.User{})
	}
	return &psDB{db: db}
}

func (ps *psDB) CreateUser(user entity.User) (*entity.User, error) {
	err := (*ps.db).Create(&user).Error
	return &user, err
}

func (ps *psDB) DeleteUser(userId int64) error {
	err := (*ps.db).Model(entity.User{}).Delete(&entity.User{Id : userId}).Error
	return err
}

func (ps *psDB) FindUser(userId int64) (*entity.User, error) {
	user := entity.User{Id : userId}
	err := (*ps.db).Model(entity.User{}).First(&user).Error
	return &user, err
}

func (ps *psDB) UpdateUser(id int64, user entity.User) (*entity.User, error) {
	tx := (*ps.db).Model(entity.User{Id : id}).Updates(&user)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	if tx.RowsAffected == 0 {
		err = errors.New("User not found")
		return nil, err
	}
	return &user, err
}