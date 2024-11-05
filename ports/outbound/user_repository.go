package ports

import domain "hexagonal/domain/entities"

type UserRepository interface {
	Save(user domain.User) error
	FindById(id string) (domain.User, error)
	FindAll() ([]domain.User, error)
	Update(user domain.User) error
	Delete(id string) error
}
