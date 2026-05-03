package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	targetPath string
}

type Config struct {
	TargetPath string `json:"targetPath"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.loadConfig()
}

func (a *App) getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".plan-saver-config.json")
}

func (a *App) loadConfig() {
	configPath := a.getConfigPath()
	data, err := os.ReadFile(configPath)
	if err == nil {
		a.targetPath = string(data)
	}
}

func (a *App) saveConfig(path string) {
	configPath := a.getConfigPath()
	_ = os.WriteFile(configPath, []byte(path), 0644)
}

// GetSavedPath returns the path from config
func (a *App) GetSavedPath() string {
	return a.targetPath
}

// SelectFolder opens a directory dialog and returns the selected path
func (a *App) SelectFolder() (string, error) {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Chọn thư mục lưu Plan",
	})
	if err != nil {
		return "", err
	}
	if selection != "" {
		a.targetPath = selection
		a.saveConfig(selection)
	}
	return selection, nil
}

// ExtractAndSave processes the content, extracts markdown blocks, and saves them to the specified path
func (a *App) ExtractAndSave(content string, targetPath string) (string, error) {
	if targetPath == "" {
		return "", fmt.Errorf("chưa chọn thư mục lưu")
	}

	blocks := ExtractMarkdownBlocks(content)
	if len(blocks) == 0 {
		return "", fmt.Errorf("không tìm thấy block markdown nào")
	}

	// Create timestamped subfolder
	timestamp := time.Now().Format("060102-1504") // YYMMDD-HHMM
	folderName := fmt.Sprintf("plan-%s", timestamp)
	saveDir := filepath.Join(targetPath, folderName)

	err := os.MkdirAll(saveDir, 0755)
	if err != nil {
		return "", fmt.Errorf("không thể tạo thư mục: %v", err)
	}

	for _, block := range blocks {
		fileName := SanitizeFileName(block.FileName)
		filePath := filepath.Join(saveDir, fileName)

		err := os.WriteFile(filePath, []byte(block.Content), 0644)
		if err != nil {
			return "", fmt.Errorf("không thể lưu file %s: %v", fileName, err)
		}
	}

	return fmt.Sprintf("Đã lưu %d file vào thư mục: %s", len(blocks), folderName), nil
}
