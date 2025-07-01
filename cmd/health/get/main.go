package main

import (
	"github.com/aaronsisler/services.email/handlers/health"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(health.HealthGetHandler)
}
