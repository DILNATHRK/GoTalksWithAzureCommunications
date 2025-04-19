package model

type EmailNotification struct {
	EmailNotificationID uint   `gorm:"primarykey" json:"email_notification_id"`
	Sender              string `gorm:"sender"     json:"sender" validate:"required"`
	Recipient           string `gorm:"recipient"  json:"recipient" validate:"required"`
	Subject             string `gorm:"subject"    json:"subject"`
	Payload             string `gorm:"payload"    json:"payload"`
}
