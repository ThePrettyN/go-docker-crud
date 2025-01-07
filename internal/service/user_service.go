package service

import "go-docker-crud/internal/repository"

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{Repo: repo}
}

func (s *UserService) CreateUser(user repository.User) (repository.User, error) {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]repository.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (repository.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(id int, user repository.User) (repository.User, error) {
	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
