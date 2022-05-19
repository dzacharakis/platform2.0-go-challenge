package repository

import (
	"database/sql"
	"errors"
	"log"

	"bitbucket.org/go-webservice/model"
	_ "github.com/lib/pq"
)

type AudienceRepository struct {
	Database *sql.DB
}

func (repository *AudienceRepository) FindAll() []model.Variable {
	sqlStatement := `
	SELECT asset_id, name, var_type, possible_values, unit
	FROM variable
	`
	rows, err := repository.Database.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var vbl model.Variable

	var varCollection []model.Variable
	for rows.Next() {
		err := rows.Scan(&vbl.AssetID, &vbl.Name, &vbl.VarType, &vbl.PossibleValues, &vbl.Unit)
		if err != nil {
			log.Println(err)
		}

		varCollection = append(varCollection, vbl)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return varCollection
}

func (repository *AudienceRepository) FindByID(id int) (model.Variable, error) {
	sqlStatement := `
	SELECT asset_id, name, var_type, possible_values, unit 
	FROM variable 
	WHERE asset_id = $1
	`

	var vbl model.Variable
	row := repository.Database.QueryRow(sqlStatement, id)
	err := row.Scan(&vbl.AssetID, &vbl.Name, &vbl.VarType, &vbl.PossibleValues, &vbl.Unit)
	if err != nil {
		WarnLogger.Println(err)
	}

	if vbl == (model.Variable{}) {
		err = errors.New("variable not found")
	}

	return vbl, err
}
