package main

import (
	"fmt"
	"log"
	"modern-blog/configs"
	"modern-blog/internal/dao"
	"modern-blog/internal/routers"
	"modern-blog/internal/util/important"
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
