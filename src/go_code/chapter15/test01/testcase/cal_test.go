package main

import (
	"testing"
)

// Test Case
func TestGetSum(t *testing.T) {
	// call
	res := getSum(10)
	if res != 55 {
		// fmt.Println("Test Fails, Expect 55, Received:", res)
		t.Fatalf("Test Fails, Expect: 55, Received: %v", res)
	} else {
		// fmt.Println("Test Pass")
		t.Log("Test Pass")
	}
}

