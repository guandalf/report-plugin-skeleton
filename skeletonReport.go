package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/getgauge/common"

	// TODO: CHANGE BELOW IMPORTS
	"github.com/sriv/report-plugin-skeleton/env"
	"github.com/sriv/report-plugin-skeleton/gauge_messages"
	"github.com/sriv/report-plugin-skeleton/listener"
)

const (
	htmlReport      = "html-report"
	setupAction     = "setup"
	executionAction = "execution"
	gaugeHost       = "localhost"
	gaugePortEnv    = "plugin_connection_port"
	pluginActionEnv = "html-report_action"
	timeFormat      = "2006-01-02 15.04.05"
)

type nameGenerator interface {
	randomName() string
}

type timeStampedNameGenerator struct {
}

func (T timeStampedNameGenerator) randomName() string {
	return time.Now().Format(timeFormat)
}

var pluginsDir string

func createExecutionReport() {
	pluginsDir, _ = os.Getwd()
	os.Chdir(env.GetProjectRoot())
	listener, err := listener.NewGaugeListener(gaugeHost, os.Getenv(gaugePortEnv))
	if err != nil {
		fmt.Println("Could not create the gauge listener")
		os.Exit(1)
	}
	listener.OnSuiteResult(createReport)
	listener.Start()
}

func createReport(suiteResult *gauge_messages.SuiteExecutionResult) {
	projectRoot, err := common.GetProjectRoot()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	reportsDir := getReportsDirectory(getNameGen())
	fmt.Println(projectRoot, reportsDir)
	// suiteResult.GetSuiteResult() will give you the execution result in a struct of gauge_messages.SuiteExecutionResult

	// TODO: insert YOUR LOGIC HERE to process the raw result

	// projectRoot and reportsDir have been evaluated. Feel free to discard if they aren't used.
	// Remember to delete unused methods/types.
}

func getNameGen() nameGenerator {
	var nameGen nameGenerator
	if env.ShouldOverwriteReports() {
		nameGen = nil
	} else {
		nameGen = timeStampedNameGenerator{}
	}
	return nameGen
}

func getReportsDirectory(nameGen nameGenerator) string {
	reportsDir, err := filepath.Abs(os.Getenv(env.GaugeReportsDirEnvName))
	if reportsDir == "" || err != nil {
		reportsDir = env.DefaultReportsDir
	}
	env.CreateDirectory(reportsDir)
	var currentReportDir string
	if nameGen != nil {
		currentReportDir = filepath.Join(reportsDir, htmlReport, nameGen.randomName())
	} else {
		currentReportDir = filepath.Join(reportsDir, htmlReport)
	}
	env.CreateDirectory(currentReportDir)
	return currentReportDir
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}
