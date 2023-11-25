package model

import "time"

type ErrorResponse struct {
	TransactionID int
	Code          string
	Message       string
	Details       any
}

type CustomLog struct {
	TransactionID int    `json:"transactionID"`
	Code          string `json:"code"`
	Status        int    `json:"status"`
	Message       string `json:"message"`
	Method        string `json:"method"`
	Path          string
	Duration      time.Duration
	ClientIp      string
	Agent         string

	Data interface{} `json:"data"`
}
