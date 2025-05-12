package nenjwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Sign(payload, privateKey []byte) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	var claims jwt.MapClaims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return "", err
	}

	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(key)

}

func Verify(token string, key []byte) (*jwt.Token, error) {
	k, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return k, nil
	})
}

func Decode(token string) (string, string, error) {
	if len(token) == 0 {
		return "", "", fmt.Errorf("jwt token is required.")
	}

	t := strings.Split(token, ".")
	if len(t) != 3 {
		return "", "", fmt.Errorf("jwt token invalid format")
	}

	header, err := base64.RawURLEncoding.DecodeString(t[0])
	if err != nil {
		return "", "", err
	}

	body, err := base64.RawURLEncoding.DecodeString(t[1])
	if err != nil {
		return "", "", err
	}

	return string(header), string(body), nil
}
