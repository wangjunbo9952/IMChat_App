package service

import (
	"IMChat_App/internal/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// 生成 JWT
func GenerateJWT(userID uint, userAccount string) (string, error) {
	return model.GenerateJWT(userID, userAccount)
}

// 验证 JWT
func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
	return model.ValidateJWT(tokenString)
}

// 检查用户是否有特定权限
func CheckPermissions(tokenString, requiredRole string) (bool, error) {
	// 验证 JWT
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return false, fmt.Errorf("invalid token: %v", err)
	}

	// 如果用户角色匹配，则返回 true
	if claims.Issuer == requiredRole {
		return true, nil
	}

	return false, fmt.Errorf("user does not have the required role")
}
