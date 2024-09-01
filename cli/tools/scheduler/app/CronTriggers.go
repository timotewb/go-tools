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

// package main

// import "fmt"

// type NumberTrigger struct{}

// type GetNum struct {
// 	NumberTrigger *NumberTrigger
// }

// func (s *NumberTrigger) GetNumberTrigger(min, max int) {
// 	fmt.Printf("Min: %v, Max: %v", min, max)
// }

// func (g *GetNum) GetMin() {
// 	g.NumberTrigger.GetNumberTrigger(1, 5)
// }

// func main() {
// 	t1 := GetNum{}
// 	t1.GetMin()
// }



func getMin(part string) []int{

	min := 0
	max := 59
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