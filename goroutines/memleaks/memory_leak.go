package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	count := 100000000
	var memprofile = "memleak.mprof"
	fmt.Println(memprofile)
	for i := 0; i < count; i++ {
		leaking()
	}
	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
	wg.Wait()
}

func leaking() int {
	ch := make(chan int, 2)
	go func() {
		ch <- 0
	}()
	go func() {
		ch <- 1
	}()
	_ = <-ch
	return <-ch
}
