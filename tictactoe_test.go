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

//subtest2のテスト
func TestPutToken04(t *testing.T) {
	b := &Board{
		tokens: []int{1, 2, 1, 2, 1, 2, 0, 0, 1},
	}
	str := b.judge()
	if str != "o" {
		t.Error("Test04 is failed")
	}
	b.show()
	fmt.Print("\n")
}
func TestPutToken05(t *testing.T) {
	b := &Board{
		tokens: []int{2, 1, 1, 2, 1, 0, 2, 0, 0},
	}
	str := b.judge()
	if str != "x" {
		t.Error("Test05 is failed")
	}
	b.show()
	fmt.Print("\n")
}
