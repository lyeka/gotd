package mongo

import (
	"context"
	"errors"
	_type "github.com/lyeka/gotd/internal/type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserTasks struct {
	Uid   int64
	Tasks []Task
}

type Task struct {
	ID        int64     // 任务ID
	Title     string    // 任务标题
	Desc      string    // 任务描述
	StartTime time.Time `bson:"start_time"` // 任务开始时间
	EndTime   time.Time `bson:"end_time"`   // 任务结束时间
	CreatedAt time.Time `bson:"created_at"` // 任务创建时间
	UpdatedAt time.Time `bson:"updated_at"` // 任务更新时间
	State     int       // 任务状态
}

// Out 转化为外部任务结构
func (t *Task) Out(uid int64) _type.Task {
	return _type.Task{
		ID:        t.ID,
		CID:       uid,
		Title:     t.Title,
		Desc:      t.Desc,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		State:     _type.TaskState(t.State),
	}
}

// GetUserTask 获取用户单个任务
// todo 换成mongo映射字段来获取单个任务
func (db *DB) GetUserTask(ctx context.Context, uid, tid int64) (task *Task, err error) {
	userTasks, exist, err := db.getUserTasks(ctx, uid)
	if err != nil {
		return nil, err
	}

	if !exist {
		err = errors.New("任务不存在")
		return
	}

	for _, t := range userTasks.Tasks {
		if t.ID == tid {
			return &t, nil
		}
	}

	return nil, errors.New("任务不存在")
}

func (db *DB) CreateTask(ctx context.Context, task *_type.Task) (id int64, err error) {
	_, exist, err := db.getUserTasks(ctx, task.CID)
	if err != nil {
		return
	}

	if !exist {
		err = db.createUserTask(ctx, task.CID)
		if err != nil {
			return
		}
	}

	now := time.Now()
	id = db.GenerateID()
	innerTask := Task{
		ID:        id,
		Title:     task.Title,
		Desc:      task.Desc,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
		CreatedAt: now,
		UpdatedAt: now,
		State:     int(task.State),
	}

	_, err = db.CollTask().UpdateOne(ctx,
		bson.M{"uid": task.CID},
		bson.M{"$push": bson.M{
			"tasks": innerTask,
		}},
	)
	if err != nil {
		return
	}

	return
}

func (db *DB) UpdateTask(ctx context.Context, task *_type.Task) (err error) {
	task.UpdatedAt = time.Now()
	res, err := db.CollTask().UpdateOne(ctx,
		bson.M{"uid": task.CID, "tasks": bson.M{"$elemMatch": bson.M{"id": task.ID}}},
		bson.M{"$set": bson.M{"tasks.$": task}},
	)

	if err != nil {
		return
	}

	if res.ModifiedCount != 1 {
		err = errors.New("wrong update task count")
	}

	return
}

func (db *DB) GetUserTasks(ctx context.Context, uid int64) ([]_type.Task, error) {
	userTasks, exist, err := db.getUserTasks(ctx, uid)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, err
	}

	tasks := []_type.Task{}
	for _, t := range userTasks.Tasks {
		tasks = append(tasks, t.Out(userTasks.Uid))
	}
	return tasks, nil
}

func (db *DB) getUserTasks(ctx context.Context, uid int64) (userTasks *UserTasks, exist bool, err error) {
	result := db.CollTask().FindOne(ctx, bson.M{"uid": uid})
	err = result.Err()

	if err == mongo.ErrNoDocuments {
		err = nil
		return
	}

	if err != nil {
		return
	}

	exist = true
	userTasks = new(UserTasks)
	err = result.Decode(userTasks)
	if err != nil {
		return
	}
	return
}

func (db *DB) createUserTask(ctx context.Context, uid int64) error {
	userTasks := new(UserTasks)
	userTasks.Uid = uid
	userTasks.Tasks = []Task{}

	_, err := db.CollTask().InsertOne(ctx, userTasks)
	return err
}
