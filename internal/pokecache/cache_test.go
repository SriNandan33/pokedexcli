package pokecache

import "testing"

func TestAddGet(t *testing.T) {
	cache := NewCache()
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
