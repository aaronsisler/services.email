package godog

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aaronsisler/services.email/handlers/health"
	"github.com/aws/aws-lambda-go/events"
	"github.com/cucumber/godog"
)

var healthResp events.APIGatewayProxyResponse
var healthHandlerErr error

func iInvokeTheHealthCheckLambdaHandler() error {
	healthResp, healthHandlerErr = health.HealthGetHandler(context.Background(), events.APIGatewayProxyRequest{})
	return healthHandlerErr
}

func theHealthResponseStatusCodeShouldBe(expected int) error {
	if healthResp.StatusCode != expected {
		return fmt.Errorf("expected status %d, got %d", expected, healthResp.StatusCode)
	}
	return nil
}

func theHealthResponseBodyShouldContain(field string) error {
	var bodyMap map[string]string
	err := json.Unmarshal([]byte(healthResp.Body), &bodyMap)
	if err != nil {
		return fmt.Errorf("failed to parse JSON body: %w", err)
	}

	if _, ok := bodyMap[field]; !ok {
		return fmt.Errorf("response body does not contain key: %s", field)
	}

	return nil
}

// Export a register function instead
func registerHealthSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I invoke the health check Lambda handler$`, iInvokeTheHealthCheckLambdaHandler)
	ctx.Step(`^the response status code should be (\d+)$`, theHealthResponseStatusCodeShouldBe)
	ctx.Step(`^the response body should contain "([^"]+)"$`, theHealthResponseBodyShouldContain)
}
