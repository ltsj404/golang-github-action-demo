package main

import (
	"testing"
)

func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
