package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gotd/internal/auth"
	_type "github.com/lyeka/gotd/internal/type"
	"net/http"
	"time"
)

type CorRequest struct {
	TaskID    int64     `form:"taskID"`
	Title     string    `form:"title"`
	Desc      string    `form:"desc"`
	StartTime time.Time `form:"startTime" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02 15:04:05" time_utc:"1"`
}

func CreateOrUpdateTask(s *Server, rg *gin.RouterGroup) {
	rg.POST("/createOrUpdate", func(c *gin.Context) {
		resp := &Response{Code: CodeOK}
		uerID, _ := c.Get(auth.UID)
		uid := uerID.(int64)

		var req CorRequest
		if err := c.ShouldBind(&req); err != nil {
			resp.Code = CodeInvalidParam
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}

		taskID := req.TaskID

		task := _type.Task{
			ID:        taskID,
			CID:       uid,
			Title:     req.Title,
			Desc:      req.Desc,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		}
		// 创建任务
		if taskID == 0 {
			tid, err := s.DB.CreateTask(c, &task)
			if err != nil {
				resp.Code = CodeErrCreateTask
				resp.Message = "创建任务失败"
				c.JSON(http.StatusOK, resp)

			} else {
				resp.Message = "创建任务成功"
				resp.Data = tid
				c.JSON(http.StatusOK, resp)
			}
			return
		}

		// 更新任务
		err := s.DB.UpdateTask(c, &task)
		if err != nil {
			resp.Code = CodeErrUpdateTask
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
		}

		resp.Message = "更新任务成功"
		c.JSON(http.StatusOK, resp)
		return
	})
}

func GetUserTasks(s *Server, rg *gin.RouterGroup) {
	rg.GET("/myTask", func(c *gin.Context) {
		resp := Response{Code: CodeOK}
		uid, _ := c.Get(auth.UID)
		userID := uid.(int64)

		tasks, err := s.DB.GetUserTasks(c, userID)
		if err != nil {
			resp.Message = "获取用户任务列表失败"
			c.JSON(http.StatusOK, resp)
			return
		}

		resp.Data = map[string]interface{}{
			"list": tasks,
		}

		c.JSON(http.StatusOK, resp)

	})
}
