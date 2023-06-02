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
	// services.UmsAdminService.UmsAdminRegister(nil, &umsAdminRequest)
	requestBody, _ := json.Marshal(umsAdminRequest)
	router := initialize.StartUp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/register", bytes.NewReader(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
