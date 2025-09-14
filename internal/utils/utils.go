package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var ErrNotFound = errors.New("not found")

func ReadWordList(fileName *string) (string, error) {
	content, err := os.ReadFile(*fileName)
	if err != nil {
		return "", fmt.Errorf("Error reading wordlist: %v", err)
	}

	// check if the wordlist is valid
	if len(content) < 1 {
		return "", fmt.Errorf("Wordlist should contain at least 1 directory!")
	}

	if content[0] != '/' && content[len(content)-1] != '/' {
		return "", fmt.Errorf("invalid directory format")
	}

	return string(content), nil
}

func ReturnDirectories(content string) []string {
	lines := strings.Split(content, "\n")
	var answer []string
	for _, line := range lines {
		if strings.HasPrefix(line, "/") {
			answer = append(answer, line)
		}
	}

	return answer
}

func CheckDirectory(url string) (bool, error) {
	// initialize a new get request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("Invalid request: %v", err)
	}

	// Initialize a client and send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusNotFound {
		return false, ErrNotFound
	}

	return true, nil
}
