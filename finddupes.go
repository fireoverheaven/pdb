package pdb

import (
  "log"
  
  "github.com/asdine/storm/v3"
)

func InitDB(filename string) *storm.DB {
  db, err := storm.Open(filename)
  if err != nil {
    log.Fatalf("Couldn't open storm db at '%s'", filename)
  }
  err = db.Init(&FileMetadata{})
  if err != nil {
    log.Fatalf("Couldn't initialize db ('%s') as FileMetadata", filename)
  }

  return db
}

func ReindexDB(db *storm.DB) {
  err := db.ReIndex(&FileMetadata{})
  if err != nil {
    log.Fatalf("Couldn't reindex: %s", err)
  }
  err = db.Commit()
  if err != nil {
    log.Fatalf("Couldn't commit: %s", err)
  }
}

func PrintBySize(db *storm.DB) {
  var records []FileMetadata
  err := db.AllByIndex("Size", &records)
  if err != nil {
    log.Fatalf("Couldn't get all records by size: %s", err)
  }
  for _, record := range records {
    log.Printf("%d %s", record.Size, record.Path)
  }
}
