package pokecache

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCacheAdd(t *testing.T) {
	cache := NewCache()

	val := []byte("test-value")
	cache.Add("test", val)

	expected := []byte("test-value")
	got, ok := cache.Get("test")
	if !ok {
		t.Fatal("expected key to be present")
	}
	if diff := cmp.Diff(got, expected); diff != "" {
		t.Fatalf("expected: %s, got: %s", expected, got)
	}
}

func TestCacheAddSameKeyTwice(t *testing.T) {
	cache := NewCache()
	firstVal := []byte("test-value")
	cache.Add("test", firstVal)

	secondVal := []byte("updated-value")
	cache.Add("test", secondVal)

	expected := []byte("updated-value")
	got, ok := cache.Get("test")
	if !ok {
		t.Fatal("expected key to be present")
	}
	if diff := cmp.Diff(got, expected); diff != "" {
		t.Fatalf("expected: %s, got: %s", expected, got)
	}
}

func TestCacheGetMissingKey(t *testing.T) {
	cache := NewCache()

	_, ok := cache.Get("test")

	if ok {
		t.Fatalf("expected: %v, got: %v", false, ok)
	}
}

func TestAddTwoDifferentKeys(t *testing.T) {
	cache := NewCache()
	firstVal := []byte("test-one")
	cache.Add("testOne", firstVal)

	secondVal := []byte("test-two")
	cache.Add("testTwo", secondVal)

	first, ok := cache.Get("testOne")
	if !ok {
		t.Fatal("expected key to be present")
	}
	second, ok := cache.Get("testTwo")
	if !ok {
		t.Fatal("expected key to be present")
	}

	firstDiff := cmp.Diff(firstVal, first)
	secondDiff := cmp.Diff(secondVal, second)
	if firstDiff != "" || secondDiff != "" {
		t.Fatal()
	}
}
