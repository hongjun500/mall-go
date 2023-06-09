/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 使用 golang-jwt 实现 封装一些函数
 */

package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hongjun500/mall-go/internal/conf"
	"time"
)

const (
	sub     = "sub"
	created = "created"
	expired = "expired"
)

type CustomClaims struct {
	Sub     string    `json:"sub"`
	Created time.Time `json:"created"`
	jwt.RegisteredClaims
}

// GenerateTokenExpire 生成 token 的过期时间
func GenerateTokenExpire() time.Time {
	// 604800 秒 = 7 天
	return time.Now().Add(time.Duration(conf.GlobalJwtConfigProperties.Expiration) * time.Second)
}

// generatePrivateKey 生成 token 的私钥
func generatePrivateKey() []byte {
	return []byte(conf.GlobalJwtConfigProperties.Secret)
}

// GenerateTokenFromClaims 根据自定义声明生成 token
func GenerateTokenFromClaims(claimsMap map[string]any) string {
	expiredTime := GenerateTokenExpire()
	if claimsMap[expired] != nil {
		expiredTime = claimsMap[expired].(time.Time)
	}
	claims := CustomClaims{
		claimsMap[sub].(string),
		claimsMap[created].(time.Time),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	// 用这个算法实现的签名方法直接可以利用一个字符串，无需私钥或者公钥的麻烦操作
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, _ := token.SignedString(generatePrivateKey())
	return tokenString
}

// GenerateToken 根据用户名生成 token
func GenerateToken(username string) string {
	claimsMap := make(map[string]any)
	claimsMap[sub] = username
	claimsMap[created] = time.Now()
	return GenerateTokenFromClaims(claimsMap)
}

// GetClaimsFromToken 从 token 中获取 claims
func GetClaimsFromToken(tokenString string) (CustomClaims, error) {
	claims := CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return generatePrivateKey(), nil
	})
	if err != nil {
		return claims, err
	}
	if token.Valid {
		return claims, nil
	}
	return claims, err
}

// GetUsernameFromToken 从 token 中获取 username
func GetUsernameFromToken(tokenString string) (string, error) {
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Sub, nil
}

// TokenIsExpired 判断 token 是否过期
func TokenIsExpired(tokenString string) bool {
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return true
	}
	return claims.ExpiresAt.Time.Before(time.Now())
}

// TokenValid 判断 token 是否有效
func TokenValid(tokenString string, username string) bool {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return generatePrivateKey(), nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid && claims.Sub == username {
		return true
	}
	return false
}

// GetTokenExpireTime 获取 token 的过期时间
func GetTokenExpireTime(tokenString string) time.Time {
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return time.Time{}
	}
	return claims.ExpiresAt.Time
}

// RefreshToken 刷新 token
func RefreshToken(oldTokenString string) (string, error) {
	if oldTokenString == "" {
		return "", nil
	}
	token := oldTokenString[7:]
	if length := len(token); length == 0 {
		return "", nil
	}
	// 验证 token 是否有效
	claims, err := GetClaimsFromToken(token)
	if err != nil {
		return "", err
	}
	if &claims == nil {
		return "", nil
	}
	// 验证 token 是否过期,过期不支持刷新
	if TokenIsExpired(token) {
		return "", nil
	}
	// 验证 token 是否在指定时间内刷新过,30 分钟内不支持刷新
	if TokenRefreshJustBeforeExpired(token, time.Minute*30) {
		return token, nil
	} else {
		// 重新生成token
		// todo 这里会有一个问题：原先的 token 只要在有效期内仍然可以使用
		newToken := GenerateToken(claims.Sub)
		return newToken, nil
	}
}

// TokenRefreshJustBeforeExpired 判断 token 是否在指定时间内刷新过
func TokenRefreshJustBeforeExpired(tokenString string, tm time.Duration) bool {
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return false
	}
	ct := claims.Created
	refreshTime := time.Now()
	if refreshTime.After(ct) && refreshTime.Before(ct.Add(tm)) {
		return true
	}
	return false
}
