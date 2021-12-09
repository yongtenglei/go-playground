package mysql

import (
	"fmt"
	"go-blog/models"
)

func SignIn(user models.User) (err error) {
	return db.Create(&user).Error
}

//func LogIn(user models.User) (err error) {
//var u models.User
//db.First(&u)
//if u.Password == user.Password {
//return
//}

//return
//}

func UserExist(user models.User) (models.User, bool) {
	var u models.User
	//db.Model(models.User{}).Where("username = ?", user.Username).First(&u)
	db.Model(models.User{}).Where("username = ?", user.Username).First(&u)
	fmt.Println(u)

	if u.ID > 0 {
		return u, true
	}

	return models.User{}, false
}
