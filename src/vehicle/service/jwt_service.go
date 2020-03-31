package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/response"
)

/**
第一部分我们称它为头部（header)
第二部分我们称其为载荷（payload, 类似于飞机上承载的物品)
第三部分是签证（signature)

1、header:
声明类型，这里是jwt
声明加密的算法 通常直接使用 HMAC SHA256

2、playload:
标准中注册的声明
公共的声明
私有的声明

Audience  string `json:"aud,omitempty"`
ExpiresAt int64  `json:"exp,omitempty"`
Id        string `json:"jti,omitempty"`
IssuedAt  int64  `json:"iat,omitempty"`
Issuer    string `json:"iss,omitempty"`
NotBefore int64  `json:"nbf,omitempty"`
Subject   string `json:"sub,omitempty"`

iss: jwt签发者
sub: jwt所面向的用户
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
nbf: 定义在什么时间之前，该jwt都是不可用的.
iat: jwt的签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击

3、signature:
header (base64后的)
payload (base64后的)
secret
var encodedString = base64UrlEncode(header) + '.' + base64UrlEncode(payload);
var signature = HMACSHA256(encodedString, 'secret');
 */
var (
	ExpiresAt = time.Now().Add(conf.Expires * time.Hour).Unix()
	//token过期
	TokenExpired      = errors.New(response.TokenExpiredStr)
	//token未激活
	TokenNotValidYet  = errors.New(response.TokenNotValidYetStr)
	//token不合法
	TokenMalformed    = errors.New(response.TokenMalformedStr)
	//token未知
	TokenInvalid      = errors.New(response.TokenInvalidStr)
	//签名信息错误，无法验证token
	ValidationErrorUnverifiable = errors.New(response.ValidationErrorUnverifiableStr)
	//签名验证失败
	ValidationErrorSignatureInvalid = errors.New(response.ValidationErrorSignatureInvalidStr)
	Jwt *JWT
	SignKey = conf.SignKey
)
func init() {
	Jwt = newJWT()
}

func newJWT() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}
type JWT struct {
	SigningKey []byte
}


type VehicleClaims struct {
	UserId   string  `json:"user_id"`
	UserName string  `json:"user_name"`
	PassWord string  `json:"pass_word"`
	jwt.StandardClaims
}
//生成token
func (j *JWT) CreateToken(claims VehicleClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

/**
自定义校验，这里省却了
*/
func (j *JWT) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.SigningKey, nil
}
//解析token
func (j *JWT) ParseToken(tokenString string) (*VehicleClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &VehicleClaims{},j.keyFunc)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			}else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				return nil, ValidationErrorUnverifiable
			}else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, ValidationErrorSignatureInvalid
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*VehicleClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &VehicleClaims{},j.keyFunc)

	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*VehicleClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = ExpiresAt
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}












