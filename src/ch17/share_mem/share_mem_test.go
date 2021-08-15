package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	// 并发竞争
	counter := 0

	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 线程安全
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	// 并发竞争
	counter := 0

	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 线程安全
func TestCounterThreadSafeWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	// 并发竞争
	counter := 0

	for i := 0; i < 5000; i++ {
		// 增加等待量
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	// 阻塞语句
	wg.Wait()
	t.Logf("counter = %d", counter)
}
