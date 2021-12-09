package mysql

import (
	"app/models"
	"app/pkg/MD5"
	"database/sql"
	"errors"
)

func IsUserExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func QueryByUserName() error {
	return nil
}

func InsertUser(user *models.User) error {
	sqlStr := `insert into user (user_id, username, password) values (?, ?, ?)`
	if _, err := db.Exec(sqlStr, user.UserID, user.UserName, user.Password); err != nil {
		return err
	}
	return nil
}

func LogIn(user *models.User) error {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err := db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return errors.New("Invalid username or password")
	}
	if err != nil {
		return err
	}

	if MD5.EncryptPassword(oPassword) != user.Password {
		return errors.New("Invalid password")
	}

	return nil

}

func GetUserByID(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
