package repository

import (
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestQueryBuilder_DescriptionIsEmpty_ReturnQueryInsertWithoutDescription(t *testing.T) {
	reqCreateSource := &proto.CreateTaskRequest{
		CategoryId:  "cate_id",
		TaskName:    "task_name",
		TaskStatus:  0,
		TaskCreated: &wrappers.StringValue{Value: time.Now().Format(time.RFC3339)},
	}
	expectedQuery := "INSERT INTO tasks(id,name,status,created_date,cate_id) VALUES($1,$2,$3,$4,$5)"
	assert.Equal(t, expectedQuery, queryInsertBuilder(reqCreateSource))
}

func TestQueryBuilder_DescriptionIsNotEmpty_ReturnQueryInsertWithDescription(t *testing.T) {
	reqCreateSource := &proto.CreateTaskRequest{
		CategoryId:      "cate_id",
		TaskName:        "task_name",
		TaskDescription: &wrappers.StringValue{Value: "task_description"},
		TaskStatus:      0,
		TaskCreated:     &wrappers.StringValue{Value: time.Now().Format(time.RFC3339)},
	}
	expectedQuery := "INSERT INTO tasks(id,name,status,created_date,cate_id,description) VALUES($1,$2,$3,$4,$5,$6)"
	assert.Equal(t, expectedQuery, queryInsertBuilder(reqCreateSource))
}
