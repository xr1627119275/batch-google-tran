package test

import (
	"fmt"
	"os"
	"testing"
)

func TestStat(t *testing.T) {
	fmt.Println(os.Args[0])
	stat, err := os.Stat(os.Args[0])
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stat.Size())
}
