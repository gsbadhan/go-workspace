package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("threading started..")
	runParallel()
	runAfterAnother()
	time.Sleep(10 * time.Second)

	log.Println("main exit")
}

func runAfterAnother() {
	var wg sync.WaitGroup
	wg.Add(1)
	go waitgroupThread(1, &wg)
	wg.Wait()

	wg.Add(1)
	go waitgroupThread(2, &wg)
	wg.Wait()
}

func waitgroupThread(id int, wg *sync.WaitGroup) {
	log.Printf("%v runAfterAnother WaitGroup\n", id)
	time.Sleep(5 * time.Second)
	defer wg.Done()
}

func runParallel() {
	log.Println("running parallel threads..")
	go parallelThread(1)
	go parallelThread(2)
}

func parallelThread(id int) {
	log.Printf("%v parallelThread\n", id)
}
