package jwt

import (
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 10

var UserSecret = []byte("yongtenglei.user.com")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type UserClaims struct {
	Username string `json:"username"`
	jwtgo.StandardClaims
}

// GenToken 生成JWT
func GenToken(userName string) (string, error) {
	// 创建一个我们自己的声明
	c := UserClaims{
		userName, // 自定义字段
		jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "web_app",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(UserSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*UserClaims, error) {
	userClaims := new(UserClaims)
	// 解析token
	token, err := jwtgo.ParseWithClaims(tokenString, userClaims, func(token *jwtgo.Token) (i interface{}, err error) {
		return UserSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid { // 校验token
		return userClaims, nil
	}
	return nil, errors.New("Invalid token")
}
