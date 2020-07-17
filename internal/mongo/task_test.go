package mongo

import (
	"context"
	"errors"
	"fmt"
	_type "github.com/lyeka/gotd/internal/type"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestTaskCURD(t *testing.T) {
	db := getTestDB()
	ctx := context.Background()
	uid := int64(1)

	task1 := _type.Task{
		CID:   uid,
		Title: "title1",
	}
	task2 := _type.Task{
		CID:   uid,
		Title: "title2",
	}

	// 清理数据，保证下次运行测试用例不会出错
	_, err := db.CollTask().DeleteMany(ctx, bson.M{})
	if err != nil {
		t.Fatal("clear test db failed: ", err)
	}

	// 验证插入
	taskID1, err := checkTaskInsert(db, ctx, &task1, 1, &task1)
	if err != nil {
		t.Fatal(err)
	}
	task1.ID = taskID1

	taskID2, err := checkTaskInsert(db, ctx, &task2, 2, &task2)
	if err != nil {
		t.Fatal(err)
	}
	task2.ID = taskID2

	task1.Title = "title1 updated"
	task1.Desc = "desc1"
	err = db.UpdateTask(ctx, &task1)
	if err != nil {
		t.Fatal("update task failed: ", err)
	}
	err = checkTaskUpdate(db, ctx, &task1, &task1)
	if err != nil {
		t.Fatal(err)
	}

}

// 验证一个任务有没有被正确创建
func checkTaskInsert(db *DB, ctx context.Context, task *_type.Task, expectedTaskCount int, expectedTask *_type.Task) (id int64, err error) {
	id, err = db.CreateTask(ctx, task)
	if err != nil {
		err = errors.New(fmt.Sprintf("create task failed: %v", err))
	}
	fmt.Println("task id: ", id)

	userTasks, exist, err := db.getUserTasks(ctx, task.CID)
	if err != nil || !exist {
		err = errors.New(fmt.Sprintf("get user tasks failed: %v", err))
		return
	}

	if len(userTasks.Tasks) != expectedTaskCount {
		err = errors.New(fmt.Sprintf("wrong user task count: %v", userTasks.Tasks))
	}

	exist = false
	for _, t := range userTasks.Tasks {
		if t.Title == expectedTask.Title && t.ID == id {
			exist = true
			break
		}
	}

	if !exist {
		err = errors.New("do not insert task successfully")
	}
	return
}

// 检测用户
func checkTaskUpdate(db *DB, ctx context.Context, updatedTask, expectedTask *_type.Task) error {
	innerTask, err := db.GetUserTask(ctx, updatedTask.CID, updatedTask.ID)
	if err != nil {
		return err
	}

	if innerTask.Title != expectedTask.Title || innerTask.Desc != expectedTask.Desc {
		return errors.New("update task content failed")
	}
	return nil
}
