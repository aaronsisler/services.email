package main

import (
	"log"

	"github.com/aaronsisler/services.email/handlers/email"
	"github.com/aaronsisler/services.email/services"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	emailSender, err := services.NewDefaultEmailSender()

	if err != nil {
		log.Fatalf("failed to initialize email sender: %v", err)
	}

	handler := &email.EmailHandler{
		Sender: emailSender,
	}

	lambda.Start(handler.EmailPostHandler)
}
