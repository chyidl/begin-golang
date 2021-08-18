package ch35

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	t.Log(os.Getwd())
	t.Log(os.Environ())
	t.Log(os.Getenv("PATH"))
}