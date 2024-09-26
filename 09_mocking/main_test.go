package main

import (
	"reflect"
	"testing"
)

type SpyCountDownOperation struct {
	Calls []string
}

func (s *SpyCountDownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountDown(t *testing.T) {
	spy := SpyCountDownOperation{}
	CountDown(&spy, &spy)

	got := spy.Calls
	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q but got %q", want, got)
	}

}
