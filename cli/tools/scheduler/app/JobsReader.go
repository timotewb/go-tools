package app

import (
	"encoding/json"
	"os"

	"github.com/timotewb/go-tools/cli/tools/scheduler/models"
)

func JobsReader() (models.Jobs, error) {
	var resp models.Jobs

	// read in
	b, err := os.ReadFile("tools/scheduler/data/jobs.json")
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}