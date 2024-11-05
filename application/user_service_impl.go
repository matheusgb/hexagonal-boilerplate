package application

import (
	domain "hexagonal/domain/entities"
	ports "hexagonal/ports/outbound"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user domain.User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetById(id string) (domain.User, error) {
	return s.repo.FindById(id)
}

func (s *UserService) GetAll() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Update(user domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
