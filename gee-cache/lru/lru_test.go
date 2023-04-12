package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	lru := New(0, nil) // maxBytes 0 表示无内存限制
	lru.Add("key1", String("123"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "123" {
		t.Fatalf("cache hit key1=123 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestCache_UseRemoveOldElement(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "v1", "v2", "value3"
	testCap := len(k1 + k2 + v1 + v2)
	lru := New(int64(testCap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get(k1); ok || lru.Len() == 2 {
		t.Fatalf("RemoveOldElement %s failed", k1)
	}
}

func TestCache_UseOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(4, callback)
	lru.Add("k1", String("v1"))
	lru.Add("k2", String("v2"))
	lru.Add("k3", String("v3"))
	lru.Add("k4", String("v4"))

	expect := []string{"k1", "k2", "k3"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed keys %s , expect keys equals to %s", keys, expect)
	}
}
