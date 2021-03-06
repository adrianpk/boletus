// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package here

import (
	"sync"
)

// infoMap wraps sync.Map and uses the following types:
// key:   string
// value: Info
type infoMap struct {
	data *sync.Map
}

// Load the key from the map.
// Returns Info or bool.
// A false return indicates either the key was not found
// or the value is not of type Info
func (m *infoMap) Load(key string) (Info, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return Info{}, false
	}
	s, ok := i.(Info)
	return s, ok
}

// Store a Info in the map
func (m *infoMap) Store(key string, value Info) {
	m.data.Store(key, value)
}
