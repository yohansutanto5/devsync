package main

import (
	"app/Integration"
	"app/handler"
	"app/service"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

// Rest API
func setupRoutes() *gin.Engine {
	r := gin.New()
	// Setup Middleware
	r.Use(middleware, gin.Recovery())
	// Setup Admin Endpoint
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Initiate all services and dependency
	IntegrationService := Integration.NewExternalService(&configuration)
	userService := service.NewUserService(database)
	UserProfileService := service.NewUserProfileService(database)
	ApplicationService := service.NewApplicationService(database)
	ReleaseOPSService := service.NewReleaseOPSService(database, IntegrationService)
	DebtService := service.NewDebtService(database)
	requestService := service.NewRequestService(database)
	// Define The route Path
	// System API
	r.GET("/health", func(c *gin.Context) { handler.GetSystemHealth(c, database) })

	// Post Method
	r.POST("/user", func(c *gin.Context) { handler.InsertUser(c, userService) })
	r.POST("/userprofile", func(c *gin.Context) { handler.InsertUserProfile(c, UserProfileService) })
	r.POST("/onboarding/application", func(c *gin.Context) { handler.InsertApplication(c, ApplicationService) })
	r.POST("/debt", func(c *gin.Context) { handler.InsertDebt(c, DebtService) })

	// Get Method
	r.GET("/user", func(c *gin.Context) { handler.GetUser(c, userService) })
	r.GET("/userprofile", func(c *gin.Context) { handler.GetUserProfile(c, UserProfileService) })
	r.GET("/application", func(c *gin.Context) { handler.GetApplication(c, ApplicationService) })
	r.GET("/debt", func(c *gin.Context) { handler.InsertDebt(c, DebtService) })
	r.GET("/request", func(c *gin.Context) { handler.GetRequest(c, requestService) })
	r.GET("/request/action/:token", func(c *gin.Context) { handler.GetRequest(c, requestService) })

	// RELEASE OPS MODULE
	r.POST("/releaseops/ticket", func(c *gin.Context) { handler.InsertReleaseTicket(c, ReleaseOPSService) })
	r.GET("/releaseops/ticket", func(c *gin.Context) { handler.GetReleaseTicket(c, ReleaseOPSService) })
	r.POST("/releaseops/trigger-build/:ID", func(c *gin.Context) { handler.TriggerBuild(c, ReleaseOPSService) })

	return r
}

// Event Driven
func topicRoutes(message *kafka.Message) {
	IntegrationService := Integration.NewExternalService(&configuration)
	ReleaseOPSService := service.NewReleaseOPSService(database, IntegrationService)

	// Perform processing based on the content of the message
	if *message.TopicPartition.Topic == configuration.Kafka.ApprovalTopic {
		handler.ApprovalEventAction(message, ReleaseOPSService)
	} else {
		// Handle messages from other topics or implement a default behavior
		// fmt.Printf("Received message from unknown topic: %s\n", message.TopicPartition.Topic)
	}

}
