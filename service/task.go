package service

import (
	"Todo/pkg/ctl"
	"Todo/pkg/e"
	"Todo/pkg/util"
	"Todo/repository/db/dao"
	"Todo/repository/db/model"
	"Todo/types"
	"context"
	"sync"
	"time"
)

type TaskService struct {
}

var TaskServiceOnce sync.Once
var TaskServiceIns *TaskService

func GetTaskService() *TaskService {
	TaskServiceOnce.Do(func() {
		TaskServiceIns = &TaskService{}
	})
	return TaskServiceIns
}

func (s *TaskService) CreateTask(ctx context.Context, req *types.TaskCreateReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	task := &model.Task{
		Uid:       u.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	taskDao := dao.NewTaskDao(ctx)
	err = taskDao.CreateTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskCreate
		return ctl.RespError(err, code), err
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *types.TaskUpdateReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	task, err := taskDao.FindTaskById(req.Id, u.ID)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Content != "" {
		task.Content = req.Content
	}
	if req.Status == 0 {
		task.Status = 0
		task.EndTime = ""
	} else if req.Status == 1 {
		task.Status = 1
		task.EndTime = time.Now().Format("2006-01-02 15:04:05")
	}
	err = taskDao.Update(task)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskUpdate
		return ctl.RespError(err, code), err
	}

	return ctl.RespSuccess(), nil
}

func (s *TaskService) ShowTask(ctx context.Context, req *types.TaskShowReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	task, err := taskDao.FindTaskById(req.Id, u.ID)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	taskInfoResp := &types.TaskInfoResp{
		Id:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(),
		Status:    task.Status,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
	task.AddView()
	return ctl.RespSuccessWithData(taskInfoResp), nil
}

func (s *TaskService) ListTasks(ctx context.Context, req *types.TaskListReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	tasks, count, err := taskDao.ListTasks(u.ID, req.Limit, req.Start)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	resps := make([]*types.TaskInfoResp, 0)
	for _, v := range tasks {
		resps = append(resps, &types.TaskInfoResp{
			Id:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(resps, count), nil
}

func (s *TaskService) SearchByInfo(ctx context.Context, req *types.TaskSearchReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	tasks, count, err := taskDao.SearchTasksByInfo(u.ID, req.Info, req.Start)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	resps := make([]*types.TaskInfoResp, 0)
	for _, v := range tasks {
		resps = append(resps, &types.TaskInfoResp{
			Id:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(resps, count), nil
}

func (s *TaskService) SearchByStatus(ctx context.Context, req *types.TaskSearchReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	tasks, count, err := taskDao.SearchTasksByStatus(u.ID, req.Status, req.Start)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	resps := make([]*types.TaskInfoResp, 0)
	for _, v := range tasks {
		resps = append(resps, &types.TaskInfoResp{
			Id:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(resps, count), nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *types.TaskDeleteReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	task, err := taskDao.FindTaskById(req.Id, u.ID)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	err = taskDao.Delete(task)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorDatabase
		return ctl.RespError(err, code), err
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskService) DeleteTasksWithStatus(ctx context.Context, req *types.TaskDeleteReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	taskDao := dao.NewTaskDao(ctx)
	err = taskDao.DeleteWithStatus(u.ID, req.Status)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorTaskNotExists
		return ctl.RespError(err, code), err
	}
	return ctl.RespSuccess(), nil
}
