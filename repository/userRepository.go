package repository

import (
	"database/sql"
	"errors"
	"log"

	model "bitbucket.org/go-webservice/model"
)

type UserRepository struct {
	Database *sql.DB
}

func (repository *UserRepository) FindAll() []model.User {
	sqlStatement := `
	SELECT id, username, password, email 
	FROM users
	`
	rows, err := repository.Database.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var u model.User
	var userCollection []model.User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
		if err != nil {
			log.Println(err)
		}

		userCollection = append(userCollection, u)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return userCollection
}

func (repository *UserRepository) FindByID(id int) (model.User, error) {
	sqlStatement := `
	SELECT id, username, password, email 
	FROM users 
	WHERE id = $1
	`
	var u model.User
	row := repository.Database.QueryRow(sqlStatement, id)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		WarnLogger.Println(err)
	}

	if u == (model.User{}) {
		err = errors.New("user not found")
	}

	return u, err
}
