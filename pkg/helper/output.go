package helper

import (
	"fmt"
	"github.com/merge/pkg/interval"
)

func Print(title string, results []*interval.Interval) string {
	var result string
	fmt.Println(title)
	for _,interval := range results {
		interval := fmt.Sprintf("[%d,%d]", interval.Low, interval.High)
		result = result + interval
		fmt.Print(interval)
	}
	fmt.Println()
	return result
}
