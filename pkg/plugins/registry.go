package plugins

import (
	"sync"
)

var (
	registry = make(map[string]Plugin)
	mu       sync.RWMutex
)

// Register adds a plugin to the global registry
func Register(p Plugin) {
	mu.Lock()
	defer mu.Unlock()
	registry[p.Name()] = p
}

// GetPlugin retrieves a plugin by name
func GetPlugin(name string) (Plugin, bool) {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := registry[name]
	return p, ok
}

// ListPlugins returns names of all registered plugins
func ListPlugins() []string {
	mu.RLock()
	defer mu.RUnlock()
	var list []string
	for name := range registry {
		list = append(list, name)
	}
	return list
}
