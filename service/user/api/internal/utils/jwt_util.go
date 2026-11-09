package utils

import (
	"encoding/base64"
	"infinitechatgo/service/user/api/internal/types/constants"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func getSignInKey() []byte {
	keyBytes, err := base64.StdEncoding.DecodeString(constants.TokenSecretKey)
	if err != nil {
		return nil
	}

	return keyBytes
}

/*
 * 生成JWT（支持自定义过期时间）
 */
func GenerateAccessToken(userId int64) string {

	now := time.Now()

	claims := jwt.MapClaims{
		"subject":    strconv.FormatInt(userId, 10),
		"issuedAt":   now,
		"expiration": now.Add(time.Duration(30) * time.Minute),
	}

	res, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(getSignInKey())
	return res
}
