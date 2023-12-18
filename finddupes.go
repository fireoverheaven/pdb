package pdb

import (
  "log"
  
  "github.com/asdine/storm/v3"
)

func InitDB(filename string) *storm.DB {
  log.Printf("Opening db at '%s'", filename)
  db, err := storm.Open(filename)
  if err != nil {
    log.Fatalf("Couldn't open storm db at '%s'", filename)
  }

  log.Printf("Initializing db")
  err = db.Init(&FileMetadata{})
  if err != nil {
    log.Fatalf("Couldn't initialize db ('%s') as FileMetadata", filename)
  }

  return db
}

func ReIndexDB(db *storm.DB) error {
  log.Printf("Reindexing %s", db)
  err := db.ReIndex(&FileMetadata{})
  if err != nil {
    return err
  }

  log.Printf("Committing db")
  err = db.Commit()
  if err != nil {
    return err
  }

  return nil
}

func PrintBySize(db *storm.DB) {
  var alephrecords []FileMetadata
  var daisyrecords []FileMetadata
  err := db.Find("Host", "aleph.fireoh.com", &alephrecords)
  if err != nil {
    log.Fatalf("Couldn't get all records by size: %s", err)
  }
  err = db.Find("Host", "daisy.fireoh.com", &daisyrecords)
  if err != nil {
    log.Fatalf("Couldn't get all records by size: %s", err)
  }
  for _, record := range alephrecords {
    if record.Size < 1382740 {
      continue
    }
    for _, r2 := range daisyrecords {
      if r2.Size == record.Size {
        log.Printf("%d\n%s\n%s\n\n", record.Size, record.Path, r2.Path)
      }
    }
  }
}
