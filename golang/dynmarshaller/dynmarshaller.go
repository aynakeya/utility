package dynmarshaller

import (
	"encoding/json"
	"errors"
	"reflect"
)

type DynamicMarshallable struct {
	Type string
}

type Initilizable interface {
	Initialize() error
}

type DynamicUnmarshaller[T any] struct {
	registry map[string]T
}

func NewDynamicUnmarshaller[T any](registry map[string]T) *DynamicUnmarshaller[T] {
	return &DynamicUnmarshaller[T]{
		registry: registry,
	}
}

func (d *DynamicUnmarshaller[T]) Unmarshal(data []byte) (T, error) {
	var f DynamicMarshallable
	if err := json.Unmarshal(data, &f); err != nil {
		return *new(T), err
	}
	t, ok := d.registry[f.Type]
	if !ok {
		return *new(T), errors.New("registry not found")
	}
	v := reflect.New(reflect.TypeOf(t).Elem()).Interface()
	err := json.Unmarshal(data, v)
	if err != nil {
		return *new(T), err
	}
	if i, ok := v.(Initilizable); ok {
		return v.(T), i.Initialize()
	}
	return v.(T), nil
}
