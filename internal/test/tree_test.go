package service

import (
	"bytes"
	"io"
	"os"
	"testing"
	"tree/internal/service"
)

// Для функции DirsAndFiles
func TestDirsAndFiles(t *testing.T) {
	out := new(bytes.Buffer)
	err := captureOutput(out, func() {
		service.DirsAndFiles("../../testdata", "")
	})
	if err != nil {
		t.Fatalf("DirsAndFiles failed with error: %v", err)
	}
	expected := `├───project
│   ├───file.txt (19)
│   └───gopher.png (70372)
├───static
│   ├───a_lorem
│   │   ├───dolor.txt (empty)
│   │   ├───gopher.png (70372)
│   │   └───ipsum
│   │       └───gopher.png (70372)
│   ├───css
│   │   └───body.css (28)
│   ├───empty.txt (empty)
│   ├───html
│   │   └───index.html (57)
│   ├───js
│   │   └───site.js (10)
│   └───z_lorem
│       ├───dolor.txt (empty)
│       ├───gopher.png (70372)
│       └───ipsum
│           └───gopher.png (70372)
├───zline
│   ├───empty.txt (empty)
│   └───lorem
│       ├───dolor.txt (empty)
│       ├───gopher.png (70372)
│       └───ipsum
│           └───gopher.png (70372)
└───zzfile.txt (empty)
`
	result := out.String()
	if result != expected {
		t.Errorf("Results not match\nGot:\n%s\nExpected:\n%s", result, expected)
	}
}

// Для функции OnlyDir
func TestOnlyDir(t *testing.T) {
	out := new(bytes.Buffer)
	err := captureOutput(out, func() {
		service.OnlyDir("../../testdata", "")
	})
	if err != nil {
		t.Fatalf("OnlyDir failed with error: %v", err)
	}
	expected := `├───project
├───static
│   ├───a_lorem
│   │   └───ipsum
│   ├───css
│   ├───html
│   ├───js
│   └───z_lorem
│       └───ipsum
└───zline
    └───lorem
        └───ipsum
`
	result := out.String()
	if result != expected {
		t.Errorf("Results not match\nGot:\n%s\nExpected:\n%s", result, expected)
	}
}

func captureOutput(out io.Writer, f func()) error {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = w

	done := make(chan error)
	go func() {
		_, err := io.Copy(out, r)
		done <- err
	}()

	f()
	w.Close()
	os.Stdout = old

	return <-done
}
