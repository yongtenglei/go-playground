package models

type User struct {
	// json:"strind" means int64 can be converted to string when Marshaling to Json
	// and string type in Json can be converted to int64 when Unmarshaling to User struct
	// because the range of int64 > range of JSON number type
	UserID   int64  `json:"user_id,string" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `db:"password"`
}
