package logic

import (
	"app/dao/mysql"
	"app/models"
	"app/pkg/MD5"
	"app/pkg/jwt"
	"app/pkg/snowflake"
	"errors"
	"log"
)

func SignUp(p *models.ParamSignUp) error {
	// 判断是否存在
	exist, err := mysql.IsUserExist(p.UserName)
	if err != nil {
		// 数据库查询出错
		log.Printf("Query in IsUserExist failed: %v\n", err)
		return err
	}
	if exist {
		return errors.New("User already exist")
	}

	// 生成UID
	var userID int64
	userID = snowflake.GenID()

	log.Println("Generate userID", userID)

	//密码加密
	userPassword := MD5.EncryptPassword(p.Password)

	//  构建User实例
	user := &models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: userPassword,
	}

	// 保存到数据库
	if err := mysql.InsertUser(user); err != nil {
		log.Printf("InsertUser failed: %v\n", err)
		return err
	}

	return nil
}

func LogIn(p *models.ParamLogIn) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}

	if err := mysql.LogIn(user); err != nil {
		return "", err
	}

	// 生成jwt
	return jwt.GenToken(user.UserID, user.UserName)
}
