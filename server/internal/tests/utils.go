package tests

import "testing"

func AssertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	f()
}

func AssertNotPanic(t *testing.T, f func()) {
	go func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	f()
}
