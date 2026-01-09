# Developing New Plugins

HawkLens is built to be extensible. If you want to add support for a new social platform or data source, follow this guide.

## 🧩 The Plugin Interface

Every plugin must implement the `Plugin` interface defined in `pkg/plugins/plugin.go`:

```go
type Plugin interface {
	Name() string
	Description() string
	Fetch(ctx context.Context, query string) ([]Result, error)
}
```

## 🛠️ Steps to Create a Plugin

### 1. Create a new package
Create a new directory in `internal/plugins/[platform]`.

### 2. Implement the interface
Create a struct that satisfies the `Plugin` interface.

```go
package myplatform

import (
	"context"
	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

type MyPlugin struct{}

func (p *MyPlugin) Name() string { return "myplatform" }
func (p *MyPlugin) Description() string { return "Scrapes data from MyPlatform" }

func (p *MyPlugin) Fetch(ctx context.Context, query string) ([]plugins.Result, error) {
	// 1. Authenticate (optional)
	// 2. Fetch data
	// 3. Map to plugins.Result
	return []plugins.Result{...}, nil
}
```

### 3. Register your plugin
In your plugin's package, add an `init()` function or register it in `internal/cli/init.go`.

```go
import "github.com/ismailtsdln/HawkLens/pkg/plugins"

func init() {
    plugins.Register(&MyPlugin{})
}
```

## 📝 Result Mapping Guidelines

When mapping your platform's data to the `plugins.Result` struct, aim for consistency:

- **Platform**: Use lowercase name (e.g., `twitter`).
- **DataType**: Specify the content type (e.g., `tweet`, `video`, `post`).
- **Data**: A map containing the raw but relevant fields.
  - For text content, use the key `text`, `title`, or `caption`.

## 🧪 Testing Your Plugin

Create a `[platform]_test.go` file in your package. Use the `plugins.GetPlugin()` function to retrieve and verify your implementation.
