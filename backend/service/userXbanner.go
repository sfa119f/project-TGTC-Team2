package service

import (
	"tgtc/backend/database"
	"tgtc/backend/dictionary"
)

func InsertUserXBanner(userXbanner dictionary.User_X_Banner) error {
	db := database.GetDB()

	query :=`
		INSERT INTO user_x_banner (user_id, banner_id) VALUES ($1, $2)
	`

	_, err := db.Exec(query, userXbanner.UserId, userXbanner.BannerId)

	return err
}

func GetBannersOfUser(userId int64) ([]int, error) {
	db := database.GetDB()

	query :=`
	SELECT banner_id 
	FROM user_x_banner 
	WHERE user_id = $1
	`

	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners_id []int
	for rows.Next() {
		var uXb dictionary.User_X_Banner
		if err:= rows.Scan(&uXb.BannerId); err != nil {
			return banners_id, err
		}
		banners_id = append(banners_id, uXb.BannerId)
	}
	if err = rows.Err(); err != nil {
		return banners_id, err
	}
	return banners_id, nil
}
