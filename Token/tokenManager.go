package Token

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GetToken(username string) string {
	//Creating Access Token
	//os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return token

}