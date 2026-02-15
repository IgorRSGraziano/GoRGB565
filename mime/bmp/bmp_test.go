package bmp

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFunc(t *testing.T) {
	f, err := os.Open("testdata/blu.bmp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	bmp, err := Read(f)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(bmp.Header.FileSize)
	fmt.Println(bmp.Header.DataOffset)
}
