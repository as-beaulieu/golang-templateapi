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

	_, err := d.connection.Query(query, user.ID, user.Name)
	if err != nil {
		return fmt.Errorf("error with insert statement of user in database: %+v", err)
	}

	return nil
}

func (d dao) GetUsers() ([]*models.User, error) {
	results := make([]*models.User, 0)
	query := `SELECT * FROM users;`

	rows, err := d.connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error with select all statement of user in database: %+v", err)
	}

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Println(err)
			return nil, err
		}

		results = append(results, &user)
	}

	return results, nil
}

func (d dao) GetUserById(id string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`

	row, err := d.connection.Query(query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(&user.ID, &user.Name); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &user, nil
}

func (d dao) UpdateUser(user models.User) error {
	query := `UPDATE users
				SET name = $2
				WHERE id = $1;`

	_, err := d.connection.Query(query, user.ID, user.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d dao) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1;`

	_, err := d.connection.Query(query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
