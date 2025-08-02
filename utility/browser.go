// Package browser provides helpers to open files, readers, and urls in a browser window.
//
// The choice of which browser is started is entirely client dependant.
package utility

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Stdout is the io.Writer to which executed commands write standard output.
var Stdout io.Writer = os.Stdout

// Stderr is the io.Writer to which executed commands write standard error.
var Stderr io.Writer = os.Stderr

// OpenFile opens new browser window for the file path.
func OpenFile(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	return OpenURL("file://" + path)
}

// OpenReader consumes the contents of r and presents the
// results in a new browser window.
func OpenReader(r io.Reader) error {
	f, err := os.CreateTemp("", "browser.*.html")
	if err != nil {
		return fmt.Errorf("browser: could not create temporary file: %v", err)
	}
	defer f.Close()
	defer os.Remove(f.Name()) // Clean up temp file

	if _, err := io.Copy(f, r); err != nil {
		return fmt.Errorf("browser: caching temporary file failed: %v", err)
	}
	return OpenFile(f.Name())
}

// OpenURL opens a new browser window pointing to url.
func OpenURL(url string) error {
	return openBrowser(url)
}
