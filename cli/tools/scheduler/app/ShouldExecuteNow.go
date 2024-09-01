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
    Hour   []int
    DOM    []int
    Month  []int
    DOW    []int
}

func paresCronTimeTrigger(cronStr string) (*CronExpression, error) {
    c := CrontTrigger{}
    parts := strings.Split(cronStr, " ")
    if len(parts) != 5 {
        return nil, fmt.Errorf("invalid cron format")
    }

    expr := &CronExpression{
        Minute: c.MinTrigger(parts[0]),
        Hour: c.HourTrigger(parts[1]), 
        DOM: c.HourTrigger(parts[1]), 
        Month: c.HourTrigger(parts[1]), 
        DOW: c.HourTrigger(parts[1]), 
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
    if checkIntInIntArray(expr.Minute, -1) || checkIntInIntArray(expr.Hour, -1) || checkIntInIntArray(expr.DOM, -1) || checkIntInIntArray(expr.Month, -1) || checkIntInIntArray(expr.DOW, -1) {
        log.Print("Error: invalid value passed in time trigger. Please check input time trigger format.")
        return false
    }
    if checkIntInIntArray(expr.Minute, tn.Minute()) && checkIntInIntArray(expr.Hour, tn.Hour()) && checkIntInIntArray(expr.DOM, tn.Day()) && checkIntInIntArray(expr.Month, int(tn.Month())) && checkIntInIntArray(expr.DOW, int(tn.Weekday())) {
        return true
    }
    return false
}
