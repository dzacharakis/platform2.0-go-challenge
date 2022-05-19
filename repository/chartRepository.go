package repository

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"bitbucket.org/go-webservice/model"
	_ "github.com/lib/pq"
)

var WarnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

type ChartRepository struct {
	Database *sql.DB
}

func (repository *ChartRepository) FindAll() []model.Chart {
	sqlStatement := `
	SELECT asset_id, title, x_name, y_name 
	FROM chart
	`
	rows, err := repository.Database.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var c model.Chart
	var chCollection []model.Chart
	for rows.Next() {
		err := rows.Scan(&c.AssetID, &c.Title, &c.XName, &c.YName)
		if err != nil {
			log.Println(err)
		}

		chCollection = append(chCollection, c)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return chCollection
}

func (repository *ChartRepository) FindByID(id int) (model.Chart, error) {
	sqlStatement := `
	SELECT asset_id, title, x_name, y_name 
	FROM chart
	WHERE asset_id = $1
	`
	var c model.Chart
	row := repository.Database.QueryRow(sqlStatement, id)
	err := row.Scan(&c.AssetID, &c.Title, &c.XName, &c.YName)
	if err != nil {
		WarnLogger.Println(err)
	}

	if c.AssetID == 0 { // default initialization
		err = errors.New("chart not found")
	}

	return c, err
}
