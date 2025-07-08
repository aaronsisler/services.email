package main

import (
	"github.com/aaronsisler/services.email/handlers/email"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(email.EmailPostHandler)
}
