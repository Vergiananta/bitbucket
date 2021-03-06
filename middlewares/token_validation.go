package middlewares

import (
	"bitbucket/models/dto"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateToken(user_id *dto.LoginResponse) (string, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() //Token expires after 1 month
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("API_SECRET")))

}

func TokenValid(r *http.Request) error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(r *http.Request) (string, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return "",err
	}
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return "", err
		}
		return string(uid), nil
	}
	return "", nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}