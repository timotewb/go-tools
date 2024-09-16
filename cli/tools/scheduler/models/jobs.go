package models

type Jobs []struct {
	CronTimeTrigger string `json:"cron_time_trigger"`
	Command string `json:"command"`
}