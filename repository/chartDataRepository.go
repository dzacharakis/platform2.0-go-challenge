package repository

import (
	"database/sql"
	"log"

	"bitbucket.org/go-webservice/model"
	_ "github.com/lib/pq"
)

type ChartDataRepository struct {
	Database *sql.DB
}

func (repository *ChartDataRepository) FindDataByChartID(id int) ([]model.ChartData, error) {
	sqlStatement := `
	SELECT id, x_value, y_value 
	FROM chart_data
	WHERE chart_id = $1
	`
	rows, err := repository.Database.Query(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var cd model.ChartData
	var chartDataCollection []model.ChartData

	for rows.Next() {
		err := rows.Scan(&cd.ID, &cd.XValue, &cd.YValue)
		if err != nil {
			log.Println(err)
		}

		chartDataCollection = append(chartDataCollection, cd)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return chartDataCollection, err
}
