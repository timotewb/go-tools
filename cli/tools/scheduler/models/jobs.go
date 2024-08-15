package models

type Jobs []struct {
	Minute string `json:"minute"`
	Hour string `json:"hour"`
	DOM string `json:"dom"`
	Month string `json:"month"`
	DOW string `json:"dow"`
	ExecutingCommand string `json:"executing_command"`
}