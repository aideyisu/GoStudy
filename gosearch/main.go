package main

import (
	"fmt"
	"io/ioutil"
	// 	"sync"
	"time"
)

var query = "test"
var matches int

var workerCount = 0
var maxWorkerCount = 32

var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundMatch = make(chan bool)

func main() {
	start := time.Now()
	workerCount = 1
	go search("/Users/mfz/Code/aideyisu/", true)
	waitForWorkers()
	fmt.Println(matches, " matches")
	fmt.Println(time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range files {
		name := file.Name()
		if name == query {
			foundMatch <- true
		}
		if file.IsDir() {
			if workerCount < maxWorkerCount {
				searchRequest <- path + name + "/"
			} else {
				go search(path+name+"/", false)
			}
		}
	}
	if master {
		workerDone <- true
	}
}

/*

// sync方法
func main() {
	var A sync.WaitGroup
	A.Add(1)
	start := time.Now()
	go search("/Users/mfz/Code/aideyisu/", &A)
	A.Wait()
	fmt.Println(matches, " matches")
	fmt.Println(time.Since(start))
}

func search(path string, wg *sync.WaitGroup) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range files {
		name := file.Name()
		if name == query {
			matches++
		}
		if file.IsDir() {
			var A sync.WaitGroup
			A.Add(1)
			go search(path+name+"/", &A)
			A.Wait()
		}
	}
	wg.Done()
}
*/
