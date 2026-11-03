package main

import (
	"bytes"
	"testing"
)

func TestRunPrintsOneLinePerNumber(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if err := run([]string{"5"}, &stdout, &stderr); err != nil {
		t.Fatalf("run: %v", err)
	}
	if want := "1\n2\nFizz\n4\nBuzz\n"; stdout.String() != want {
		t.Errorf("stdout = %q, want %q", stdout.String(), want)
	}
}

func TestRunRejectsBadArguments(t *testing.T) {
	for _, args := range [][]string{{"five"}, {"1", "2"}, {"-unknown"}} {
		var stdout, stderr bytes.Buffer
		if err := run(args, &stdout, &stderr); err == nil {
			t.Errorf("run(%q) succeeded, want an error", args)
		}
	}
}

func TestRunPrintsVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if err := run([]string{"-version"}, &stdout, &stderr); err != nil {
		t.Fatalf("run: %v", err)
	}
	if stdout.Len() == 0 {
		t.Error("no version printed")
	}
}
