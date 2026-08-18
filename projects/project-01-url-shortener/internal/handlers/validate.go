package handlers

import (
	"fmt"
	"net/url"
)

func ValidateURL(raw string) error {
	if raw == "" {
		return fmt.Errorf("url required")
	}
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid url")
	}
	if u.Scheme != "https" {
		return fmt.Errorf("only https scheme allowed")
	}
	return nil
}
