package validation

import (
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestEmptyFieldCreateSource_CategoryIdIsEmpty_ReturnErrorCodeInvalidArgument(t *testing.T) {
	req := &proto.CreateTaskRequest{
		TaskName:   "task_name",
		TaskStatus: 0,
	}
	err := EmptyFieldCreateSource(req)
	assert.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.InvalidArgument)
	assert.Equal(t, s.Message(), "category id is required")
}

func TestEmptyFieldCreateSource_TaskNameIsEmpty_ReturnErrorCodeInvalidArgument(t *testing.T) {
	req := &proto.CreateTaskRequest{
		CategoryId: "cate_Id",
		TaskStatus: 0,
	}
	err := EmptyFieldCreateSource(req)
	assert.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.InvalidArgument)
	assert.Equal(t, s.Message(), "task name is required")
}
