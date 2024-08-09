package pprof

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

// https://www.liwenzhou.com/posts/Go/pprof/

func TestPProf(t *testing.T) {
	file, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Printf("open")
		return
	}
	err = pprof.StartCPUProfile(file)
	if err != nil {
		return
	}
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	pprof.StopCPUProfile()
}
