package model

type ErrorResponse struct {
	TransactionID int
	Code          string
	Message       string
	Details       any
}
