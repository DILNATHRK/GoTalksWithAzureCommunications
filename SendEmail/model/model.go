package model

type EmailNotification struct {
	EmailNotificationID uint   `json:"email_notification_id"`
	Sender              string `json:"sender"     validate:"required"`
	Recipient           string `json:"recipient"  validate:"required"`
	Subject             string `json:"subject"`
	Payload             string `json:"payload"`
}
