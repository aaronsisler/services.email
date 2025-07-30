package services

import (
	"fmt"

	"github.com/aaronsisler/services.email/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SESEmailSender struct {
	sesClient *ses.SES
}

func NewSESEmailSender(region string) (*SESEmailSender, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &SESEmailSender{
		sesClient: ses.New(sess),
	}, nil
}

func (s *SESEmailSender) SendEmail(email models.Email) error {
	charSet := "UTF-8"

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice([]string{email.Header.To}),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(email.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(email.Header.Subject),
			},
		},
		Source: aws.String(email.Header.From),
	}

	result, err := s.sesClient.SendEmail(input)
	if err != nil {
		fmt.Println("SES send failed:", err)
		return err
	}

	fmt.Println("SES send success:", *result.MessageId)
	return nil
}
