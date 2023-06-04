/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: gin_jwt_test.go
 */

package mid

import (
	"fmt"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := mid.GenerateToken("hongjun500")
	assert.NotEmpty(t, token)
	fmt.Println("GenerateToken token: ", token)
	time := mid.GetTokenExpireTime(token)
	assert.NotEmpty(t, time)
	fmt.Println("GetTokenExpireTime time: ", time)
	expired := mid.TokenIsExpired(token)
	assert.False(t, expired)
	fmt.Println("IsTokenExpired expired: ", expired)
	username, _ := mid.GetUsernameFromToken(token)
	assert.Equal(t, "hongjun500", username)
	fmt.Println("GetUsernameFromToken username: ", username)
	valid := mid.TokenValid(token, username)
	assert.True(t, valid)
	fmt.Println("TokenValid valid: ", valid)
	refreshToken, err := mid.RefreshToken(token)
	if err != nil {
		return
	}
	assert.Empty(t, refreshToken)
	fmt.Println("RefreshToken refreshToken: ", refreshToken)
}
