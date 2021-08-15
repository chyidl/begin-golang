package remote_package_test

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
	// Create a new map.
	m := cmap.New()

	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")
	fmt.Println(m)

	// Retrieve item from map.
	if tmp, ok := m.Get("foo"); ok {
		bar := tmp.(string)
		fmt.Println(bar)
	}

	// Removes item under key "foo"
	m.Remove("foo")
	fmt.Println(m)
}
