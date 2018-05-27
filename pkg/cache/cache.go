// Package cache provides a library for abstracting the caching layer.
package cache

// Cache is an interface describing the cache layer for storing computed values.
type Cache interface {
	Set(KeyValuePair) error
	Get(Key) (*KeyValuePair, error)
}

// Key represents a key in a cache.
type Key = int64

// Value represents a value in a cache.
type Value = string

// KeyValuePair represents a key-value pair in a cache.
type KeyValuePair struct {
	Key   Key
	Value Value
}
