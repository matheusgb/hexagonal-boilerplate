package in_memory

import (
	"errors"
	domain "hexagonal/domain/entities"
)

type InMemoryUserRepo struct {
	users map[string]domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make(map[string]domain.User)}
}

func (r *InMemoryUserRepo) Save(user domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindById(id string) (domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}

	// chama dto de output e retorna (não usar domain)
	return user, nil
}

func (r *InMemoryUserRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	for _, user := range r.users {
		users = append(users, user)
	}

	// chama dto de output e retorna (não usar domain)
	return users, nil
}

func (r *InMemoryUserRepo) Update(user domain.User) error {
	_, exists := r.users[user.ID]
	if !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user

	// chama dto de output e retorna (não usar domain)
	return nil
}

func (r *InMemoryUserRepo) Delete(id string) error {
	delete(r.users, id)
	return nil
}
