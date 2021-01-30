package main

import (
	"testing"
)

// Test Case
func TestGetDiff(t *testing.T) {
	// call
	res := getDiff(10, 3)
	if res != 7 {
		// fmt.Println("Test Fails, Expect 55, Received:", res)
		t.Fatalf("Test Fails, Expect: 7, Received: %v", res)
	} else {
		// fmt.Println("Test Pass")
		t.Log("Test Pass")
	}
}

