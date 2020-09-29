package service

import (
	"TemplateApi/src/models"
	"go.uber.org/zap"
)

func (s service) CreateUser(user models.User) (*models.User, error) {
	logger := s.logger.Named("s.CreateUser").With(zap.Object("user", user))

	logger.Info("executing creation of user")
	err := s.postgres.CreateUser(user)
	if err != nil {
		logger.Error("error calling dao method for create user", zap.Error(err))
		return nil, err
	}

	logger.Info("successful creation of user")
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
