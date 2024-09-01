package app

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//ref: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html

// "parent" type to inject into the called struct
type NumberTrigger struct{}

// called struct
type CrontTrigger struct{
	NumberTrigger *NumberTrigger
}

// the "child" functions
func (c *CrontTrigger) MinTrigger(part string) []int{
	return c.NumberTrigger.GetTriggers(part, 0,59)
}
func (c *CrontTrigger) HourTrigger(part string) []int{
	return c.NumberTrigger.GetTriggers(part, 0,23)
}
func (c *CrontTrigger) DOMTrigger(part string) []int{
	return c.NumberTrigger.GetTriggers(part, 1,31)
}
func (c *CrontTrigger) MonthTrigger(part string) []int{
	return c.NumberTrigger.GetTriggers(part, 1,12)
}
func (c *CrontTrigger) DOWTrigger(part string) []int{
	return c.NumberTrigger.GetTriggers(part, 1,7)
}

func (n * NumberTrigger) GetTriggers(part string, min, max int) []int{

	// check if start and return
	if strings.Contains(part, "*") {
        result := make([]int, max)
		for i := range result {
            result[i] = min+i
        }
		return result
    }

	//----------------------------------------------------------------------------------------
	// process each value
	//----------------------------------------------------------------------------------------
	var result []int
	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])$`)

	//----------------------------------------------------------------------------------------
	// split values (e.g. 1,4,5) into parts - comma is priority 2 (after space)
	//----------------------------------------------------------------------------------------
	parts := strings.Split(part, ",")
	for _, p := range parts{
		if strings.Contains(p, "-") {
			//----------------------------------------------------------------------------------------
			// split range (e.g. 1-4) into pieces - hyphen is priority 3
			//----------------------------------------------------------------------------------------
			pieces := strings.Split(p, "-")
			if len(pieces) != 2 {
				log.Printf("Error: length of '%v' not equal to two.", pieces)
				return []int{-1}
			} else {
				//----------------------------------------------------------------------------------------
				// get to and from values from range
				//----------------------------------------------------------------------------------------
				from, err := strconv.Atoi(pieces[0])
				if err != nil{
					log.Printf("Error: could not convert '%v' to integer", pieces[0])
					return []int{-1}
				}
				to, err := strconv.Atoi(pieces[1])
				if err != nil{
					log.Printf("Error: could not convert '%v' to integer", pieces[1])
					return []int{-1}
				}
				for i := range make([]int, to - from+1) {
					if from+i >= min && from+i <= max{
						result = append(result, from+i)
					} else {
						log.Printf("Error: invalid value from range '%v'", from+1)
						return []int{-1}
					}
				}
			}
		} else if strings.Contains(p, "/") {
			//----------------------------------------------------------------------------------------
			// split range (e.g. 1-4) into pieces - slash is priority 3
			//----------------------------------------------------------------------------------------
			pieces := strings.Split(p, "/")
			fmt.Printf(" - Start: %v\n", pieces[0])
			fmt.Printf(" - Increment: %v\n", pieces[1])
			i, err := strconv.Atoi(pieces[0])
			if err != nil{
				log.Printf("Error: could not convert '%v' to integer", pieces[0])
				return []int{-1}
			}
			inc, err := strconv.Atoi(pieces[1])
			if err != nil{
				log.Printf("Error: could not convert '%v' to integer", pieces[0])
				return []int{-1}
			}
			result = append(result, i)
			for i + inc < max {
				i = i + inc
				result = append(result, i)
			}

		} else {
			//----------------------------------------------------------------------------------------
			// check string for value match
			//----------------------------------------------------------------------------------------
			if re.MatchString(p) {
				val, _ := strconv.Atoi(p)
				if val >=min && val <= max {
					result = append(result, val)
				}
			} else {
				log.Printf("Error: could not match '%v' to required format", p)
				return []int{-1}
			}
		}
	}
	return removeDuplicates(result)
}


func removeDuplicates(arr []int) []int {
    seen := map[int]bool{}
    uniqueArr := make([]int, 0)
    
    for _, num := range arr {
        if !seen[num] {
            seen[num] = true
            uniqueArr = append(uniqueArr, num)
        }
    }
	
    // Sort the array in ascending order
    sort.Ints(uniqueArr)
    
    return uniqueArr
}