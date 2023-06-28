package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/hongjun500/mall-go/docs/mall_search"
	// _ "github.com/hongjun500/mall-go/docs/mall_admin"
	"github.com/hongjun500/mall-go/cmd/admin"
	"github.com/hongjun500/mall-go/cmd/portal"
	"github.com/hongjun500/mall-go/cmd/search"
	"github.com/hongjun500/mall-go/internal/conf"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	// 初始化全部配置
	// conf.InitAllConfigProperties()
	// 修改于 2021.6.15 13:30 将上面的函数名修改为 init , 以便于在 main 函数之前执行并无需手动调用

	adminServer := &http.Server{
		Addr:        conf.GlobalAdminServerConfigProperties.Host + ":" + conf.GlobalAdminServerConfigProperties.Port,
		Handler:     admin.HandlerAdmin(),
		ReadTimeout: time.Duration(conf.GlobalAdminServerConfigProperties.ReadTimeout) * time.Second,
		// WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		err := adminServer.ListenAndServe()
		if err != nil {
			log.Printf("adminServer.ListenAndServe() error: %v\n", err)
		}
		return err
	})

	if conf.GlobalPortalServerConfigProperties.Enable {
		portalServer := &http.Server{
			Addr:        conf.GlobalPortalServerConfigProperties.Host + ":" + conf.GlobalPortalServerConfigProperties.Port,
			Handler:     portal.HandlerPortal(),
			ReadTimeout: time.Duration(conf.GlobalPortalServerConfigProperties.ReadTimeout) * time.Second,
			// WriteTimeout: 10 * time.Second,
		}

		g.Go(func() error {
			err := portalServer.ListenAndServe()
			if err != nil {
				log.Printf("portalServer.ListenAndServe() error: %v\n", err)
			}
			return err
		})
	}

	if conf.GlobalSearchServerConfigProperties.Enable {
		searchServer := &http.Server{
			Addr:        conf.GlobalSearchServerConfigProperties.Host + ":" + conf.GlobalSearchServerConfigProperties.Port,
			Handler:     search.HandlerSearch(),
			ReadTimeout: time.Duration(conf.GlobalSearchServerConfigProperties.ReadTimeout) * time.Second,
			// WriteTimeout: 10 * time.Second,
		}
		g.Go(func() error {
			err := searchServer.ListenAndServe()
			if err != nil {
				log.Printf("searchServer.ListenAndServe() error: %v\n", err)
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

/*
func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8080",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil {
			log.Printf("server01 error: %v", err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil {
			log.Printf("server02 error: %v", err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}*/
