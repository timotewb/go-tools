package app

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// CronExpression represents a cron expression.
type CronExpression struct {
    Minute []int
    Hour   int
    Dom    int
    Month  int
    DOW    int
}

func paresCronTimeTrigger(cronStr string) (*CronExpression, error) {
    parts := strings.Split(cronStr, " ")
    if len(parts) != 5 {
        return nil, fmt.Errorf("invalid cron format")
    }

    expr := &CronExpression{
        Minute: getMin(parts[0]),
    }
    return expr, nil
}



func checkIntInIntArray(arr []int, target int) bool {
    for _, value := range arr {
        if value == target {
            return true
        }
    }
    return false
}

func ShouldExecuteNow(cronStr string) bool {
	expr, err := paresCronTimeTrigger(cronStr)
    fmt.Printf(" - expr: %v",expr)
	if err != nil {
		log.Fatal(err)
	}
	// parse time to match crontab time trigger
    tn := time.Now()

    // check for any errors
    if checkIntInIntArray(expr.Minute, -1) {
        log.Print("Error: invalid value passed in time trigger. Please check input time trigger format.")
        return false
    }
    if checkIntInIntArray(expr.Minute, tn.Minute()) {
        return true
    }

    return false
}
