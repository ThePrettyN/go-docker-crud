package repository

import "gorm.io/gorm"

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=0"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(id int, user User) (User, error) {
	if err := r.db.Model(&User{ID: uint(id)}).Updates(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
