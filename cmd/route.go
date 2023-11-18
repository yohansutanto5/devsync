package main

import (
	"app/handler"
	"app/service"

	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware, gin.Recovery())
	// Initiate all services and dependency
	// studentService := service.NewStudentService(database)

	userService := service.NewUserService(database)
	UserProfileService := service.NewUserProfileService(database)

	// Define The route Path
	r.POST("/user", func(c *gin.Context) { handler.InsertUser(c, userService) })
	r.POST("/userprofile", func(c *gin.Context) { handler.InsertUserProfile(c, UserProfileService) })

	r.GET("/user", func(c *gin.Context) { handler.GetUser(c, userService) })
	r.GET("/userprofile", func(c *gin.Context) { handler.GetUserProfile(c, UserProfileService) })

	// r.DELETE("/template/:id", handler.DeleteStudent)
	// r.PUT("/template/:id", handler.UpdateStudent)
	return r
}
