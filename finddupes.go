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
  var records []FileMetadata
  err := db.Find("Host", "aleph.fireoh.com", &records)
  if err != nil {
    log.Fatalf("Couldn't get all records by size: %s", err)
  }
  for _, record := range records {
    db.Find("Size", record.Size, &FileMetadata{})
    log.Printf("%d %s", record.Size, record.Path)
  }
}
