package model

// URLMapping represents a mapping from a short code to a destination URL.
type URLMapping struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}
