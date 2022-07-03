package entity

type IStorage interface {
	CreateUser(user User) (*User, error)
	DeleteUser(userId int64) error
	FindUser(userId int64) (*User, error)
	UpdateUser(userId int64, user User) (*User, error)
}

type IUsecase interface {
	CreateUser(user User) (*User, error)
	DeleteUser(userId int64) error
	FindUser(userId int64) (*User, error)
	UpdateUser(id int64, user User) (*User, error)
}

type usecase struct {
	storage *IStorage
}

func New(storage *IStorage) IUsecase {
	return &usecase{storage: storage}
}

func (u *usecase) CreateUser(user User) (*User, error) {
	return (*u.storage).CreateUser(user)
}

func (u *usecase) DeleteUser(userId int64) error {
	return (*u.storage).DeleteUser(userId)
}

func (u *usecase) FindUser(userId int64) (*User, error) {
	return (*u.storage).FindUser(userId)
}

func (u *usecase) UpdateUser(id int64, user User) (*User, error) {
	return (*u.storage).UpdateUser(id, user)
}