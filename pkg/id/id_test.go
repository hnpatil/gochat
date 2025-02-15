package id

import (
	"testing"
	"time"
)

func Test_ID(t *testing.T) {
	store := map[string]struct{}{}
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Microsecond)
		println(New().String())
		store[New().String()] = struct{}{}
	}

	println(len(store))
}
