package godog

import (
	"github.com/cucumber/godog"
)

// This function will register all steps from different files
func InitializeScenario(ctx *godog.ScenarioContext) {
	// Register health steps
	registerHealthSteps(ctx)

	// Register email steps
	registerEmailSteps(ctx)
}
