package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cuonglm/gogi"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func finder() (string, error) {
	gogiClient, _ := gogi.NewHTTPClient()
	data, err := gogiClient.List()
	if err != nil {
		return "", fmt.Errorf("failed to fetch the list of languages from gitignore.io: %w", err)
	}

	langList := strings.Split(data, ",")

	data, err = gogiClient.Create("go")
	if err != nil {
		return "", fmt.Errorf("failed to create the gogi client: %w", err)
	}

	// fmt.Println(data)

	idx, err := fuzzyfinder.Find(langList, func(i int) string {
		return langList[i]
	})
	if err != nil {
		if errors.Is(fuzzyfinder.ErrAbort, err) {
		}

		return "", fmt.Errorf("an error occurred on fuzzyfinder: %w", err)
	}

	return langList[idx], nil
}
