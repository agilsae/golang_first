package helper

import (
	"diary-api/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))


func GenerateJWT(user model.User) (string, time.Time, time.Time, error) {
    var tmIat time.Time
    var tmEat time.Time

    tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  user.ID,
        "iat": time.Now().Unix(),
        "eat": time.Now().Add(time.Hour * time.Duration(tokenTTL)).Unix(),
    })

    key, err := token.SignedString(privateKey)

    claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return key, tmIat, tmEat, errors.New("couldn't parse claims")
	}

    fmt.Println("claims", claims)

    switch iat := claims["iat"].(type) {
    case int64:
        tmIat = time.Unix(iat, 0)
    }

    switch eat := claims["eat"].(type) {
    case int64:
        tmEat = time.Unix(eat, 0)
    }

    return key, tmIat, tmEat, err
}

func ValidateJWT(context *gin.Context) error {
    token, err := getToken(context)
    if err != nil {
        return err
    }
    // _, ok := token.Claims.(jwt.MapClaims)
    // if ok && token.Valid {
    //     return nil
    // }
    // return errors.New("invalid token provided")

    claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("couldn't parse claims")
	}
    var tm time.Time
    switch eat := claims["eat"].(type) {
    case float64:
        tm = time.Unix(int64(eat), 0)
    case json.Number:
        v, _ := eat.Int64()
        tm = time.Unix(v, 0)
    }

    fmt.Println("eat", tm, "local tm", time.Unix(time.Now().Local().Unix(),0))
	if tm.Unix() < time.Now().Local().Unix() {
		return  errors.New("token expired")
		
	}

	return nil
}


func CurrentUser(context *gin.Context) (model.User, error) {
    err := ValidateJWT(context)
    if err != nil {
        return model.User{}, err
    }
    token, _ := getToken(context)
    claims, _ := token.Claims.(jwt.MapClaims)
    userId := uint(claims["id"].(float64))

    user, err := model.FindUserById(userId)
    if err != nil {
        return model.User{}, err
    }
    return user, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
    tokenString := getTokenFromRequest(context)
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return privateKey, nil
    })
    return token, err
}

func getTokenFromRequest(context *gin.Context) string {
    bearerToken := context.Request.Header.Get("Authorization")
    splitToken := strings.Split(bearerToken, " ")
    if len(splitToken) == 2 {
        return splitToken[1]
    }
    return ""
}

