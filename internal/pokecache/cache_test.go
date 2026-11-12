package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(2 * time.Second)
	cache.Add("first", []byte("first value"))
	cache.Add("second", []byte("second value"))
	cache.Add("first", []byte("updated first value"))

	type test struct {
		key   string
		want  string
		found bool
	}

	cases := []test{
		{key: "first", want: "updated first value", found: true},
		{key: "second", want: "second value", found: true},
		{key: "notFound", found: false},
	}

	for _, c := range cases {
		t.Run(c.key, func(t *testing.T) {
			got, ok := cache.Get(c.key)
			if ok != c.found || string(got) != c.want {
				t.Fatalf("Get(%q) = %q, %v; Want: %q, %v", c.key, got, ok, c.want, c.found)
			}
		})
	}
}

func TestReaping(t *testing.T) {
	cache := NewCache(2 * time.Second)
	cache.Add("first", []byte("first value"))
	cache.Add("second", []byte("second value"))

	type test struct {
		key   string
		want  string
		found bool
	}

	cases := []test{
		{key: "first", want: "first value", found: true},
		{key: "second", want: "second value", found: true},
	}

	time.Sleep(2 * time.Second)
	for _, c := range cases {
		t.Run(fmt.Sprintf("Expiry for %s", c.key), func(t *testing.T) {
			_, ok := cache.Get(c.key)
			if ok {
				t.Fatalf("Expected cache entry to be reaped")
			}
		})
	}
}
