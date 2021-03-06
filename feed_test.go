package fugoblog

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	tc, err := getTime()
	if err != nil {
		t.Fatalf("could not laod location")
	}
	if time.Now().Before(tc) {
		t.Errorf("should not be newer time")
	}
}

func TestCheckIfRSSUpdated(t *testing.T) {
	_, err := CheckIfRSSUpdated()
	if err != nil {
		t.Errorf("could not retrieve RSS")
	}
}
