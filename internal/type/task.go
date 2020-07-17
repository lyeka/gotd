package _type

import "time"

type TaskState int

const (
	TaskTODO        TaskState = 1 // 表示还没到任务的开始时间
	TasKProgressing TaskState = 2 // 表示任务处于安排时间范围内
	TaskDone        TaskState = 4 // 表示任务已经完成
	TaskExpired     TaskState = 8 // 表示任务过期前没有完成
)

// Task 代表一个待办事项
type Task struct {
	ID        int64     `json:"id"`         // 任务ID
	CID       int64     `json:"cid"`        // 创建者ID
	Title     string    `json:"title"`      // 任务标题
	Desc      string    `json:"desc"`       // 任务描述
	StartTime time.Time `json:"start_time"` // 任务开始时间
	EndTime   time.Time `json:"end_time"`   // 任务结束时间
	CreatedAt time.Time `json:"created_at"` // 任务创建时间
	UpdatedAt time.Time `json:"updated_at"` // 任务更新时间
	State     TaskState `json:"state"`      // 任务状态
}
