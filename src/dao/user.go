package dao

import (
	"TemplateApi/src/models"
	"fmt"
)

func (d dao) CreateUser(user models.User) error {
	query := `INSERT INTO users 
				(id, name) 
				VALUES $1, $2
				ON CONFLICT(id) DO NOTHING;`

	result, err := d.db.Query(query, user.ID, user.Name)
	if err != nil {
		return fmt.Errorf("error with insert statement of user in database: %+v", err)
	}

	if !result.Next() {
		return fmt.Errorf("error insert statement of user in database failed to save: %+v", err)
	}

	return nil
}

func (d dao) GetUsers() ([]*models.User, error) {

	return nil, nil
}

func (d dao) GetUserById(id string) (*models.User, error) {

	return nil, nil
}

func (d dao) UpdateUser(user models.User) error {

	return nil
}

func (d dao) DeleteUser(id string) error {

	return nil
}
