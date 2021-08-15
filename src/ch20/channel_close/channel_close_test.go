package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭channel
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func dataReceiver1(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i <= 20; i++ {
			// 通道被关闭 持续获取通道类型的默认值
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	//wg.Add(1)
	//dataReceiver(ch, &wg)
	//wg.Add(1)
	//dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver1(ch, &wg)
	wg.Wait()
}
