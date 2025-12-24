package saver

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/YuanJun-93/CodeGenesis/internal/pkg/constant"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/parser"
)

// SaveCode saves the parsed code content
func SaveCode(result *parser.CodeResult, typeStr string) (string, error) {
	// Base directory for generated code
	baseDir := "generated"
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}

	timestamp := time.Now().Format("20060102_150405")

	if typeStr == constant.GeneratorTypeHtml {
		// Single file mode
		filename := fmt.Sprintf("gen_%s.html", timestamp)
		fullPath := filepath.Join(baseDir, filename)
		if err := os.WriteFile(fullPath, []byte(result.Html), 0644); err != nil {
			return "", err
		}
		return filepath.Abs(fullPath)
	}

	// Multi-file mode: Create a subdirectory
	projectDirName := fmt.Sprintf("project_%s", timestamp)
	projectDirPath := filepath.Join(baseDir, projectDirName)
	if err := os.MkdirAll(projectDirPath, 0755); err != nil {
		return "", err
	}

	// Save index.html
	if result.Html != "" {
		if err := os.WriteFile(filepath.Join(projectDirPath, "index.html"), []byte(result.Html), 0644); err != nil {
			return "", err
		}
	}
	// Save style.css
	if result.Css != "" {
		if err := os.WriteFile(filepath.Join(projectDirPath, "style.css"), []byte(result.Css), 0644); err != nil {
			return "", err
		}
	}
	// Save script.js
	if result.Js != "" {
		if err := os.WriteFile(filepath.Join(projectDirPath, "script.js"), []byte(result.Js), 0644); err != nil {
			return "", err
		}
	}

	return filepath.Abs(projectDirPath)
}
