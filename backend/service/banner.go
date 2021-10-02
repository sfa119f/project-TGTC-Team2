package service

import (
	"tgtc/backend/database"
	"tgtc/backend/dictionary"
)

func InsertBanner(banner dictionary.Banner) error {
	db := database.GetDB()

	query := `
	INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.Exec(query,
		banner.Name, banner.Image, banner.Url, banner.DateStart, banner.EndDate)

	// sampe sini banner id masih kosong

	return err
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
