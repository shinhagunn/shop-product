package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/zsmartex/pkg/jwt"
)

func UpperFirstLetter(s string) string {
	return strings.Title(strings.ToLower(s))
}

func RandomUID() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9999999999
	return fmt.Sprintf("UID%v", rand.Intn(max-min+1)+min)
}

func CheckJWT(token string) (*jwt.Auth, error) {
	ks := jwt.KeyStore{}
	ks.LoadPublicKeyFromFile("./public.key")
	j, err := jwt.ParseAndValidate(token, ks.PublicKey)

	if err != nil {
		return nil, err
	}

	return &j, nil
}
