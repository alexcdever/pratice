package main

import (
	"blog/configs"
	"blog/internal/dao"
	"blog/internal/routers"
	"blog/internal/util/important"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	engine := routers.NewRouter()
	log.Printf("config: %v", configs.Con)

	dao.InitDb(configs.Con)
	important.ImportPosts(configs.Con)

	listenPort := fmt.Sprintf(":%v", configs.Con.Port)
	s := &http.Server{
		Addr:           listenPort,
		Handler:        engine,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
