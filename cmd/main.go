package main

import (
	"context"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/init/grpc"
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()

	zap.ReplaceGlobals(logger)
}

func main() {
	grpcPort := "50001"
	restPort := "8080"
	go func() {
		_ = grpc.RunGateway(context.Background(),grpcPort,restPort)
	}()

	if err := grpc.RunServer(grpcPort); err != nil {
		zap.S().Errorw("cannot run grpc server", zap.String("error", err.Error()))
	}
}