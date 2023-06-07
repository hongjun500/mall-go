package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/initialize"
	"testing"
)

var (
	TestServiceEngine *gin.Engine
)

func TestMain(m *testing.M) {
	TestServiceEngine = initialize.StartUp()
	m.Run()
}
