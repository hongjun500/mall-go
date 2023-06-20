// @author hongjun500
// @date 2023/6/19 18:19
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

// NewEsTypedClient 初始化 适用于 go api 的 es 连接
func NewEsTypedClient() (*elasticsearch.TypedClient, error) {
	var typedClient *elasticsearch.TypedClient
	var err error
	once := sync.Once{}

	// ca证书
	cert, _ := os.ReadFile("D:\\elasticsearch-8.8.0\\config\\certs\\http_ca.crt")

	config := elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Username:  "elastic",
		Password:  "elastic",
		// CloudID:                "",
		// APIKey:                 "",
		// ServiceToken:           "",
		// CertificateFingerprint: "",
		// Header:                 nil,
		CACert: cert,
	}

	once.Do(func() {
		// client, err = elasticsearch.NewClient(config)
		typedClient, err = elasticsearch.NewTypedClient(config)
		if err != nil {
			log.Fatal("NewEsClient Fail, ERR = ", err)
		}
		ping := typedClient.Ping()
		success, _ := ping.IsSuccess(context.Background())
		if !success {
			log.Fatalln("Elasticsearch Connected Fail, ERR = ", ping)
		}
	})
	log.Println("Elasticsearch Connected!")
	return typedClient, nil
}
