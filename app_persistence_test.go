package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigPersistence(t *testing.T) {
	app := NewApp()
	app.ctx = context.Background()

	// Use a temporary file for config during test
	home, _ := os.UserHomeDir()
	originalConfigPath := filepath.Join(home, ".plan-saver-config.json")
	
	// Backup original config if exists
	var backupData []byte
	backupExists := false
	if data, err := os.ReadFile(originalConfigPath); err == nil {
		backupData = data
		backupExists = true
	}

	// Test path
	testPath := "/tmp/test-plan-saver"

	// Test Save
	app.saveConfig(testPath)

	// Test Load
	app.targetPath = "" // Reset
	app.loadConfig()

	if app.targetPath != testPath {
		t.Errorf("Expected targetPath to be %s, got %s", testPath, app.targetPath)
	}

	// Clean up and restore backup
	os.Remove(originalConfigPath)
	if backupExists {
		os.WriteFile(originalConfigPath, backupData, 0644)
	}
}

func TestGetSavedPath(t *testing.T) {
	app := NewApp()
	app.targetPath = "/some/path"
	
	if app.GetSavedPath() != "/some/path" {
		t.Errorf("Expected GetSavedPath to return /some/path, got %s", app.GetSavedPath())
	}
}
