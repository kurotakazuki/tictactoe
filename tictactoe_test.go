package main

import (
	"fmt"
	"testing"
)
//subtest1のテスト
func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	if b.get(0, 0) != "o" {
		t.Error("Test01 is failed.")
	}
	b.show()
}
func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 2, 0, 0, 2, 0, 1, 0},
	}
	b.put(1, 1, "x")
	if b.get(1, 1) != "x" {
		t.Error("Test02 is failed.")
	}
	b.show()
	fmt.Print("\n")
}
func TestPutToken03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 1, 2, 0, 0, 0, 0, 0},
	}
	b.put(2, 2, "o")
	if b.get(2, 2) != "o" {
		t.Error("Test03 is failed.")
	}
	b.show()
	fmt.Print("\n")
}