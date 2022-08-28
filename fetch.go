package main

import (
	"fmt"
	"strings"

	"github.com/cuonglm/gogi"
)

func fetch(langs []string) (string, error) {
	langStr := strings.Join(langs, ",")
	gogiClient, _ := gogi.NewHTTPClient()

	data, err := gogiClient.Create(langStr)
	if err != nil {
		return "", fmt.Errorf("failed to fetch the content from gitignore.io: %w", err)
	}

	return data, nil
}
