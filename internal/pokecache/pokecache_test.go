package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	c := NewCache(5 * time.Second)

	key := "abc"
	val := []byte("123")

	c.Add(key, val)

	got, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected key %q to exist", key)
	}
	if string(got) != string(val) {
		t.Fatalf("expected %q, got %q", string(val), string(got))
	}
}

func TestCacheGetMissing(t *testing.T) {
	c := NewCache(5 * time.Second)

	got, ok := c.Get("missing")
	if ok {
		t.Fatalf("expected key to be missing, got ok=true with val=%v", got)
	}
}

func TestCacheReapExpired(t *testing.T) {
	interval := 50 * time.Millisecond
	c := NewCache(interval)

	key := "expired"
	val := []byte("xxx")

	c.Add(key, val)

	time.Sleep(2 * interval)

	got, ok := c.Get(key)
	if ok {
		t.Fatalf("expected key %q to be reaped, got value %q", key, string(got))
	}
}

func TestCacheNotReapedBeforeInterval(t *testing.T) {
	interval := 200 * time.Millisecond
	c := NewCache(interval)

	key := "alive"
	val := []byte("yyy")

	c.Add(key, val)

	time.Sleep(interval / 2)

	got, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected key %q to exist", key)
	}
	if string(got) != string(val) {
		t.Fatalf("expected %q, got %q", string(val), string(got))
	}
}
