package storage

import (
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

const db_name = "templatefactory.db"
const bucket_name = "templatefactory_main.bucket"
const db_permissions = 0644

var db_path = filepath.Join(TF_HOME, db_name)

func save(templateName, encodedTemplate string) error {
	db, err := bolt.Open(db_path, db_permissions, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	defer db.Close()

	key := []byte(templateName)
	value := []byte(encodedTemplate)

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		if err != nil {
			return err
		}

		return bucket.Put(key, value)

	})

	return err
}

func load(templateName string) (encodedTemplate string, err error) {
	db, err := bolt.Open(db_path, db_permissions, &bolt.Options{
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

		encodedTemplate = string(bucket.Get([]byte(templateName)))
		return nil
	})

	if err != nil {
		return "", err
	}

	return encodedTemplate, nil
}
