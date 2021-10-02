package service

import (
	"database/sql"
	"errors"
	"fmt"
	"tgtc/backend/database"
	"tgtc/backend/dictionary"
)

func InsertUser(user dictionary.User) error {
	db := database.GetDB()

	query :=`
	NSERT INTO users (user_name, balance, membership)
	VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, user.Name, user.Balance, user.Member)

	return err
}

func GetUser(id int64) (dictionary.User, error) {
	db := database.GetDB()

	query := `
		SELECT * FROM users WHERE user_id = $1
	`
	var res dictionary.User
	if err := db.QueryRow(query, id).Scan(&res.Id, &res.Name, &res.Balance, &res.Member); err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("user nil")
		}
		fmt.Println(err)
		return res, errors.New("error")
	}
	return res, nil
}