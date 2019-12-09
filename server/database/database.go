package database

import (
	"log"
	"os"
	"time"
	"github.com/boltdb/bolt"
)

var dbName = "DATABASE.db"
var db *bolt.DB
func Start(str string) {
	var err error
	dbName = str
	db, err = bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func Stop(){
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func Init(str string) {
	if _,err := os.Open(str)  ; err == nil{
		log.Println("database is already exist . If you want to initialze it , please delete it and try again")
		return
	}
	Start(str)
	if err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("users"))
		tx.CreateBucket([]byte("films"))
		tx.CreateBucket([]byte("people"))
		tx.CreateBucket([]byte("planets"))
		tx.CreateBucket([]byte("species"))
		tx.CreateBucket([]byte("starships"))
		tx.CreateBucket([]byte("vehicles"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	Stop()
}


func Update(bucketName []byte, key []byte, value []byte) {
	if err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(bucketName).Put(key, value); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func GetValue(bucketName []byte, key []byte) string {
	var result []byte
	if err := db.View(func(tx *bolt.Tx) error {
		byteLen := len(tx.Bucket([]byte(bucketName)).Get(key))
		result = make([]byte, byteLen)
		copy(result[:], tx.Bucket([]byte(bucketName)).Get(key)[:])
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func CheckKeyExist(bucketName []byte, key []byte) bool {
	var byteLen int
	if err := db.View(func(tx *bolt.Tx) error {
		byteLen = len(tx.Bucket([]byte(bucketName)).Get(key))
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return (byteLen != 0)
}

func GetBucketCount(bucketName []byte) (int) {
	count := 0
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			count ++
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return count
}

func CheckBucket(bucketName []byte) {
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}



