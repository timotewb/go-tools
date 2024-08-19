package app

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

//ref: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html

func getMin(part string) []int{

	min := 0
	max := 59
	// check if start and return
	if part == "*" {
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
	parts := strings.Split(part, ",")
	for _, p := range parts{
		//----------------------------------------------------------------------------------------
		// split range (e.g. 1-4) into pieces
		//----------------------------------------------------------------------------------------
		if strings.Contains(p, "-"){
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
				return result
			}
		}
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
	return result
}

func getHour(part string) []int{

	min := 0
	max := 59
	// check if start and return
	if part == "*" {
        result := make([]int, max)
		for i := range result {
            result[i] = min+i
        }
		return result
    }

	// process each value
	var result []int
	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3])$`)
	parts := strings.Split(part, ",")
	for _, p := range parts{
		if re.MatchString(p) {
			val, _ := strconv.Atoi(p)
			if val >=min && val <= max {
				result = append(result, val)
			}
		} else {
			return []int{-1}
		}
	}
	return result
}

func getDOM(part string) []int{
    re := regexp.MustCompile(`^([1-9]|0[1-9]|1[0-9]|2[0-9]|3[0-1])$`)
    if re.MatchString(part) {
        val, _ := strconv.Atoi(part)
		if val >=0 && val < 31 {
			return []int{val}
		}
    } else if part == "*" {
        return make([]int, 30)
    }
    // error
    return []int{-1}
}

func getMonth(part string) []int{
    re := regexp.MustCompile(`^([1-9]|0[1-9]|1[0-2])$`)
    if re.MatchString(part) {
        val, _ := strconv.Atoi(part)
		if val >=0 && val < 12 {
			return []int{val}
		}
    } else if part == "*" {
        return make([]int, 11)
    }
    // error
    return []int{-1}
}

func getDOW(part string) []int{
    re := regexp.MustCompile(`^([1-7]|0[1-7]|)$`)
    if re.MatchString(part) {
        val, _ := strconv.Atoi(part)
		if val >=0 && val < 7 {
			return []int{val}
		}
    } else if part == "*" {
        return make([]int, 6)
    }
    // error
    return []int{-1}
}