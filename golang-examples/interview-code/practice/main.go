package main

import (
	"fmt"
)

func main() {
	// i, err := strconv.ParseInt("1405544146", 10, 64)
	// if err != nil {
	// 	panic(err)
	// }
	// tm := time.Unix(i, 0)
	// fmt.Println(tm)

	var sec, min, hour, day, week int = 0, 0, 0, 0, 0
	t := 7263

	for t > 0 {
		switch {
		case t > 604800:
			week++
			t -= 604800
		case t > 86400 && t < 604800:
			day++
			t -= 86400
		case t > 3600 && t < 86400:
			hour++
			t -= 3600
		case t > 60 && t < 3600:
			min++
			t -= 60
		default:
			sec++
			t -= 1
		}
	}

	var r string
	units := 0

	if week > 0 {
		units++
	}

	if day > 0 {
		units++
	}

	if hour > 0 {
		units++
		if units > 1 {
			day++
			min = 0
			sec = 0
		}
	}

	if min > 0 {
		units++
		if units > 1 {
			hour++
			sec = 0
		}
	}

	if sec > 0 {
		units++
		if units > 1 {
			min++
		}
	}

	if week > 0 {
		r = fmt.Sprintf("%dw", week)
	}

	if day > 0 {

		r += fmt.Sprintf("%dd", day)
	}

	if hour > 0 {
		r += fmt.Sprintf("%dh", hour)
	}

	if min > 0 {
		r += fmt.Sprintf("%dm", min)
	}

	if sec > 0 {
		r += fmt.Sprintf("%ds", sec)
	}

	fmt.Println(r)
}
