package app

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
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
        Hour:   convertPart(parts[1]),
        Dom:    convertPart(parts[2]),
        Month:  convertPart(parts[3]),
        DOW:    convertPart(parts[4]),
    }
    return expr, nil
}

func getMin(part string) []int{
    re := regexp.MustCompile(`^[0-5][0-9]$`)
    if re.MatchString(part) {
        val, _ := strconv.Atoi(part)
        return []int{val}
    } else if part == "*" {
        return make([]int, 59)
    }
    // error
    return []int{-1}
}

func convertPart(part string) int {
    switch part {
    case "*":
        return -9 // Special handling for wildcard
    default:
        val, err := strconv.Atoi(part)
        if err != nil {
            return -1 // Error handling
        }
        return val
    }
}

func ShouldExecuteNow(cronStr string) bool {
	expr, err := paresCronTimeTrigger(cronStr)
	if err != nil {
		log.Fatal(err)
	}
	// parse time to match crontab time trigger
    tn := time.Now()

    // Placeholder
    if expr.Minute == -9 || expr.Hour == -9 || expr.Dom == -9 || expr.Month == -9 || expr.DOW == -9 {
        // If any field is "*", consider the condition met.
        return true
    }
    if (expr.Minute == tn.Minute() || expr.Minute == -9) && (expr.Hour == tn.Hour() || expr.Hour == -9) {
        return true
    }

    return false
}
