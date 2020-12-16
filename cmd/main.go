package main

import (
	"fmt"
	"github.com/FunnyDevP/todolist-golang-grpc-gateway/api/proto"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()

	zap.ReplaceGlobals(logger)
}

func main() {
	//2006-01-02T15:04:05Z07:00

	//layout := "2006-01-02T15:04:05Z07:00"
	//
	//t, err := time.Parse(layout,value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////time.Now().Format(time.RFC3339)
	//fmt.Println(t.String())


	//dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
	//	"localhost",5432,"postgres","mysecret","postgres")
	//db, _ := sql.Open("postgres",dataSourceName)
	//
	//if err := db.Ping(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//stmt, err := db.Prepare("INSERT INTO tasks_categories(id, name, description,created_date) VALUES($1, $2, $3,$4)")
	//if err != nil {
	//	log.Fatalf("cannot prepare stmt: %v",err)
	//}
	//value := "2020-12-14T23:52:00+07:00"
	//t,_:= time.Parse(time.RFC3339,value)
	//r, err := stmt.Exec(xid.NewWithTime(time.Now()).String(),"test","",t)
	//if err != nil {
	//	log.Fatalf("cannot exec : %v\n",err)
	//}
	//num, _ := r.RowsAffected()
	req := &proto.CreateTaskRequest{
		CategoryId:      "cateID",
		TaskName:        "test_create_tasks_name",
		TaskStatus:      1,
	}

	fmt.Println(proto.TaskStatus_value[req.TaskStatus.String()])
	//grpcPort := "50001"
	//restPort := "8080"
	//go func() {
	//	_ = grpc.RunGateway(context.Background(),grpcPort,restPort)
	//}()
	//
	//if err := grpc.RunServer(grpcPort); err != nil {
	//	zap.S().Errorw("cannot run grpc server", zap.String("error", err.Error()))
	//}
}