package components

import (
	"fmt"
	"log"
)

var (
	Registry = make(map[string]Component)
)

type Component interface {
	Greet() string
	UnmarshalSettings([]byte) Component
}

// GetComponent returns the Component specified by name from `Registry`.
func GetComponent(kind string) (Component, error) {
	// check if exists
	if _, ok := Registry[kind]; ok {
		return Registry[kind], nil
	}
	return nil, fmt.Errorf("%s is not a registered Component type", kind)
}

// Register is called by the `init` function of every `Component` to add
// the caller to the global `Registry` map. If the caller attempts to
// add a `Component` to the registry using the same name as a prior
// `Component` the call will log an error and exit.
func Register(kind string, c Component) {
	// check for name collision before adding
	if _, ok := Registry[kind]; ok {
		log.Fatalf("Component: %s has already been added to the registry", kind)
	}
	Registry[kind] = c
}
