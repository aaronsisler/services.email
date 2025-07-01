package health

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func HealthGetHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	responseBody, err := json.Marshal(map[string]string{
		"message": "The current time is " + time.Now().UTC().String(),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}
