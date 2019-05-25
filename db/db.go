package db

import (
	"time"

    "github.com/boltdb/bolt"
)

// create name key for out bucket
var authBucket = []byte("users")
var doTokenBucket = []byte("tokens")
var db *bolt.DB

// User is the struct
// that hold Digital ocean api key
// and the JWT to denote if user is logged in or not
type User struct {
    Key string 
    Value string
}

// InitDB takes in dbPath as a string and opens database connection
// also sets up relevant db buckets if they do not exist
func InitDB(dbPath string) error {
    // opens connection to database with timeout of 1 second
    db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        return err
    }
    // will update DB with callback, allows for buckets to be setup
    return db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists(authBucket)
        // will be nil if no error
        return err
    })
}

