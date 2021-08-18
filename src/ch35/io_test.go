package ch35

import (
	"io/ioutil"
	"testing"
	"github.com/chyidl/begin-golang/src/ch35"
)

func TestIOUtil(t *testing.T) {
	bts, err := ioutil.ReadFile("test/readme.md")
	if err != nil {
		logger.Info("say something")
	}
	t.Log(bts)
}
