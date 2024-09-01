package user

import "taxi-service/internal/models"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email, password string) error {
	user := &models.User{Name: name, Email: email, Password: password}
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.ListUsers()
}
