package validation

import (
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EmptyFieldCreateSource(req *proto.CreateTaskRequest) error {
	if req.CategoryId == "" {
		return status.New(codes.InvalidArgument, "category id is required").Err()
	}
	if req.TaskName == "" {
		return status.New(codes.InvalidArgument, "task name is required").Err()
	}
	return nil
}
