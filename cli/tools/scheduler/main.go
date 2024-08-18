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

	for _, job := range jobs{
		fmt.Printf("%v\n-Trigger Now: %v\n\n", job.ExecutingCommand,app.ShouldExecuteNow(job.CronTimeTrigger))
	}

}
