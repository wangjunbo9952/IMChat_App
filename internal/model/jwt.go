package model

import (
	"IMChat_App/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// 生成一个新的 JWT
func GenerateJWT(userID uint, userAccount string) (string, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(1 * time.Hour) // 设置过期时间为1小时
	sub := strconv.Itoa(int(userID))

	// 创建一个自定义的 Claims
	claims := &jwt.StandardClaims{
		Subject:   sub,
		Audience:  userAccount,
		Issuer:    "wangjunbo9952",       // 设置 JWT 的发行者
		ExpiresAt: expirationTime.Unix(), // 设置过期时间（Unix时间戳）
		IssuedAt:  time.Now().Unix(),     // 设置签发时间（Unix时间戳）
		// Id:        "time.Now().Unix()",   // 设置唯一的ID防止重放攻击,还没想好怎么实现 TODO
	}

	// 创建 JWT Token，传入声明信息
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并生成 token 字符串
	signedToken, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return signedToken, nil
}

// 验证JWT令牌的有效性
func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
	fmt.Println(config.JwtSecret)
	// 解析并验证Token
	var res jwt.StandardClaims
	token, err := jwt.ParseWithClaims(tokenString, &res, func(token *jwt.Token) (interface{}, error) {
		// 断言进行类型检查，确保Token使用的是我们预期的签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回签名密钥
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		// 处理解析错误，如Token过期等
		return nil, err
	}

	// 检查Token的有效性
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &res, nil
}
