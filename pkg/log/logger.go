package log

import (
	"app/constanta"
	"app/model"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// Set log file name using the current date
	logFileName := "/home/yohan/standard-restAPI/app-" + time.Now().Format("2006-01-02") + ".log"

	// Open the log file for writing, create it if it doesn't exist
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = logFile
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Configure logrus to produce JSON logs
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Debug(transactionID int, message string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"data":          data,
		"status":        constanta.DebugStatus,
		"Code":          constanta.DebugCode,
	}).Debug(message)
}

func DebugIssue(transactionID int, message string, mode string, data interface{}) {
	if mode == "debug" {
		log.WithFields(logrus.Fields{
			"transactionID": transactionID,
			"data":          data,
			"status":        constanta.DebugStatus,
			"Code":          constanta.DebugCode,
		}).Debug(message)
	}
}

func Error(transactionID int, message string, code string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"data":          data,
		"status":        constanta.ErrorStatus,
		"Code":          code,
	}).Error(message)
}

func Info(param model.CustomLog) {
	log.WithFields(logrus.Fields{
		"transactionID": param.TransactionID,
		"status":        param.Status,
		"Code":          param.Code,
		"clientIp":      param.ClientIp,
		"method":        param.Method,
		"path":          param.Path,
		"agent":         param.Agent,
		"duration":      param.Duration,
	}).Info(param.Message)
}

func Fatal(message string) {
	log.Fatal(message)
}

func System(message string) {
	log.Info(message)
}

func Warning(transactionID int, message string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"data":          data,
		"status":        constanta.ErrorStatus,
		"Code":          "code",
	}).Warning(message)
}
