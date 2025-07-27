package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aaronsisler/services.email/models"
	"github.com/aaronsisler/services.email/services"
	"github.com/aaronsisler/services.email/validators"
	"github.com/aws/aws-lambda-go/events"
)

func EmailPostHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var email models.Email
	err := json.Unmarshal([]byte(req.Body), &email)

	if err != nil {
		fmt.Println("Invalid JSON")
		return errorResponse(400, "Invalid JSON format"), nil
	}

	validationErrors, err := validators.ValidateEmail(email)

	if err != nil {
		fmt.Println("Validation failed with error:", err)
		return errorResponse(400, "Validation failed"), nil
	}

	if len(validationErrors) > 0 {
		fmt.Println("Validation errors found:", validationErrors)
		body, _ := json.Marshal(map[string]interface{}{
			"errors": validationErrors,
		})

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(body),
		}, nil
	}

	err = services.SendEmail()

	if err != nil {
		fmt.Println("services.SendEmail failed")
		return errorResponse(500, "services.SendEmail failed"), nil
	}

	responseBody, _ := json.Marshal(map[string]string{
		"message": "Email received successfully!",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

func errorResponse(status int, msg string) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(map[string]string{"error": msg})

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}
}
