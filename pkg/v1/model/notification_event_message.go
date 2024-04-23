package model

type NotificationEventMessage struct {
	QueueID string `json:"id"`
	Status  string `json:"status"`
}
