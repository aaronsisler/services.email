package services

import (
	"fmt"

	"github.com/aaronsisler/services.email/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendEmail(email models.Email) error {
	// Create a new session in your desired region.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Println("err from session creation")
		fmt.Println(err)
		return err
	}

	svc := ses.New(sess)

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

	result, err := svc.SendEmail(input)

	if err != nil {
		fmt.Println("err from svc.SendEmail")
		fmt.Println(err)
		return err
	}

	fmt.Println("Result MessageId")
	fmt.Println(result.MessageId)

	return nil
}
