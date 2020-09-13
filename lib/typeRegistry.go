package lib

import "reflect"

type TypeRegistry struct {
	registered map[string]reflect.Type
}

func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{registered: map[string]reflect.Type{}}
}

func (registry *TypeRegistry) Register(t reflect.Type) {
	registry.registered[t.Name()] = t
}

func (registry *TypeRegistry) Has(name string) bool {
	_, ok := registry.registered[name]
	return ok
}

func (registry *TypeRegistry) Get(name string) reflect.Type {
	return registry.registered[name]
}
