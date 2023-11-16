package storage

import (
	"time"

	"github.com/boltdb/bolt"
)

const bucket_name = "templatefactory_main.bucket"

func save(key, encodedString string) error {
	db, err := bolt.Open(db_path, db_file_permissions, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	defer db.Close()

	value := []byte(encodedString)

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(key), value)

	})

	return err
}

func getAllKeys() ([]string, error) {
	db, err := bolt.Open(db_path, db_file_permissions, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	keys := []string{}
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))
		cursor := bucket.Cursor()

		for key, _ := cursor.First(); key != nil; key, _ = cursor.Next() {
			keys = append(keys, string(key))
		}

		return nil
	})

	return keys, err
}

func delete(key string) error {
	db, err := bolt.Open(db_path, db_file_permissions, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))

		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		return bucket.Delete([]byte(key))
	})

	return err
}

func load(key string) (encodedString string, err error) {
	db, err := bolt.Open(db_path, db_file_permissions, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return "", err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))

		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		encodedString = string(bucket.Get([]byte(key)))
		return nil
	})

	if err != nil {
		return "", err
	}

	return encodedString, nil
}
