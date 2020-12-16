package integration_test_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/internal/repository"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

var port = "5433"

var queryCreateEnumTaskStatus = "create type taskstatus as enum ('0', '1', '2');"
var queryCate = "CREATE TABLE tasks_categories(" +
	"id varchar(40) primary key," +
	"name varchar(50) not null," +
	"description text," +
	"created_date timestamp not null," +
	"deleted_date timestamp" +
	")"
var queryTask = "CREATE TABLE tasks(" +
	"id varchar(40) constraint tasks_pk primary key," +
	"name text not null," +
	"description text," +
	"status taskstatus not null," +
	"created_date timestamp not null," +
	"deleted_date timestamp," +
	"cate_id varchar(40) not null constraint tasks_tasks_categories_id_fk references tasks_categories" +
	")"

var repo repository.TodoListRepository
var req *proto.CreateTaskRequest

func migrationSchema(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func seedDataCategory(db *sql.DB) (string, error) {
	cateID := xid.NewWithTime(time.Now()).String()
	stmt, err := db.Prepare("INSERT INTO tasks_categories(id, name, description,created_date) VALUES($1, $2, $3,$4)")
	if err != nil {
		return "",err
	}

	_, err = stmt.Exec(cateID,"test_cate_name","test_cate_description",time.Now().Format(time.RFC3339))
	if err != nil {
		return "",err
	}
	return cateID,nil
}

func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	// option for process pull image
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13.1-alpine",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	var db *sql.DB
	if err != nil {
		log.Fatalf("Cloud not start resource: %v", err)
	}
	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		"localhost", 5433, "postgres", "secret", "postgres")
	if err := pool.Retry(func() error {
		db, err = sql.Open("postgres", dataSourceName)
		repo = repository.NewTodoListRepository(db)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	if err := migrationSchema(db, queryCreateEnumTaskStatus); err != nil {
		log.Fatalf("cannot migrate schema enum type: %v",err)
	}

	if err := migrationSchema(db, queryCate); err != nil {
		log.Fatalf("cannot migrate schema task_categories: %v",err)
	}

	if err := migrationSchema(db, queryTask); err != nil {
		log.Fatalf("cannot migrate schema task: %v",err)
	}

	cateID, err := seedDataCategory(db)
	if err != nil {
		log.Fatalf("cannot seed data tasks categories: %v",err)
	}

	req = &proto.CreateTaskRequest{
		CategoryId:      cateID,
		TaskName:        "test_create_tasks_name",
		TaskStatus:      0,
	}

	code := m.Run()

	// removes a container and linked volumes from docker.
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %v",err)
	}

	os.Exit(code)
}

func TestCreateTasks(t *testing.T) {
	result, err := repo.Create(context.Background(),req)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test_create_tasks_name",result.TaskName)
	assert.Equal(t, "PENDING",result.TaskStatus.String())
}
