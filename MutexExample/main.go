package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	money int64
)

func init() {
	money = 1000
}

func addMoney(value int64, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Adding money %d, %d\n", value, money)
	money += value
	mutex.Unlock()
	wg.Done()
}

func withdrawMoney(value int64, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing money %d, %d\n", value, money)
	money -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go addMoney(100, &wg)
	go withdrawMoney(50, &wg)
	go addMoney(500, &wg)
	wg.Wait()

	fmt.Printf("Money I have now: %d\n", money)
}
