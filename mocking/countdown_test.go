package main

import (
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	SpySleeperPrinter := &SpyCountdownOperations{}

	Countdown(SpySleeperPrinter, SpySleeperPrinter)

	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, SpySleeperPrinter.Calls) {
		t.Errorf("wanted calls %v got %v", want, SpySleeperPrinter.Calls)
	}
}
