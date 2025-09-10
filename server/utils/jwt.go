package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT密钥（应该从配置中读取，并且保持安全）
var jwtSecret = []byte("your-secret-key")

// Claims 定义token中包含的数据
type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID int64, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour) // 过期时间7天

	claims := Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "your-app-name",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetUserIDFromToken 从token中获取用户ID
func GetUserIDFromToken(token string) (int64, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}