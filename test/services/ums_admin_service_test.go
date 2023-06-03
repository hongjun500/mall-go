package services

import (
	"bytes"
	"encoding/json"
	"github.com/hongjun500/mall-go/internal/initialize"
	"github.com/hongjun500/mall-go/internal/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*func TestMain(t *testing.M) {
	// 在测试之前，初始化数据库连接
	// initialize.StartUp()

}*/

func TestHashPassword(t *testing.T) {
	password, err := services.HashPassword("123456")
	assert.NoError(t, err, "HashPassword should not return an error")
	assert.NotEmpty(t, password, "Hashed password should not be empty")
}

func TestVerifyPassword(t *testing.T) {
	password := "123456"
	hashPassword, _ := services.HashPassword(password)
	verify := services.VerifyPassword(password, hashPassword)
	_ = "$2a$10$cDLM3NGJJgfBfQsfpjNSGeK5xImWfs8W5SrS709L.eZYV6qZRAy1e"
	assert.Equal(t, true, verify, "verify success")
}

func TestUmsAdminRegister(t *testing.T) {

	var umsAdminRequest services.UmsAdminRequest
	umsAdminRequest.Username = "hongjun"
	umsAdminRequest.Password = "123456"
	umsAdminRequest.Icon = "http://www.baidu.com"
	umsAdminRequest.Email = "emial.com"
	umsAdminRequest.Nickname = "nickname"
	umsAdminRequest.Note = "note"
	body, err := json.Marshal(umsAdminRequest)
	if err != nil {
		return
	}

	router := initialize.StartUp()

	w := httptest.NewRecorder()
	// req, err := http.Post("/admin/register", "application/json", bytes.NewReader(body))
	req, _ := http.NewRequest("POST", "/admin/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())

}
