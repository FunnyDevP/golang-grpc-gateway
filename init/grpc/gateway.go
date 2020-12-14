package grpc

import (
	"context"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
)

func RunGateway(ctx context.Context, grpcPort, restPort string) error {
	mux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	if err := proto.RegisterTodoListServiceHandlerFromEndpoint(ctx, mux, ":"+grpcPort, opt); err != nil {
		return err
	}

	zap.S().Infof("running grpc-gateway(RESTapi) on port %v...", restPort)
	return http.ListenAndServe(":"+restPort, mux)
}
