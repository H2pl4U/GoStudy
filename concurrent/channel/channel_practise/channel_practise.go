package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	once sync.Once
	lock sync.Mutex
)

func getRand(jobChan chan<- int64) {
	defer wg.Done()
	for i := 0; i < 24; i++ {
		// rand1 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		var randNum int64 = rand.Int63()
		fmt.Println(randNum)
		jobChan <- randNum
	}
	close(jobChan)
}

func putResult(jobChan <-chan int64, results chan<- int64) {
	defer wg.Done()
	for {
		i, ok := <-jobChan
		if !ok {
			break
		}
		int64Num := i
		lenNum := len(strconv.FormatInt(int64Num, 10))
		var sum int64 = 0
		for i := 1; i <= lenNum; i++ {
			last := int64Num % 10
			sum = sum + last
			int64Num = int64Num / 10
		}
		results <- sum
	}
}

func main() {
	start := time.Now()
	jobChan := make(chan int64, 100)
	results := make(chan int64, 100)

	wg.Add(1)
	go getRand(jobChan)

	for i := 1; i <= 24; i++ {
		wg.Add(1)
		go putResult(jobChan, results)
	}
	wg.Wait()
	close(results)
	for i := range results {
		fmt.Println(i)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
