package main

import (
	"bytes"
	"fmt"
	"github.com/merge/pkg"
	"github.com/merge/pkg/handler"
	"github.com/merge/pkg/helper"
	"gopkg.in/alecthomas/kingpin.v2"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// these are command line commands and flags
var (
	// the version
	Version string
	app = kingpin.New("merge", "A interval merge tool")
	// -v means verbose and prints some details while working
	verbose = app.Flag("verbose", "Verbose mode.").Short('v').Bool()
	// interval - a valid interval list must follow this command
	interval = app.Command("interval", "use an interval list as argument.")
	// the receiver of the interval list
	intervalList = interval.Arg("intervals", "input interval list to be merged in the form of [2,6][10,14][4,8].").String()
	// file - a file containing a multiline interval list name must follow this command. All whitespaces are ignored while parsing. Intervals boundaries can be twisted.
	file = app.Command("file", "put an interval list by input file.")
	// the receiver of the file name
	fileName = file.Arg("fileName", "input file with interval list to be merged.").String()
	// gendata - generates test data
	genIntervals = app.Command("gendata", "generates test data in file 'gen.intervals'")
	// the receiver of number of intervals to generate
	numGen = genIntervals.Arg("numIntervals", "number of intervals to be generated").Int()
)

func main() {
	fmt.Printf("Version %s\n", Version)
	// init command line parser
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	defer helper.Elapsed("total", *verbose)()
	var maxHeap uint64 = 0
	var minHeap uint64 = 1000000000
	var mStats runtime.MemStats

	go func() {
		for {
			runtime.ReadMemStats(&mStats)
			if mStats.HeapInuse > maxHeap {
				maxHeap = mStats.HeapInuse
			}
			if mStats.HeapInuse < minHeap {
				minHeap = mStats.HeapInuse
			}
			time.Sleep(10*time.Millisecond)
		}
	}()

	// prepare parser handler and reader
	handler := &handler.InputHandler{Verbose: *verbose}
	var reader *bytes.Reader

	switch cmd {
	// interval as argument
	case interval.FullCommand():
		reader = bytes.NewReader([]byte(*intervalList))
		parser := pkg.NewParser(reader, handler, *verbose)
		err := parser.Parse()
		if err != nil {
			panic(err)
		}

	// file as input
	case file.FullCommand():
		file, err := os.Open(*fileName)
		if err != nil {
			panic(fmt.Errorf("Error opening file %s", *fileName))
		}
		defer file.Close()
		parser := pkg.NewParser(file, handler, *verbose)
		err = parser.Parse()
		if err != nil {
			panic(err)
		}
	case genIntervals.FullCommand():
		file, err := os.OpenFile("./gen.intervals", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		rand.Seed(time.Now().UnixNano())
		intLen := 5
		max := *numGen * intLen
		if err != nil {
			panic(fmt.Errorf("invalid number for generating test data given %s", err))
		}
		for numVals := 1; numVals <= *numGen; numVals ++ {
			low := rand.Intn(max)
			high := rand.Intn(intLen*3) + low +1
			fmt.Fprintf(file, "[%d,%d]", low, high)
			if numVals % 20 == 0 {
				fmt.Fprintln(file, "")
			}
		}
		fmt.Printf("generated successfully %d intervals\n", *numGen)
		return
	}

	// print the parsed intervals
	if len(handler.Array.Inputs) < 100 {
		helper.Print("Inputs:", handler.Array.Inputs)
	} else {
		fmt.Printf("%d intervals before merge\n", len(handler.Array.Inputs))
	}

	// merge overlapping intervals
	results := handler.Merge()
	// output results
	if len(results) < 100 {
		helper.Print("Results:", results)
	} else {
		fmt.Printf("%d intervals after merge\n", len(results))
	}
	if *verbose {
		fmt.Printf("peak heap used: %d MiB\n", bToMb(maxHeap-minHeap))
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}