package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	stdOut := os.Stdout

	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	main()

	_ = writer.Close()

	result, _ := io.ReadAll(reader)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320") {
		t.Errorf("expected $34320. got %s", output)
	}

}
