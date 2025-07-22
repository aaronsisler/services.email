package godog

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

// func TestMain(m *testing.M) {
// 	status := godog.TestSuite{
// 		Name:                "bdd",
// 		ScenarioInitializer: InitializeScenario,
// 		Options: &godog.Options{
// 			Output: colors.Colored(os.Stdout),
// 			Paths:  []string{"../features"},
// 			Format: "pretty", // or "progress"
// 		},
// 	}.Run()

// 	os.Exit(status)
// }

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"../features"}, // Make sure this is correct
	}

	status := godog.TestSuite{
		Name:                "email",
		ScenarioInitializer: registerEmailSteps, // <- this must be correct
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
