package lock_test

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"
)

func init() {
	// set the integer value used in generating random numbers
	rand.Seed(time.Now().Unix())
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func reader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	// RLock and RUlock
	sleep()
	m.RLock()
	c <- 1
	sleep()
	c <- -1
	m.RUnlock()
	wg.Done()
}

func writer(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	// Lock and Unlock
	sleep()
	// new readers are blocked,
	m.Lock()
	c <- 1
	sleep()
	c <- -1
	m.Unlock()
	wg.Done()
}

func TestRWMutex(t *testing.T) {
	// RWMutex allows for either at least one reader or exactly one writer
	var m sync.RWMutex
	var rs, ws int
	rsCh := make(chan int)
	wsCh := make(chan int)

	go func() {
		for {
			select {
			case n := <-rsCh:
				rs += n
			case n := <-wsCh:
				ws += n
			}
			// How many readers and writers are currently executing it
			fmt.Printf("%s%s\n", strings.Repeat("R", rs), strings.Repeat("W", ws))
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reader(rsCh, &m, &wg)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go writer(wsCh, &m, &wg)
	}
	wg.Wait()
}
