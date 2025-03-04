package concurrency

import (
	"reflect"
	"testing"
	"time"
)

const unavailableSite = "waat:/fusdfs.dasf"

func slowWebsiteChecker(url string) bool {
	time.Sleep(5 * time.Millisecond)

	return mockWebsiteChecker(url)
}

func mockWebsiteChecker(url string) bool {
	return url != unavailableSite
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range len(urls) {
		urls[i] = "a URL"
	}
	b.ResetTimer()
	for b.Loop() {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"htps://www.google.com",
		"https://ack.com",
		unavailableSite,
	}
	want := map[string]bool{
		"htps://www.google.com": true,
		"https://ack.com":       true,
		unavailableSite:         false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
