package grpc

import (
	"fmt"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/internal/handler"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func RunServer(grpcPort string) error {
	todoListHandler := handler.NewTodolistHandler()
	listen, err := net.Listen("tcp", fmt.Sprintf(":"+grpcPort))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	proto.RegisterTodoListServiceServer(server, todoListHandler)

	zap.S().Infof("running grpc server on port %v...",grpcPort)
	return server.Serve(listen)
}
