package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		finishedList = []Philosopher{}

		dine()

		if len(finishedList) != 5 {
			t.Errorf("expected 5. got %d", len(finishedList))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"0", time.Second * 0},
		{"1/4", time.Millisecond * 250},
		{"1/2", time.Millisecond * 500},
		{"1", time.Second * 1},
	}

	for _, e := range theTests {
		finishedList = []Philosopher{}
		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay

		dine()

		if len(finishedList) != 5 {
			t.Errorf("test: %s. expected 5. got %d", e.name, len(finishedList))
		}
	}
}
