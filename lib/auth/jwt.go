package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

const (
	UsernameClaim       = "username"
	RoleClaim           = "role"
	AccountIdClimeTitle = "id"
)

var (
	ErrorInvalidToken = errors.New("Некорректный токен")
)

type TokenPayload struct {
	Username string
	Role     string
	Id       float64
}

func GetPayloadAndValidate(jwtToken, jwtSecret string) (TokenPayload, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return TokenPayload{}, errors.Wrap(ErrorInvalidToken, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return TokenPayload{}, ErrorInvalidToken
	}

	var payload TokenPayload

	payload.Username, ok = claims[UsernameClaim].(string)
	if !ok {
		return TokenPayload{}, ErrorInvalidToken
	}
	payload.Role, ok = claims[RoleClaim].(string)
	if !ok {
		return TokenPayload{}, ErrorInvalidToken
	}

	payload.Id, ok = claims[AccountIdClimeTitle].(float64)
	if !ok {
		return TokenPayload{}, ErrorInvalidToken
	}

	return payload, nil
}
