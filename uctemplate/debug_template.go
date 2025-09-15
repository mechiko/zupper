package uctemplate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mechiko/utility"
)

var pathTemplates = "../uctemplate/templates"

func rootPathTemplates() (out string) {
	defer func() {
		if r := recover(); r != nil {
			out = ""
		}
	}()
	out, err := filepath.Abs(pathTemplates)
	if err != nil {
		return ""
	}
	if utility.PathOrFileExists(out) {
		return out
	} else {
		return ""
	}
}

func readFile(fileAbs string) (string, error) {
	if !utility.PathOrFileExists(fileAbs) {
		return "", fmt.Errorf("file not found: %s", fileAbs)
	}
	contentBytes, err := os.ReadFile(fileAbs)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %w", err)
	}
	return string(contentBytes), nil
}
