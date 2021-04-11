package racer

import "testing"

func TestRacer(t *testing.T) {
	fastURL := "http://www.baidu.com"
	slowURL := "http://www.bilibili.com/"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
