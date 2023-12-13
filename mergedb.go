package pdb

import (
	"log"

	"github.com/asdine/storm/v3"
)

func MergeDB(dbfn1 string, dbfn2 string) {
	db1, err := storm.Open(dbfn1)
	if err != nil {
		log.Fatalf("Couldn't open storm db at '%s'", dbfn1)
	}

	db2, err := storm.Open(dbfn2)
	if err != nil {
		log.Fatalf("Couldn't open storm db at '%s'", dbfn2)
	}

	db1.Init(&FileMetadata{})
	db2.Init(&FileMetadata{})

	var md1 []FileMetadata
	db1.All(&md1)
	for x := range md1 {
		db2.Save(&x)
	}

}
