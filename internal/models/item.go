package models

// Channel holds the live status of a Kick channel.
type Channel struct {
	Name string `json:"name"`
	Live bool   `json:"live"`
}
