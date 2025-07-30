package godog

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aaronsisler/services.email/handlers/email"
	"github.com/aaronsisler/services.email/test/mocks"
	"github.com/aws/aws-lambda-go/events"
	"github.com/cucumber/godog"
)

var (
	emailResponse    events.APIGatewayProxyResponse
	emailHandlerErr  error
	emailRequestBody string

	mockSender   *mocks.MockEmailSender
	emailHandler *email.EmailHandler
)

func iHaveARequestWithTheMissingFromField() error {
	emailRequestBody = `{
		"header": {
			"subject": "Test Subject",
			"to": "recipient@example.com"
		}
	}`
	return nil
}

func iHaveARequestWithAllTheRequiredFields() error {
	emailRequestBody = `{
		"header": {
			"subject": "Test Subject",
			"to": "recipient@example.com",
			"from": "sender@example.com"
		},
		"body": "Hello!"
	}`
	return nil
}

func iInvokeTheEmailLambdaHandler() error {
	emailResponse, emailHandlerErr = emailHandler.EmailPostHandler(
		context.Background(),
		events.APIGatewayProxyRequest{
			Body: emailRequestBody,
		},
	)
	return emailHandlerErr
}

func theEmailResponseStatusCodeShouldBe(expected int) error {
	if emailResponse.StatusCode != expected {
		return fmt.Errorf("expected status code %d, got %d", expected, emailResponse.StatusCode)
	}
	return nil
}

func theEmailResponseBodyShouldContain(key string) error {
	var body map[string]any
	if err := json.Unmarshal([]byte(emailResponse.Body), &body); err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}

	if _, ok := body[key]; !ok {
		return fmt.Errorf("expected body to contain key: %s", key)
	}
	return nil
}

func theErrorShouldHaveTheCorrectFields() error {
	var body map[string]any
	if err := json.Unmarshal([]byte(emailResponse.Body), &body); err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}

	errors, ok := body["errors"].([]any)
	if !ok {
		return fmt.Errorf("'errors' field is not an array")
	}

	expectedFields := []string{"Email.Header.From"}

	for _, field := range expectedFields {
		found := false
		for _, errItem := range errors {
			if errMap, ok := errItem.(map[string]any); ok {
				if errStr, ok := errMap["field"].(string); ok && strings.Contains(errStr, field) {
					found = true
					break
				}
			}
		}
		if !found {
			return fmt.Errorf("expected error message for field '%s' not found", field)
		}
	}

	return nil
}

func theEmailSenderShouldBeCalled() error {
	if !mockSender.WasCalled {
		return fmt.Errorf("expected SendEmail to be called")
	}
	return nil
}

func registerEmailSteps(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Set up mock and inject into handler
		mockSender = &mocks.MockEmailSender{}
		emailHandler = &email.EmailHandler{
			Sender: mockSender,
		}
		return ctx, nil
	})

	ctx.Step(`^I have a request with all the required fields$`, iHaveARequestWithAllTheRequiredFields)
	ctx.Step(`^I have a request with the missing from field$`, iHaveARequestWithTheMissingFromField)
	ctx.Step(`^I invoke the email Lambda handler$`, iInvokeTheEmailLambdaHandler)
	ctx.Step(`^the email response status code should be (\d+)$`, theEmailResponseStatusCodeShouldBe)
	ctx.Step(`^the email response body should contain "([^"]+)"$`, theEmailResponseBodyShouldContain)
	ctx.Step(`^the "error" should have the correct fields$`, theErrorShouldHaveTheCorrectFields)
	ctx.Step(`^the email sender should be called$`, theEmailSenderShouldBeCalled)

}
