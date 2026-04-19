package storage

import (
	"errors"
	"time"
)

// Store defines the interface for the storage backend.
type Store interface {
	// Get retrieves the value for the given key.
	// Returns ErrNotFound if the key doesn't exist.
	Get(key string) ([]byte, error)

	// Set stores the value for the given key.
	Set(key string, value []byte) error

	// SetWithTTL stores the value for the given key with an expiration time.
	SetWithTTL(key string, value []byte, ttl time.Duration) error

	// Delete removes the key from the store.
	Delete(key string) error

	// Has checks whether the key exists in the store
	Has(key string) (bool, error)

	// List returns all keys.
	List() ([]string, error)

	// ListFiltered returns all keys that start with the given prefix.
	ListFiltered(prefix string) ([]string, error)

	// Close releases all resources held by the store.
	Close() error
}

// ErrNotFound is returned when a requested key doesn't exist.
var ErrNotFound = errors.New("key not found")
