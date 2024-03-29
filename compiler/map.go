package compiler

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Link represents a node of doubly linked list.
type Link struct {
	key   string
	value any
	next  *Link
	prev  *Link
}

// Map holds the elements in go's native map, also maintains the head and tail link
// to keep the elements in insertion order.
type Map struct {
	m    map[string]*Link
	head *Link
	tail *Link
}

// newLink creates a new link node with the provided key and value.
func newLink(key string, value any) *Link {
	return &Link{key: key, value: value, next: nil, prev: nil}
}

// NewMap instantiates a linked hash map.
func NewMap() *Map {
	return &Map{m: make(map[string]*Link), head: nil, tail: nil}
}

// ConvertMap converts the input to a linked hash map.
func ConvertMap(in any) any {
	switch v := in.(type) {
	case map[string]any:
		out := NewMap()
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			switch val := v[key].(type) {
			case json.Number:
				if n, err := val.Int64(); err == nil {
					v[key] = n
				} else if f, err := val.Float64(); err == nil {
					v[key] = f
				}
			}
			out.Set(key, ConvertMap(v[key]))
		}
		return out
	case []any:
		for i, item := range v {
			v[i] = ConvertMap(item)
		}
	}
	return in
}

// LoadMap instantiates a linked hash map and initializing it from map[string]any.
func LoadMap(init map[string]any) (ret *Map) {
	ret = NewMap()
	keys := make([]string, 0, len(init))
	for key := range init {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, v := range keys {
		ret.Set(v, ConvertMap(init[v]))
	}
	return
}

// Set inserts an element into the map.
func (m *Map) Set(key string, value any) {
	link, found := m.m[key]
	if !found {
		link = newLink(key, value)
		if m.tail == nil {
			m.head = link
			m.tail = link
		} else {
			m.tail.next = link
			link.prev = m.tail
			m.tail = link
		}
		m.m[key] = link
	} else {
		link.value = value
	}
}

// Get searches the element in the map by key and returns its value or nil if key doesn't exists.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key string) (value any, found bool) {
	var link *Link
	link, found = m.m[key]
	if found {
		value = link.value
	} else {
		value = nil
	}
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key string) {
	link, found := m.m[key]
	if found {
		delete(m.m, key)
		if m.head == link && m.tail == link {
			m.head = nil
			m.tail = nil
		} else if m.tail == link {
			m.tail = link.prev
			link.prev.next = nil
		} else if m.head == link {
			m.head = link.next
			link.next.prev = nil
		} else {
			link.prev.next = link.next
			link.next.prev = link.prev
		}
	}
}

// IsEmpty returns true if map does not contain any elements.
func (m *Map) IsEmpty() bool {
	return m == nil || m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Keys returns all keys of the map (insertion order).
func (m *Map) Keys() []string {
	keys := make([]string, m.Size())
	count := 0
	for current := m.head; current != nil; current = current.next {
		keys[count] = current.key
		count++
	}
	return keys
}

// Values returns all values of the map (insertion order).
func (m *Map) Values() []any {
	values := make([]any, m.Size())
	count := 0
	for current := m.head; current != nil; current = current.next {
		values[count] = current.value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.m = make(map[string]*Link)
	m.head = nil
	m.tail = nil
}

// String returns a string representation of container.
func (m *Map) String() string {
	str := "map["
	for current := m.head; current != nil; current = current.next {
		str += fmt.Sprintf("%v:%v ", current.key, current.value)
	}
	return strings.TrimRight(str, " ") + "]"
}

// MarshalJSON marshals the Map into JSON.
func (m *Map) MarshalJSON() ([]byte, error) {
	s := "{"
	for current := m.head; current != nil; current = current.next {
		k := current.key
		escaped := strings.Replace(k, `"`, `\"`, -1)
		s = s + `"` + escaped + `":`
		v := current.value
		vBytes, err := json.Marshal(v)
		if err != nil {
			return []byte{}, err
		}
		s = s + string(vBytes) + ","
	}
	if len(s) > 1 {
		s = s[0 : len(s)-1]
	}
	s = s + "}"
	return []byte(s), nil
}

// NewFile creates a new file with default values.
func NewFile() *Map {
	return LoadMap(map[string]any{
		"Name":     "",
		"MimeType": "",
		"Body":     []byte{},
	})
}

// NewFileFromMap creates a new file from a map.
func NewFileFromMap(m map[string]any) (f *Map, ok bool) {
	var v any
	f = NewFile()

	if v, ok = m["Name"].(string); !ok {
		return
	}
	f.Set("Name", v)
	if v, ok = m["MimeType"].(string); !ok {
		return
	}
	f.Set("MimeType", v)
	if v, ok = m["Body"].([]byte); !ok {
		return
	}
	f.Set("Body", v)
	return
}
