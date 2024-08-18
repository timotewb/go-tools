package models

type Jobs []struct {
	CronTimeTrigger string `json:"cron_time_trigger"`
	ExecutingCommand string `json:"executing_command"`
}