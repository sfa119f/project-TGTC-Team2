package service

import (
	"tgtc/backend/database"
	"tgtc/backend/dictionary"
)

func InsertUserXBanner(userXbanner dictionary.User_X_Banner) error {
	db := database.GetDB()

	query1 :=`
		SELECT * FROM user_x_banner WHERE user_id = $1, banner_id = $2
	`
	query2 :=`
		INSERT INTO user_x_banner (user_id, banner_id) VALUES ($1, $2)
	`

	var res dictionary.User_X_Banner
	if err := db.QueryRow(query, id).Scan(&res.user_id, &res.banner_id); err != nil {
		if err == sql.ErrNoRows {
			_, err := db.Exec(query, userXbanner.UserId, userXbanner.BannerId)
			return err
		}
		fmt.Println(err)
		return res, errors.New("error")
	} else {
		err = errors.New("data is already")
	}

	return err
}