/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 使用 golang-jwt 实现 JWT 认证,并将其封装成 gin 的中间件
 */

package mid

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	sub     = "sub"
	created = "created"
	// TokenHeader todo 从配置中获取
	TokenHeader = "Authorization"
	// TokenHead todo 从配置中获取
	TokenHead = "Bearer "
)

type CustomClaims struct {
	Sub     string    `json:"sub"`
	Created time.Time `json:"created"`
	jwt.RegisteredClaims
}

// GinJWTMiddleware 自定义使用 golang-jwt 实现 JWT 认证的 gin 中间件
func GinJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := http.StatusOK
		header := c.GetHeader(TokenHeader)
		if header != "" && strings.HasPrefix(header, TokenHead) {
			authToken := header[len(TokenHead):]
			username, err := GetUsernameFromToken(authToken)
			if err != nil {
				log.Printf("get username from token[%v] is fail %d\n", authToken, err)
				status = gin_common.CodeInvalidToken
			}
			log.Printf("checking username: %v", username)
			if username == "" {
				status = gin_common.CodeInvalidToken
			}
			valid := TokenValid(authToken, username)
			if !valid {
				status = gin_common.CodeInvalidToken
			}
			log.Printf("authenticated user: %v", username)
			c.Set("username", username)
		} else {
			status = gin_common.Unauthorized
		}
		if status != http.StatusOK {
			gin_common.CreateFail(status, c)
			// c.AbortWithStatus(status)
			c.Abort()
			return
		}
		// 请求前
		c.Next()
		// 请求后
		// 暂无
	}
}

// GenerateTokenExpire 生成 token 的过期时间
func GenerateTokenExpire() time.Time {
	// todo 从配置中获取
	return time.Now().Add(time.Hour * 24 * 7)
}

// generatePrivateKey 生成 token 的私钥
func generatePrivateKey() []byte {
	// todo 从配置中获取
	return []byte("mall-go-jwt-secret")
}

// GenerateTokenFromClaims 根据自定义声明生成 token
func GenerateTokenFromClaims(claimsMap map[string]any) string {
	claims := CustomClaims{
		claimsMap[sub].(string),
		claimsMap[created].(time.Time),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(GenerateTokenExpire()),
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
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return false
	}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return generatePrivateKey(), nil
	})
	if err != nil {
		return false
	}
	return token.Valid && claims.Sub == username
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
	if length := len(oldTokenString[7:]); length == 0 {
		return "", nil
	}
	// 验证 token 是否有效
	claims, err := GetClaimsFromToken(oldTokenString)
	if err != nil {
		return "", err
	}
	if &claims == nil {
		return "", nil
	}
	// 验证 token 是否过期,过期不支持刷新
	if TokenIsExpired(oldTokenString) {
		return "", nil
	}
	// 验证 token 是否在指定时间内刷新过,30 分钟内不支持刷新
	if TokenRefreshJustBeforeExpired(oldTokenString, time.Minute*30) {
		return "", nil
	}
	return "", nil
}

// TokenRefreshJustBeforeExpired 判断 token 是否在指定时间内刷新过
func TokenRefreshJustBeforeExpired(tokenString string, tm time.Duration) bool {
	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return false
	}
	created := claims.Created
	refreshTime := time.Now()
	if refreshTime.After(created) && refreshTime.Before(created.Add(tm)) {
		return true
	}
	return false
}
