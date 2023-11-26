package model

type ApprovalTopicMessage struct {
	Name          string `convert:"Name"`
	ID            int
	Action        string
	Service       string
	TransactionID int
}
