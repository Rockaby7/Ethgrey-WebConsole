package storage

import (
	"errors"
	"log"
	"time"

	"github.com/dgraph-io/badger/v3"
)

type storage struct {
	db *badger.DB
}

var Storage *storage

func init() {
	open, err := badger.Open(badger.DefaultOptions("./agent_storage"))
	if err != nil {
		log.Fatalln(err)
	}

	Storage = &storage{db: open}
}

func (s *storage) GetDB() *badger.DB {
	return s.db
}

func (s *storage) SetNX(key string, value []byte, timeout time.Duration) error {
	return s.db.Update(func(txn *badger.Txn) error {
		ttl := badger.NewEntry([]byte(key), value)
		if timeout > 0 {
			ttl = ttl.WithTTL(timeout)
		}
		return txn.SetEntry(ttl)
	})
}

func (s *storage) Get(key string) (value []byte, err error) {
	return value, s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			if val == nil {
				return errors.New("val is nil")
			}

			value = val
			return nil
		})
	})
}

func (s *storage) Del(key string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (s *storage) Prefix(prefix string) (map[string][]byte, error) {
	result := make(map[string][]byte)

	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek([]byte(prefix)); it.ValidForPrefix([]byte(prefix)); it.Next() {
			item := it.Item()
			k := item.Key()
			item.Value(func(val []byte) error {
				result[string(k)] = val
				return nil
			})
		}

		return nil
	})

	return result, err
}
