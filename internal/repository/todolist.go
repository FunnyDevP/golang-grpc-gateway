package repository

import (
	"context"
	"database/sql"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/rs/xid"
	"strconv"
	"time"
)

type TodoListRepository interface {
	Create(ctx context.Context, req *proto.CreateTaskRequest) (*proto.Task, error)
	Gets(ctx context.Context) ([]*proto.Task, error)
}

type todoListRepository struct {
	db *sql.DB
}

func NewTodoListRepository(db *sql.DB) *todoListRepository {
	return &todoListRepository{db: db}
}

func queryInsertBuilder(req *proto.CreateTaskRequest) string {
	q1 := "INSERT INTO tasks(id,name,status,created_date,cate_id"
	val := "VALUES($1,$2,$3,$4,$5"

	if req.TaskDescription != nil {
		q1 += ",description)"
		val += ",$6)"
	} else {
		q1 += ")"
		val += ")"
	}

	q1 = q1 + " " + val

	return q1
}

func (r todoListRepository) Create(ctx context.Context, req *proto.CreateTaskRequest) (*proto.Task, error) {
	query := queryInsertBuilder(req)
	if req.TaskCreated == nil {
		req.TaskCreated = &wrappers.StringValue{Value: time.Now().Format(time.RFC3339)}
	}

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	taskID := xid.NewWithTime(time.Now()).String()
	arg := []interface{}{taskID, req.TaskName, strconv.Itoa(int(proto.TaskStatus_value[req.TaskStatus.String()])), req.TaskCreated.Value, req.CategoryId}
	if req.TaskDescription != nil {
		arg = append(arg, req.TaskDescription.Value)
	}

	_, err = stmt.Exec(arg...)
	if err != nil {
		return nil, err
	}

	return &proto.Task{
		TaskId:          taskID,
		CategoryId:      req.CategoryId,
		TaskName:        req.TaskName,
		TaskDescription: req.TaskDescription,
		TaskStatus:      req.TaskStatus,
		TaskCreated:     req.TaskCreated.Value,
	}, nil
}

func (r todoListRepository) Gets(ctx context.Context) ([]*proto.Task, error) {
	return nil, nil
}
