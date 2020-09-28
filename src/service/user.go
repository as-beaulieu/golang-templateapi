package service

import "TemplateApi/src/models"

func (s service) CreateUser(user models.User) (*models.User, error) {
	err := s.postgres.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s service) GetUsers() ([]*models.User, error) {
	users, err := s.postgres.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s service) GetUserByID(userID string) (*models.User, error) {
	user, err := s.postgres.GetUserById(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) UpdateUser(user models.User) (*models.User, error) {
	err := s.postgres.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s service) DeleteUser(userID string) error {
	err := s.postgres.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}
