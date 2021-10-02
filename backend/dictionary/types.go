package dictionary

import (
	"time"
)

const (
	UndisclosedError	= "Something went wrong"
	NoError				= "None"
	NotFoundError		= "Not found"
	InvalidParamError	= "Invalid parameter"
)

type User struct {
	Id			int 		`json:"id"`
	Name		string 	`json:"name"`
	Balance	int			`json:"balance"`
	Member	int			`json:"membership"`
}

type Banner struct {
	Id			int			`json:"id"`
	Name		string		`json:"banner_name"`
	Image		string		`json:"banner_image"`
	Url			string		`json:"next_url"`
	DateStart	time.Time	`json:"date_start"`
	EndDate		time.Time	`json:"date_end"`
}

type User_X_Banner struct {
	UserId			int			`json:"user_id"`
	BannerId		int			`json:"banner_id"`
}

type APIResponse struct {
	Data	interface{}	`json:"data"`
	Error	string		`json:"error"`
}
