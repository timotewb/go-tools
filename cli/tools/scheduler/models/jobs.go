package models

type Jobs []struct {
	CronTimeTrigger string `json:"cron_time_trigger"`
	EndPoint string `json:"end_point"`
	Body string `json:"body"`
}