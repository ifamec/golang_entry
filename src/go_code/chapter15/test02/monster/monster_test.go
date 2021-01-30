package monster

import (
	"os"
	"testing"
)

var testFilePath = "./store.txt"

func TestMonster_Store(t *testing.T) {
	var m *Monster = &Monster{"AAA", 20, "Climbing"}
	var status bool = (*m).Store(testFilePath)
	if !status {
		t.Fatalf("Test Monster Store() Fails\n")
	}
	t.Logf("Test Monster Store() Success")
}

func TestMonster_ReStore(t *testing.T) {
	var m *Monster = &Monster{}
	var status bool = (*m).ReStore(testFilePath)

	if !status {
		t.Fatalf("Test Monster ReStore() Fails\n")
	}
	if (*m).Name != "AAA" {
		t.Fatalf("Test Monster ReStore() Object.Name Expect: %v, Received: %v\n", "AAA", (*m).Name)
	}

	t.Logf("Test Monster ReStore() Success")
}

func TestMDeleteStoreFile(t *testing.T) {
	err := os.Remove(testFilePath)

	if err != nil {
		t.Fatalf("Remove Test File Fails\n")
	}
	t.Logf("Remove Test File Success")
}

