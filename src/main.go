package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	day := flag.Int("day", -1, "which day to run")
	flag.Parse()

	if *day == -1 {
		fmt.Println("Please specify a day to run with -day=N")
		os.Exit(1)
	}

	days := map[int]func(){
		1: D1,
		2: D2,
	}

	if run, ok := days[*day]; ok {
		run()
	} else {
		fmt.Printf("Day %d is not implemented yet.\n", *day)
		os.Exit(1)
	}
}
