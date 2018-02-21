package nem

import (
    "testing"
)

func TestGetHeartbeat(t *testing.T) {
    result, err := GetHeartbeat()
    if err != nil {
        t.Fatalf("failed test %#v", err)
    }
    if result.Message != "ok" {
        t.Fatal("failed test")
    }
}
