package route

import (
	"Todo/api"
	"Todo/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api")
	{
		user := v1.Group("user")
		user.POST("register", api.UserRegisterHandler())
		user.POST("login", api.UserLoginHandler())

		task := v1.Group("task")
		task.Use(middleware.JWT())
		{
			task.POST("create", api.CreateTaskHandler())
			task.POST("update", api.UpdateTaskHandler())
			task.GET("show", api.ShowTaskHandler())
			task.GET("list", api.ListTasksHandler())
			task.POST("search/info", api.SearchByInfoHandler())
			task.POST("search/status", api.SearchByStatusHandler())
			task.DELETE("delete", api.DeleteTaskByIdHandler())
			task.DELETE("delete/status", api.DeleteTasksByStatusHandler())
		}
	}

	return r
}
