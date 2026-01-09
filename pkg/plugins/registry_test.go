package plugins

import (
	"context"
	"testing"
)

type MockPlugin struct {
	name string
}

func (m *MockPlugin) Name() string        { return m.name }
func (m *MockPlugin) Description() string { return "mock description" }
func (m *MockPlugin) Fetch(ctx context.Context, query string) ([]Result, error) {
	return []Result{{Platform: m.name, DataType: "mock"}}, nil
}

func TestRegistry(t *testing.T) {
	p := &MockPlugin{name: "test-plugin"}
	Register(p)

	// Test ListPlugins
	list := ListPlugins()
	found := false
	for _, name := range list {
		if name == "test-plugin" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Register failed: test-plugin not in ListPlugins")
	}

	// Test GetPlugin success
	retrieved, err := GetPlugin("test-plugin")
	if err != nil {
		t.Errorf("GetPlugin failed: %v", err)
	}
	if retrieved.Name() != "test-plugin" {
		t.Errorf("Expected test-plugin, got %s", retrieved.Name())
	}

	// Test GetPlugin failure
	_, err = GetPlugin("non-existent")
	if err == nil {
		t.Error("GetPlugin should fail for non-existent plugin")
	}
}
