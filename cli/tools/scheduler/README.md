# Scheduler

## schedule data input

data/jobs.json
[
{
"cron_time_trigger": "55 0,6,14,18 * * *",
"executing_command": "echo hello"
},
{
"cron_time_trigger": "0-4,11 * * * *",
"executing_command": "echo now!"
}
]
