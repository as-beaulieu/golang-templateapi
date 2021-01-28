package service

import (
	"TemplateApi/src/models"
	"go.uber.org/zap"
)

type UserOperator interface {
	CreateUser(user models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUser(userID string) error
}

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
	logger := s.logger.Named("s.GetUsers")

	logger.Info("executing get of all users")

	users, err := s.postgres.GetUsers()
	if err != nil {
		logger.Error("error calling dao for get all users", zap.Error(err))
		return nil, err
	}

	logger.Info("successful get of all users from dao")
	return users, nil
}

func (s service) GetUserByID(userID string) (*models.User, error) {
	logger := s.logger.Named("s.GetUserById").With(zap.String("user_id", userID))

	logger.Info("executing get of user by id")
	user, err := s.postgres.GetUserById(userID)
	if err != nil {
		logger.Error("error calling dao for get of user by id", zap.Error(err))
		return nil, err
	}

	logger.Info("successful get of user by Id")
	return user, nil
}

func (s service) UpdateUser(user models.User) (*models.User, error) {
	logger := s.logger.Named("s.UpdateUser").With(zap.Object("user", user))

	logger.Info("executing update user")
	err := s.postgres.UpdateUser(user)

	if err != nil {
		logger.Error("error calling dao for update user", zap.Error(err))
		return nil, err
	}

	logger.Info("successful update of user")
	return nil, nil
}

func (s service) DeleteUser(userID string) error {
	logger := s.logger.Named("s.DeleteUser").With(zap.String("user_id", userID))

	logger.Info("executing delete user by id")
	err := s.postgres.DeleteUser(userID)
	if err != nil {
		logger.Error("error calling dao for delete user by id", zap.Error(err))
		return err
	}

	logger.Info("successful delete of user by id")
	return nil
}
