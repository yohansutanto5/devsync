package main

import (
	"app/constanta"
	"app/model"
	"app/pkg/log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	transactionID := generateTransactionID()
	c.Set("transactionID", transactionID)
	start := time.Now()
	// Incoming request logging
	incomingLog := model.CustomLog{
		Agent:         c.Request.UserAgent(),
		Method:        c.Request.Method,
		ClientIp:      c.ClientIP(),
		Path:          c.Request.URL.Path,
		TransactionID: transactionID,
		Status:        200,
		Code:          constanta.CodeOK,
		Message:       "Incoming Request",
	}
	log.Info(incomingLog)

	defer func() {
		if err := recover(); err != nil {
			// Handle the error, log it, and send an appropriate response.
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			// log.Error(fmt.Sprintf("%v", err))
		}
	}()

	// Process the request
	c.Next()

	// Response logging
	responseLog := model.CustomLog{
		Agent:         c.Request.UserAgent(),
		Method:        c.Request.Method,
		ClientIp:      c.ClientIP(),
		Path:          c.Request.URL.Path,
		TransactionID: transactionID,
		Status:        c.Writer.Status(),
		Duration:      time.Duration(time.Since(start).Milliseconds()),
		Code:          constanta.CodeOK,
		Message:       c.Errors.String(),
	}
	log.Info(responseLog)

}

func generateTransactionID() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
