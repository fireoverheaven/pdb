package pdb

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/asdine/storm/v3"
)

func MergeDB(dbfn1 string, dbfn2 string) {
	db1, err := bolt.Open(dbfn1, 0600, nil)
	if err != nil {
		log.Fatalf("Couldn't open storm db at '%s'", dbfn1)
	}

	db2, err := bolt.Open(dbfn2, 0600, nil)
	if err != nil {
		log.Fatalf("Couldn't open storm db at '%s'", dbfn2)
	}

  copyBucket(db1, db2, "FileMetadata")
  defer db1.Close()
  defer db2.Close()

}


func copyBucket(idb, odb *bolt.DB, bucket string) error {
	return idb.View(func(itx *bolt.Tx) error {
		ib := itx.Bucket([]byte(bucket))
		return odb.Update(func(otx *bolt.Tx) error {
			ob, err := otx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
			return ib.ForEach(func(k, v []byte) error {
				return ob.Put(k, v)
			})
		})
	})
}
