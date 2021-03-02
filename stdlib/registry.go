package stdlib

import "github.com/regeda/expr/delegate"

var registry map[string]delegate.Delegator

// Register adds a delegate to the registry.
func Register(s string, f delegate.Delegator) {
	if registry == nil {
		registry = make(map[string]delegate.Delegator)
	}
	registry[s] = f
}

// Registry returns the registry of delegators.
func Registry() map[string]delegate.Delegator {
	return registry
}
