package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "hello world"

	wg.Add(2)

	go updateMessage("x")
	go updateMessage("goodbye")
	wg.Wait()

	if msg != "goodbye" {
		t.Error("incorrect")
	}
}
