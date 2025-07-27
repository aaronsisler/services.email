package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendEmail() error {
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

	sender := "aaron.sisler@eandbsolutions.com"
	recipient := "aaron.sisler@gmail.com"
	subject := "My SES Test Email"
	htmlBody := "<h1>Hello from AWS SES!</h1><p>This is a test email sent with Go.</p>"
	textBody := "Hello from AWS SES! This is a test email sent with Go."
	charSet := "UTF-8"

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice([]string{recipient}),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		fmt.Println("err from svc.SendEmail")
		fmt.Println(err)
		return err
	}

	fmt.Println("Result")
	fmt.Println(result)

	return nil
}
