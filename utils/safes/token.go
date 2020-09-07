// token
package safes

import (
	"crypto/rsa"
	"errors"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"daily/cmd/common"
	"daily/cmd/logger"
)

const (
	tokenTime = time.Hour * 720
)

var privateKey = string(common.PrivateKeyString)
var publicKey = string(common.PublicKeyString)

// Create 根据登录时间生成token
func Create(userId, role string, loginTime int64, license bool) (string, time.Time, error) {
	var tokenSring string
	var expire time.Time
	var err error
	// Create the Claims
	var expireTime = time.Now().Add(tokenTime)
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := token.Claims.(jwt.MapClaims)

	// 设置token参数
	expire = expireTime
	claims["iss"] = "login_token"
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	claims["userid"] = userId
	claims["role"] = role
	claims["logintime"] = strconv.FormatInt(loginTime, 10)
	claims["license"] = license
	// 生成token字符串
	key, err := rsaPrivateKey(privateKey)
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field(err))
		return "", expire, err
	}
	tokenSring, err = token.SignedString(key)
	return tokenSring, expire, err
}

// Parse 深入解析token
func Parse(tokenString string) (map[string]interface{}, error) {
	key, err := rsaPublickey(publicKey)
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field(err))
		return nil, err
	}
	//	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
	//		return key, nil
	//	})
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, errors.New("invalid signing algorithm")
		}
		return key, nil
	})
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field(err))
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// Refresh 刷新token
func Refresh(token *jwt.Token, role, userId string, nowTime int64, license bool) (string, time.Time, error) {
	var expire = time.Now().Add(tokenTime)
	var tokenString string
	claims := token.Claims.(jwt.MapClaims)

	// new sign
	claims["iss"] = "refresh_token"
	claims["exp"] = expire.Unix()
	claims["iat"] = nowTime
	claims["license"] = license
	claims["userid"] = userId
	claims["role"] = role
	claims["logintime"] = strconv.FormatInt(nowTime, 10)
	claims["license"] = license

	key, err := rsaPrivateKey(privateKey)
	if err != nil {
		return tokenString, expire, nil
	}
	tokenString, err = token.SignedString(key)
	return tokenString, expire, err
}

// Get 把http-haed里面的token给初步解析出来，解析结果为token字符串
func Get(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("auth header empty")
	}
	// 切割token,使其成为数组
	parts := strings.SplitN(authorization, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("invalid auth header")
	}
	return parts[1], nil
}

// TokenFormat token的string转为*jwt.token
func Format(tokenString string) (*jwt.Token, error) {
	key, err := rsaPublickey(publicKey)
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field(err))
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, errors.New("invalid signing algorithm")
		}
		return key, nil
	})
	return token, err
}

// 将私钥钥用rsa算法转换
func rsaPrivateKey(key string) (*rsa.PrivateKey, error) {
	keys := []byte(key)
	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(keys)
	return rsaKey, err
}

//将公钥用rsa算法转换
func rsaPublickey(key string) (*rsa.PublicKey, error) {
	keys := []byte(key)
	rsaKey, err := jwt.ParseRSAPublicKeyFromPEM(keys)
	return rsaKey, err
}
