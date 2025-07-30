// services/email_sender.go
package services

import "github.com/aaronsisler/services.email/models"

type EmailSender interface {
	SendEmail(email models.Email) error
}
