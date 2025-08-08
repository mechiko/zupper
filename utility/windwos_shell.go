package utility

import (
	"net/url"
	"path/filepath"

	"golang.org/x/sys/windows"
)

func OpenFileInShell(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	urlFile := url.URL{
		Scheme: "file",
		Path:   path,
	}
	str, err := url.PathUnescape(urlFile.String())
	if err != nil {
		return err
	}
	return startShell(str)
}

// открыть ссылку в браузере
func OpenHttpLinkInShell(urlRaw string) error {
	parsedURL, err := url.Parse("//" + urlRaw)
	if err != nil {
		return err
	}
	parsedURL.Scheme = "http"
	return startShell(parsedURL.String())
}

func OpenHttpsLinkInShell(urlRaw string) error {
	parsedURL, err := url.Parse("//" + urlRaw)
	if err != nil {
		return err
	}
	parsedURL.Scheme = "https"
	return startShell(parsedURL.String())
}

func startShell(url string) error {
	return windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(url), nil, nil, windows.SW_SHOWNORMAL)
}
