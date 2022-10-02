package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)
func fabonacci(n int) int{
	fabonacci := make([]int, n)
	fabonacci[0] = 1
	fabonacci[1] = 2
	for i := 2; i < n; i++{
		fabonacci[i] = fabonacci[i - 1] + fabonacci[i - 2]
	}
	return fabonacci[n - 1]
}
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	res := fabonacci(7)
	fmt.Println(res)
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}
