package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sensitive_words_check/constant"
	"sensitive_words_check/rpc/sensitive_word_rpc"
	"sensitive_words_check/rpc/sensitive_word_rpc/sensitive_word_rpc_service"
)

func NewRpcServer() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", constant.RpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	// 注册rpc服务
	sensitive_word_rpc.RegisterSensitiveWordServer(grpcServer, sensitive_word_rpc_service.NewService())
	reflection.Register(grpcServer)

	log.Printf("server listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
