package cache

import (
	"fmt"
	"testing"
	"time"
)

var cases = []struct {
	key   string
	value []byte
}{
	{key: "https://test.com", value: []byte("data-1")},
	{key: "https://test.com/test-path", value: []byte("data-2")},
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := New(interval)
			cache.Add(c.key, c.value)
			data, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(data) != string(c.value) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const wait = interval * 2
	cache := New(interval)
	testCase := cases[0]
	cache.Add(testCase.key, testCase.value)

	_, ok := cache.Get(testCase.key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(wait)

	_, ok = cache.Get(testCase.key)
	if ok {
		t.Errorf("expected to find key")
		return
	}
}
