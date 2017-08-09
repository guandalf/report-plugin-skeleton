package main

import (
	"os"

	// TODO: CHANGE BELOW IMPORTS
	"github.com/sriv/report-plugin-skeleton/env"
)

func main() {
	action := os.Getenv(pluginActionEnv)
	if action == setupAction {
		env.AddDefaultPropertiesToProject()
	} else if action == executionAction {
		createExecutionReport()
	}
}
