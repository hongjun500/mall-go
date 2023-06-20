// @author hongjun500
// @date 2023/6/20 14:08
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package database

import (
	"context"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hongjun500/mall-go/internal/database"
)

var esTypedClient *elasticsearch.TypedClient
var ctx = context.Background()

func TestMain(m *testing.M) {
	// 获取 es 连接
	esTypedClient, _ = database.NewEsTypedClient()
	m.Run()
}
