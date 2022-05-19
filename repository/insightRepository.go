package repository

import (
	"database/sql"
	"errors"
	"log"

	"bitbucket.org/go-webservice/model"
	_ "github.com/lib/pq"
)

type InsightRepository struct {
	Database *sql.DB
}

func (repository *InsightRepository) FindAll() []model.Insight {
	sqlStatement := `
	SELECT asset_id, chart_id, statement 
	FROM insight
	`
	rows, err := repository.Database.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var ins model.Insight
	var insCollection []model.Insight
	for rows.Next() {
		err := rows.Scan(&ins.AssetID, &ins.ChartID, &ins.Statement)
		if err != nil {
			log.Println(err)
		}

		insCollection = append(insCollection, ins)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return insCollection
}

func (repository *InsightRepository) FindByID(id int) (model.Insight, error) {
	sqlStatement := `
	SELECT asset_id, chart_id, statement 
	FROM insight 
	WHERE asset_id = $1
	`

	var ins model.Insight
	row := repository.Database.QueryRow(sqlStatement, id)
	err := row.Scan(&ins.AssetID, &ins.ChartID, &ins.Statement)
	if err != nil {
		WarnLogger.Println(err)
	}

	if ins == (model.Insight{}) {
		err = errors.New("insight not found")
	}

	return ins, err
}
