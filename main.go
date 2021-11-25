package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sensitive_words_check/config"
	"sensitive_words_check/constant"
	router2 "sensitive_words_check/http/router"
	"sensitive_words_check/model"
	"sensitive_words_check/rpc"
	"sensitive_words_check/setup"
	"time"
)


func init() {
	env := os.Getenv("APP_RUN_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	fmt.Println(env)
	config.InitConfig(env)
	model.Setup()
	setup.CheckTable()
}
// @title Words Check API DOC
// @version 1.0

// @contact.name cixn
// @contact.url https://blog.zhangziwen.cn
// @contact.email zeco11320@163.com
func main() {
	go func() {
		rpc.NewRpcServer()
	}()

	router := router2.NewRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", constant.HttpServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server start err: %s", err.Error())
	}
}
