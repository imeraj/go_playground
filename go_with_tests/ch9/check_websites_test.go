package concurrency

import (
	"testing"
	"reflect"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.abc.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://blog.abc.com": true,
		"waat://furhurterwe.geds": false,
	}

	got := checkWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected %v, got %v", want, got)
	}
}