package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	// TODO: CHANGE BELOW IMPORTS
	"github.com/sriv/report-plugin-skeleton/env"
	helper "github.com/sriv/report-plugin-skeleton/test_helper"
)

var now = time.Now()

type testNameGenerator struct {
}

func (T testNameGenerator) randomName() string {
	return now.Format(timeFormat)
}

func TestGetReportsDirectory(t *testing.T) {
	userSetReportsDir := filepath.Join(os.TempDir(), randomName())
	os.Setenv(env.GaugeReportsDirEnvName, userSetReportsDir)
	expectedReportsDir := filepath.Join(userSetReportsDir, htmlReport)
	defer os.RemoveAll(userSetReportsDir)

	reportsDir := getReportsDirectory(nil)

	if reportsDir != expectedReportsDir {
		t.Errorf("Expected reportsDir == %s, got: %s\n", expectedReportsDir, reportsDir)
	}
	if !helper.FileExists(expectedReportsDir) {
		t.Errorf("Expected %s report directory doesn't exist", expectedReportsDir)
	}
}

func TestGetReportsDirectoryWithOverrideFlag(t *testing.T) {
	userSetReportsDir := filepath.Join(os.TempDir(), randomName())
	os.Setenv(env.GaugeReportsDirEnvName, userSetReportsDir)
	os.Setenv(env.OverwriteReportsEnvProperty, "true")
	nameGen := &testNameGenerator{}
	expectedReportsDir := filepath.Join(userSetReportsDir, htmlReport, nameGen.randomName())
	defer os.RemoveAll(userSetReportsDir)

	reportsDir := getReportsDirectory(nameGen)

	if reportsDir != expectedReportsDir {
		t.Errorf("Expected reportsDir == %s, got: %s\n", expectedReportsDir, reportsDir)
	}
	if !helper.FileExists(expectedReportsDir) {
		t.Errorf("Expected %s report directory doesn't exist", expectedReportsDir)
	}
}

func randomName() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func TestCreatingReportShouldOverwriteReportsBasedOnEnv(t *testing.T) {
	os.Setenv(env.OverwriteReportsEnvProperty, "true")
	nameGen := getNameGen()
	if nameGen != nil {
		t.Errorf("Expected nameGen == nil, got %s", nameGen)
	}

	os.Setenv(env.OverwriteReportsEnvProperty, "false")
	nameGen = getNameGen()
	switch nameGen.(type) {
	case timeStampedNameGenerator:
	default:
		t.Errorf("Expected nameGen to be type timeStampedNameGenerator, got %s", reflect.TypeOf(nameGen))
	}
}
