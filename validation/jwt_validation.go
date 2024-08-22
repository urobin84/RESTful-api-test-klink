package validation

import (
	"github.com/form3tech-oss/jwt-go"
	"RESTful-api-test-klink/config"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/exception"
	"time"
)

type UserClaims struct {
	Sub      string `json:"sub"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(user entity.User) (string, error) {
	configuration := config.New()

	claims := UserClaims{
		Sub:      user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
		Role:     "Admin", // Please change using your role
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(configuration.Get("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ParseToken(stringToken string) (userAuth entity.User, err error) {
	configuration := config.New()
	token, err := jwt.ParseWithClaims(stringToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configuration.Get("JWT_SECRET")), nil
	})
	exception.PanicIfNeeded(err)

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return entity.User{
			Username: claims.Username,
			Email:    claims.Email,
		}, nil
	} else {
		return userAuth, err
	}
}
