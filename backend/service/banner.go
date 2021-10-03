package service

import (
	"database/sql"
	"errors"
	"fmt"
	"tgtc/backend/database"
	"tgtc/backend/dictionary"
)

func InsertBanner(banner *dictionary.Banner) error {
	db := database.GetDB()

	query := `
	INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.Exec(query,
		banner.Name, banner.Image, banner.Url, banner.DateStart, banner.EndDate)
	
	query2 := `
	SELECT currval(pg_get_serial_sequence('banners', 'banner_id'))
	`
	
	row := db.QueryRow(query2)
	row.Scan(&banner.Id)

	return err
}

func GetBanner(id int64) (dictionary.Banner, error) {
	db := database.GetDB()

	query := `
	SELECT banner_id, banner_name, banner_image, next_url, date_start, date_end
	FROM banners WHERE banner_id = $1
	`

	var res dictionary.Banner
	if err := db.QueryRow(query, id).Scan(&res.Id, &res.Image, &res.Image, &res.Url, &res.DateStart, &res.EndDate); err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("banner not found")
		}
		fmt.Println(err)
		return res, errors.New("error")
	}
	return res, nil
}

func UpdateBanner(banner dictionary.Banner) error {
	db := database.GetDB()

	query := `
	UPDATE banners
	SET
		banner_name = $2,
		banner_image = $3,
		next_url = $4,
		date_start = $5,
		date_end = $6
	WHERE
		banner_id = $1
	`

	_, err := db.Exec(query,
		banner.Id,
		banner.Name, banner.Image, banner.Url, banner.DateStart, banner.EndDate)

	return err
}
