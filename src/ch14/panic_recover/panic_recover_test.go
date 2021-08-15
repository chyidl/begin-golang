package panic_recover_test

import (
	"errors"
	"fmt"
	//"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	/*
		defer func() {
			fmt.Println("Finally")
		}()
	*/

	defer func() {
		// Let it Crash 往往是恢复不确定性错误的最好方法
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	fmt.Println("Start")
	// os.Exit(-1)
	panic(errors.New("Something wrong"))
}
