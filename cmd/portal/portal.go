package portal

import (
	"github.com/hongjun500/mall-go/internal/initialize"
	"net/http"
)

func HandlerPortal() http.Handler {
	engine := initialize.StartUpPortal()
	return engine
}
