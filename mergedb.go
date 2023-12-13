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

  err = db1.Init(&FileMetadata{})
  if err != nil {
    log.Printf("Could not init db1")
  }
	db2.Init(&FileMetadata{})
  if err != nil {
    log.Printf("Could not init db2")
  }

	var md1 []FileMetadata
  err = db1.All(&md1)
  defer db1.Close()
  if err != nil {
    log.Printf("Could not get all from db1")
  }
	for x := range md1 {
    log.Printf("Moving %d", &x)
    err := db2.Save(&md1[x])
    if err != nil {
      log.Printf("Could not save '%s' to db", &md1[x])
    }
	}

  defer db2.Close()
}
