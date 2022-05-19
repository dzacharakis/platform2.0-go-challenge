package repository

import (
	"database/sql"
	"errors"
	"log"

	"bitbucket.org/go-webservice/model"
	_ "github.com/lib/pq"
)

type FavouriteRepository struct {
	Database *sql.DB
}

func (repository *FavouriteRepository) FindAll() []model.Favourite {
	sqlStatement := `
	SELECT id, user_id, asset_id 
	FROM favourite
	`
	rows, err := repository.Database.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var fav model.Favourite

	var favCollection []model.Favourite
	for rows.Next() {
		err := rows.Scan(&fav.ID, &fav.UserID, &fav.AssetID)
		if err != nil {
			log.Println(err)
		}

		favCollection = append(favCollection, fav)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return favCollection
}

func (repository *FavouriteRepository) FindByID(id int) (model.Favourite, error) {
	sqlStatement := `
	SELECT id, user_id, asset_id 
	FROM favourite
	WHERE id = $1
	`
	var fav model.Favourite
	row := repository.Database.QueryRow(sqlStatement, id)
	err := row.Scan(&fav.ID, &fav.UserID, &fav.AssetID)
	if err != nil {
		WarnLogger.Println(err)
	}

	if fav == (model.Favourite{}) {
		err = errors.New("favourite not found")
	}

	return fav, err
}

func (repository *FavouriteRepository) FindFavouriteAudiencesByUserID(id int) ([]model.Variable, error) {
	sqlStatement := `
	SELECT v.asset_id, v.asset_type, v.name, v.var_type, v.possible_values, v.unit
	FROM favourite f NATURAL JOIN variable v 
	WHERE f.user_id = $1
	`
	rows, err := repository.Database.Query(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var vbl model.Variable
	var varCollection []model.Variable
	for rows.Next() {
		err := rows.Scan(&vbl.AssetID, &vbl.AssetType, &vbl.Name, &vbl.VarType, &vbl.PossibleValues, &vbl.Unit)
		if err != nil {
			log.Println(err)
		}

		varCollection = append(varCollection, vbl)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return varCollection, err
}

func (repository *FavouriteRepository) FindFavouriteChartsByUserID(id int) ([]model.Chart, error) {
	sqlStatement := `
	SELECT c.asset_id, c.asset_type, c.title, c.x_name, c.y_name 
	FROM favourite f NATURAL JOIN chart c 
	WHERE f.user_id = $1
	`
	rows, err := repository.Database.Query(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var c model.Chart
	var chCollection []model.Chart
	for rows.Next() {
		err := rows.Scan(&c.AssetID, &c.AssetType, &c.Title, &c.XName, &c.YName)
		if err != nil {
			log.Println(err)
		}

		chCollection = append(chCollection, c)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return chCollection, err
}

func (repository *FavouriteRepository) FindFavouriteInsightsByUserID(id int) ([]model.Insight, error) {
	sqlStatement := `
	SELECT i.asset_id, i.asset_type, i.chart_id, i.statement 
	FROM favourite f NATURAL JOIN insight i 
	WHERE f.user_id = $1
	`
	rows, err := repository.Database.Query(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close() // use defer mechanism to be sure

	var ins model.Insight

	var insCollection []model.Insight
	for rows.Next() {
		err := rows.Scan(&ins.AssetID, &ins.AssetType, &ins.ChartID, &ins.Statement)
		if err != nil {
			log.Println(err)
		}

		insCollection = append(insCollection, ins)
	}

	err = rows.Err()
	if err != nil {
		WarnLogger.Println(err)
	}

	return insCollection, err
}

func (repository *FavouriteRepository) DeleteFavouriteFromUser(userID, assetID int) (int, error) {
	sqlStatement := `
	DELETE 
	FROM favourite 
	WHERE user_id = $1 AND asset_id = $2`

	res, err := repository.Database.Exec(sqlStatement, userID, assetID)
	if err != nil {
		WarnLogger.Println(err)
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		WarnLogger.Println(err)
	}
	log.Printf("%d rows affected\n", count)

	return int(count), err
}

func (repository *FavouriteRepository) CreateFavourite(favourite model.Favourite) (int, error) {
	sqlStatement := `
	INSERT INTO favourite (user_id, asset_id) 
	VALUES ($1, $2)
	RETURNING id`

	var id int
	err := repository.Database.QueryRow(sqlStatement, favourite.UserID, favourite.AssetID).Scan(&id)
	if err != nil {
		WarnLogger.Println(err)
	}

	return id, err
}
