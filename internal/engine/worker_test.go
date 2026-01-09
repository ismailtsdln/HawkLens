package engine

import (
	"context"
	"testing"
	"time"

	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

type MockEnginePlugin struct {
	name string
}

func (m *MockEnginePlugin) Name() string        { return m.name }
func (m *MockEnginePlugin) Description() string { return "engine mock description" }
func (m *MockEnginePlugin) Fetch(ctx context.Context, query string) ([]plugins.Result, error) {
	return []plugins.Result{{Platform: m.name, DataType: "test"}}, nil
}

func TestDispatcher(t *testing.T) {
	// Register mock plugin
	plugins.Register(&MockEnginePlugin{name: "engine-mock"})

	dispatcher := NewDispatcher(2)
	ctx := context.Background()
	dispatcher.Run(ctx)

	dispatcher.Submit("engine-mock", "test-query")

	// Collect result with timeout
	select {
	case wrapper := <-dispatcher.Results():
		if wrapper.Error != nil {
			t.Errorf("Job failed: %v", wrapper.Error)
		}
		if wrapper.Platform != "engine-mock" {
			t.Errorf("Expected engine-mock, got %s", wrapper.Platform)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for dispatcher results")
	}

	dispatcher.Wait()
}
