package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	go PULL_UPDATE.Add(1)
	func() {
		fmt.Println("testing bar")
		time.Sleep(2000 * time.Millisecond)
	}()
	go MERGE_UPDATE.Add(1)
	func() {
		fmt.Println("testing bar2")
		time.Sleep(2000 * time.Millisecond)
	}()
	go BUILD_UPDATE.Add(1)
	func() {
		fmt.Println("testing bar3")
		time.Sleep(2000 * time.Millisecond)
	}()
}
