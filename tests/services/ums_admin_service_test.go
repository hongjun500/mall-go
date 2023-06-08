package services

import (
	"bytes"
	"encoding/json"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
	"github.com/hongjun500/mall-go/internal/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	var umsAdminRequest ums_admin.UmsAdminRegisterRequest
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

	w := httptest.NewRecorder()
	// req, err := http.Post("/admin/register", "application/json", bytes.NewReader(body))
	req, _ := http.NewRequest("POST", "/admin/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	TestServiceEngine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestUmsAdminLogin(t *testing.T) {
	var umsAdminLogin ums_admin.UmsAdminLogin
	umsAdminLogin.Username = "hongjun500"
	umsAdminLogin.Password = "123456"
	body, err := json.Marshal(umsAdminLogin)
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	TestServiceEngine.ServeHTTP(w, req)

	s := w.Body.String()
	var ginCommonResponse gin_common.GinCommonResponse

	err = json.Unmarshal([]byte(s), &ginCommonResponse)
	assert.NoError(t, err, "json.Unmarshal should not return an error")
	assert.Equal(t, "success", ginCommonResponse.Status)
	assert.True(t, ginCommonResponse.Data != nil, "data should not be nil")
	result := ginCommonResponse.Data.(map[string]interface{})
	assert.NotEmpty(t, result["token"])
	assert.NotEmpty(t, result["tokenHead"])
}

func TestUmsAdminInfo(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/admin/info/1", bytes.NewReader([]byte("1")))
	req.Header.Set(conf.GlobalJwtConfigProperties.TokenHeader, "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJob25nanVuNTAwIiwiY3JlYXRlZCI6IjIwMjMtMDYtMDdUMTQ6NDQ6MDMuNDI4MTYzNSswODowMCIsImV4cCI6MTY4NjcyNTA0M30.1CRZPhEbRevxQSvKB5to1hRqniXmufBhcJ9r7XO8_9cYpT2TK7WK2yQLW-0Ki-2uDC8-oLGMnQPe8HYlJc04zA")
	TestServiceEngine.ServeHTTP(w, req)

	var ginCommonResponse gin_common.GinCommonResponse
	err := json.Unmarshal([]byte(w.Body.String()), &ginCommonResponse)
	assert.NoError(t, err, "json.Unmarshal should not return an error")
	assert.Equal(t, "success", ginCommonResponse.Status)

	assert.True(t, ginCommonResponse.Data != nil, "data should not be nil")
	result := ginCommonResponse.Data.(map[string]any)

	assert.Equal(t, "tests", result["username"])

}
