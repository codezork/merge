package helper

import (
	"fmt"
	"time"
)

func Elapsed(what string, verbose bool) func() {
	start := time.Now()
	return func() {
		if verbose {
			fmt.Printf("%s took %v\n", what, time.Since(start))
		}
	}
}
