package admin

import (
	"github.com/hongjun500/mall-go/internal/initialize"
	"net/http"
)

func HandlerAdmin() http.Handler {
	ginEngine := initialize.StartUpAdmin()
	return ginEngine
}

/*
func main() {
	fmt.Println("hello, mall-go")

	ginEngine := initialize.StartUpAdmin()

	mallAdminServer := &http.Server{
		Addr:        conf.GlobalServerConfigProperties.Host + ":" + conf.GlobalServerConfigProperties.Port,
		Handler:     ginEngine,
		ReadTimeout: time.Duration(conf.GlobalServerConfigProperties.ReadTimeout) * time.Second,
	}

	go func() {
		// 服务连接
		if err := mallAdminServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mallAdminServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
*/