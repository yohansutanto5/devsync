package main

import (
	"app/Integration"
	"app/handler"
	"app/service"

	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware, gin.Recovery())
	// Initiate all services and dependency
	IntegrationService := Integration.NewExternalService(&configuration)
	// studentService := service.NewStudentService(database)
	userService := service.NewUserService(database)
	UserProfileService := service.NewUserProfileService(database)
	ApplicationService := service.NewApplicationService(database)
	ReleaseOPSService := service.NewReleaseOPSService(database, IntegrationService)
	// Define The route Path
	// System API
	r.GET("/health", func(c *gin.Context) { handler.GetSystemHealth(c, database) })

	// Post Method
	r.POST("/user", func(c *gin.Context) { handler.InsertUser(c, userService) })
	r.POST("/userprofile", func(c *gin.Context) { handler.InsertUserProfile(c, UserProfileService) })
	r.POST("/application", func(c *gin.Context) { handler.InsertApplication(c, ApplicationService) })

	// Get Method
	r.GET("/user", func(c *gin.Context) { handler.GetUser(c, userService) })
	r.GET("/userprofile", func(c *gin.Context) { handler.GetUserProfile(c, UserProfileService) })
	r.GET("/application", func(c *gin.Context) { handler.GetApplication(c, ApplicationService) })

	// RELEASE OPS MODULE
	r.POST("/releaseops/ticket", func(c *gin.Context) { handler.InsertReleaseTicket(c, ReleaseOPSService) })
	r.GET("/releaseops/ticket", func(c *gin.Context) { handler.GetReleaseTicket(c, ReleaseOPSService) })
	r.POST("/releaseops/trigger-build/:ID", func(c *gin.Context) { handler.TriggerBuild(c, ReleaseOPSService) })

	// r.DELETE("/template/:id", handler.DeleteStudent)
	// r.PUT("/template/:id", handler.UpdateStudent)
	return r
}
