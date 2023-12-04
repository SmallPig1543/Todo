package api

import (
	"Todo/pkg/util"
	"Todo/service"
	"Todo/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskCreateReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.CreateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskUpdateReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.UpdateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func ShowTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskShowReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.ShowTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func ListTasksHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskListReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.ListTasks(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func SearchByInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskSearchReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.SearchByInfo(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func SearchByStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskSearchReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.SearchByStatus(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func DeleteTaskByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskDeleteReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

func DeleteTasksByStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskDeleteReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskService()
			resp, err := l.DeleteTasksWithStatus(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}
