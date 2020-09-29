package dao

import (
	"TemplateApi/src/models"
	"fmt"
	"log"
)

func (d dao) CreateUser(user models.User) error {
	query := `INSERT INTO users 
				(id, name) 
				VALUES ($1, $2)
				ON CONFLICT(id) DO NOTHING;`

	_, err := d.db.Query(query, user.ID, user.Name)
	if err != nil {
		return fmt.Errorf("error with insert statement of user in database: %+v", err)
	}

	return nil
}

func (d dao) GetUsers() ([]*models.User, error) {
	results := make([]*models.User, 0)
	query := `SELECT * FROM users;`

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error with select all statement of user in database: %+v", err)
	}

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Println(err)
		}

		results = append(results, &user)
	}

	return results, nil
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
