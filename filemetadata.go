package pdb

import (
	"github.com/google/uuid"
	"time"
)

type FileMetadata struct {
	UUID      uuid.UUID `storm:"id"`
	Host      string
	Filename  string    `storm:"index"`
	ParentDir string    `storm:"index"`
	Path      string    `storm:"index"`
	Size      int64     `storm:"index"`
	Mime      string    `storm:"index"`
	LastScan  time.Time `storm:"index"`
}
