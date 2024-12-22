package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {
	parsed, err := url.Parse(u)

	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	hostAndPath := parsed.Host + parsed.Path
	hostAndPath = strings.ToLower(hostAndPath)

	return hostAndPath, nil
}
