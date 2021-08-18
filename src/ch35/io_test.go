package ch35

import (
	"github.com/chyidl/begin-golang/src/log"
	"io/ioutil"
	"testing"
)

func TestIOUtil(t *testing.T) {
	logger := log.Get("TestIOUtil")
	bts, err := ioutil.ReadFile("test/readme.md")
	if err != nil {
		logger.Info("say something")
	}
	logger.Info(bts)
}
