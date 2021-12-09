package main

import "testing"

func TestShow(t *testing.T) {
	s := Show()
	want := "test successfully"
	if s != want {
		t.Errorf("Show() = %v, want = %v", s, want)
	}
	t.Logf("Show() = %v, want = %v", s, want)

}
