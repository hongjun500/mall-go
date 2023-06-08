package search

import (
	"github.com/hongjun500/mall-go/internal/initialize"
	"net/http"
)

func HandlerSearch() http.Handler {
	engine := initialize.StartUpSearch()
	return engine
}
