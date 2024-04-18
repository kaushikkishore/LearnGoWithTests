package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCores := runtime.NumCPU() // Get the number of CPU cores
	desiredMaxProcs := 2         // Set the desired number of max procs
	runtime.GOMAXPROCS(desiredMaxProcs)
	fmt.Printf("Set GOMAXPROCS to %d (actual number of CPU cores: %d)\n", desiredMaxProcs, numCores)
}
