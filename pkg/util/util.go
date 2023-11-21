package util

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
)

func EchoTest() string {
	return "Success"
}

func DateConvert() time.Time {
	return Date()
}

func GetTransactionID(c *gin.Context) int {
	// transactionID, _ := c.Get("transactionID")
	return ConvertToInt(GetTransactionIDString(c))
}
func GetTransactionIDString(c *gin.Context) string {
	transactionID, _ := c.Get("transactionID")
	return ConvertToString(transactionID)
}
func GenerateRandomString(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	// Remove any non-alphanumeric characters to get a 5-character string
	return randomString[:length]
}
