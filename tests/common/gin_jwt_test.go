/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: gin_jwt_test.go
 */

package common

import (
	"fmt"
	"testing"

	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/pkg/security"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {

	conf.InitAdminConfigProperties()

	token := security.GenerateToken("hongjun500", 11)
	assert.NotEmpty(t, token)
	fmt.Println("GenerateToken token: ", token)
	time := security.GetTokenExpireTime(token)
	assert.NotEmpty(t, time)
	fmt.Println("GetTokenExpireTime time: ", time)
	expired := security.TokenIsExpired(token)
	assert.False(t, expired)
	fmt.Println("IsTokenExpired expired: ", expired)
	username, _, _ := security.GetUsernameAndUserIdFromToken(token)
	assert.Equal(t, "hongjun500", username)
	fmt.Println("GetUsernameFromToken username: ", username)
	valid := security.TokenValid(token, username)
	assert.True(t, valid)
	fmt.Println("TokenValid valid: ", valid)
	refreshToken, err := security.RefreshToken(token)
	if err != nil {
		return
	}
	assert.Empty(t, refreshToken)
	fmt.Println("RefreshToken refreshToken: ", refreshToken)
}
