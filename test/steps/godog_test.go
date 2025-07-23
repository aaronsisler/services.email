package godog

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                "bdd",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Output: colors.Colored(os.Stdout),
			Paths:  []string{"../features"},
			Format: "pretty", // or "progress"
		},
	}.Run()

	os.Exit(status)
}
