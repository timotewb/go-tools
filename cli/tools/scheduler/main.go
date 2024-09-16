package scheduler

import (
	"fmt"
	"log"

	"github.com/timotewb/go-tools/cli/tools/scheduler/app"
)

func Main() {
	// read in jobs and execute
	jobs, err := app.JobsReader()
	if err != nil {
		log.Print(err)
		return
	}

	app.ExecuteCommand("echo timmy!")

	for _, job := range jobs{
		fmt.Printf("-Trigger Now: %v\n", app.ShouldExecuteNow(job.CronTimeTrigger))
		fmt.Printf("-Command: %v\n\n", job.Command)
		if app.ShouldExecuteNow(job.CronTimeTrigger){
			go app.ExecuteCommand(job.Command)
		}
	}

}
