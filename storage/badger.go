package storage

import (
	"fmt"
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

// BadgerStore implements Store using BadgerDB.
type BadgerStore struct {
	db *badger.DB
}

// BadgerOptions holds the configuration for a Badger instance.
type BadgerOptions struct {
	// Dir is the directory where Badger stores its data files.
	Dir string

	// ValueDir is the directory for value log files.
	// If empty, defaults to Dir.
	ValueDir string

	// InMemory runs Badger in pure in-memory mode.
	InMemory bool
}

// NewBadgerStore creates a new Badger-backed Store.
func NewBadgerStore(opts BadgerOptions) (*BadgerStore, error) {
	options := badger.DefaultOptions(opts.Dir)

	if opts.ValueDir != "" {
		options = options.WithValueDir(opts.ValueDir)
	}

	if opts.InMemory {
		options = options.WithInMemory(true)
	}

	db, err := badger.Open(options)
	if err != nil {
		return nil, fmt.Errorf("open badger: %w", err)
	}

	return &BadgerStore{db: db}, nil
}

// Get retrieves the value for the given key.
// Returns ErrNotFound if the key doesn't exist.
func (s *BadgerStore) Get(key string) ([]byte, error) {
	var value []byte

	err := s.db.View(
		func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(key))
			if err != nil {
				if err == badger.ErrKeyNotFound {
					return ErrNotFound
				}

				return fmt.Errorf("get: %w", err)
			}

			value, err = item.ValueCopy(nil)
			if err != nil {
				return fmt.Errorf("value copy: %w", err)
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return value, nil
}

// Set stores the value for the given key.
func (s *BadgerStore) Set(key string, value []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		if err != nil {
			return fmt.Errorf("set: %w", err)
		}

		return nil
	})
}

// SetWithTTL stores the value for the given key with an expiration time.
func (s *BadgerStore) SetWithTTL(key string, value []byte, ttl time.Duration) error {
	return s.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), value).WithTTL(ttl)
		err := txn.SetEntry(entry)
		if err != nil {
			return fmt.Errorf("set with ttl: %w", err)
		}

		return nil
	})
}

// Delete removes the key from the store.
func (s *BadgerStore) Delete(key string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		if err != nil {
			return fmt.Errorf("delete: %w", err)
		}

		return nil
	})
}

// Has checks whether the key exists in the store.
func (s *BadgerStore) Has(key string) (bool, error) {
	exists := false

	err := s.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				exists = false

				return nil
			}

			return fmt.Errorf("has: %w", err)
		}

		exists = true

		return nil
	})

	if err != nil {
		return false, err
	}

	return exists, nil
}

// List returns all keys.
func (s *BadgerStore) List() ([]string, error) {
	var keys []string

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			keys = append(keys, string(item.Key()))
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("keys: %w", err)
	}

	return keys, nil
}

// ListFiltered returns all keys that start with the given prefix.
func (s *BadgerStore) ListFiltered(prefix string) ([]string, error) {
	var keys []string

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		it := txn.NewIterator(opts)
		defer it.Close()

		prefixBytes := []byte(prefix)
		for it.Seek(prefixBytes); it.ValidForPrefix(prefixBytes); it.Next() {
			item := it.Item()
			keys = append(keys, string(item.Key()))
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("keys: %w", err)
	}

	return keys, nil
}

// Close release all resources held by the store.
func (s *BadgerStore) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("close badger: %w", err)
	}

	return nil
}
