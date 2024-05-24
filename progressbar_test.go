package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	go STAGE1.Add(1)
	func() {
		fmt.Println("testing bar")
		time.Sleep(2000 * time.Millisecond)
	}()
	go STAGE2.Add(1)
	func() {
		fmt.Println("testing bar2")
		time.Sleep(2000 * time.Millisecond)
	}()
	go STAGE3.Add(1)
	func() {
		fmt.Println("testing bar3")
		time.Sleep(2000 * time.Millisecond)
	}()
	go STAGE4.Add(1)
	func() {
		fmt.Println("testing bar4")
		time.Sleep(2000 * time.Millisecond)
	}()
}
