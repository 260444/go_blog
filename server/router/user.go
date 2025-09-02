package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	userRouter := router.Group("user")
	userPublicRoute := PublicRouter.Group("user")
	userLoginRouter := PublicRouter.Group("user").Use(middleware.LoginRecord())
	userAdminRoute := AdminRouter.Group("user")
	userApi := api.ApiGroupApp.UserApi
	{
		userRouter.POST("logout", userApi.Logout)
		userRouter.PUT("resetPassword", userApi.UserResetPassword)
		userRouter.GET("info", userApi.UserInfo)
		userRouter.PUT("changeInfo", userApi.UserChangeInfo)
		userRouter.GET("weather", userApi.UserWeather)
		userRouter.GET("chart", userApi.UserChart)
	}
	userRouter.POST("logout", userApi.Logout)
	userRouter.PUT("resetPassword", userApi.UserResetPassword)
	userRouter.GET("info", userApi.UserInfo)
	userRouter.PUT("changeInfo", userApi.UserChangeInfo)
	userRouter.GET("weather", userApi.UserWeather)
	userRouter.GET("chart", userApi.UserChart)
	{
		userPublicRoute.GET("list", userApi.UserList)
		userLoginRouter.GET("list", userApi.UserList)
		userAdminRoute.GET("list", userApi.UserList)
	}
}
