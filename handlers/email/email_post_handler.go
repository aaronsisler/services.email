package email

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Signature struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
}

type Header struct {
	Subject string `json:"subject"`
	From    string `json:"from"`
	To      string `json:"to"`
}

type Email struct {
	Header    Header    `json:"header"`
	Body      string    `json:"body"`
	Signature Signature `json:"signature"`
}

func EmailPostHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// var email Email

	// Unmarshal the JSON string in the body
	// err := json.Unmarshal([]byte(req.Body), &email)

	// if err != nil {
	// 	return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	// }

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
