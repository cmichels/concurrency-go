package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("epsilon")

	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("expcected epsilon. got %s", msg)
	}
}

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	msg = "epsilon2"

	printMessage()

	_ = writer.Close()

	result, _ := io.ReadAll(reader)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon2") {
		t.Errorf("expected epsilon2. got %s", output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	reader, writer, _ := os.Pipe()

	os.Stdout = writer

  main()

	_ = writer.Close()

	result, _ := io.ReadAll(reader)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("expected Hello, universe!. got %s", output)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expected Hello, cosmos!. got %s", output)
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("expected Hello, world!. got %s", output)
	}
}
