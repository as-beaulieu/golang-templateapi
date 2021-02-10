package user

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	"go.uber.org/zap"
)

type local_service struct {
	*service.TemplateService
}

type UserOperator interface {
	CreateUser(user models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUser(userID string) error
}

func (s local_service) CreateUser(user models.User) (*models.User, error) {
	Logger := s.Logger.Named("s.CreateUser").With(zap.Object("user", user))

	Logger.Info("executing creation of user")
	err := s.Postgres.CreateUser(user)
	if err != nil {
		Logger.Error("error calling dao method for create user", zap.Error(err))
		return nil, err
	}

	Logger.Info("successful creation of user")
	return &user, nil
}

func (s *local_service) GetUsers() ([]*models.User, error) {
	Logger := s.Logger.Named("s.GetUsers")

	Logger.Info("executing get of all users")

	users, err := s.Postgres.GetUsers()
	if err != nil {
		Logger.Error("error calling dao for get all users", zap.Error(err))
		return nil, err
	}

	Logger.Info("successful get of all users from dao")
	return users, nil
}

func (s local_service) GetUserByID(userID string) (*models.User, error) {
	Logger := s.Logger.Named("s.GetUserById").With(zap.String("user_id", userID))

	Logger.Info("executing get of user by id")
	user, err := s.Postgres.GetUserById(userID)
	if err != nil {
		Logger.Error("error calling dao for get of user by id", zap.Error(err))
		return nil, err
	}

	Logger.Info("successful get of user by Id")
	return user, nil
}

func (s local_service) UpdateUser(user models.User) (*models.User, error) {
	Logger := s.Logger.Named("s.UpdateUser").With(zap.Object("user", user))

	Logger.Info("executing update user")
	err := s.Postgres.UpdateUser(user)

	if err != nil {
		Logger.Error("error calling dao for update user", zap.Error(err))
		return nil, err
	}

	Logger.Info("successful update of user")
	return nil, nil
}

func (s local_service) DeleteUser(userID string) error {
	Logger := s.Logger.Named("s.DeleteUser").With(zap.String("user_id", userID))

	Logger.Info("executing delete user by id")
	err := s.Postgres.DeleteUser(userID)
	if err != nil {
		Logger.Error("error calling dao for delete user by id", zap.Error(err))
		return err
	}

	Logger.Info("successful delete of user by id")
	return nil
}
