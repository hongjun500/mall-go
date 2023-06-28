//	@author	hongjun500
//	@date	2023/6/19 18:19
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

type Es struct {
	Cli *elasticsearch.Client
	// 适用于 go api 的 elasticsearch 连接
	TypedCli *elasticsearch.TypedClient
}

// NewEsTypedClient 初始化 适用于 go api 的 elasticsearch 连接
func NewEsTypedClient() (*Es, error) {
	es := new(Es)
	var typedClient *elasticsearch.TypedClient
	var client *elasticsearch.Client
	var err error
	once := sync.Once{}

	// ca证书
	cert, _ := os.ReadFile("D:\\elasticsearch-8.7.0\\config\\certs\\http_ca.crt")

	config := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
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
		typedClient, err = elasticsearch.NewTypedClient(config)
		client, err = elasticsearch.NewClient(config)
		if err != nil {
			log.Fatal("NewEsClient Fail, ERR = ", err)
		}
		ping := typedClient.Ping()
		success, err := ping.IsSuccess(context.Background())
		if !success {
			log.Fatalln("Elasticsearch Connected Fail, ERR = ", err.Error())
		}
	})
	es.TypedCli = typedClient
	es.Cli = client
	log.Println("Elasticsearch Connected!")
	return es, nil
}
