package pdb

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

func ScanDir(dbfilename string, startdir string) {
	db, err := storm.Open(dbfilename)
	if err != nil {
		log.Fatalf("Couldn't open storm db at '%s'", dbfilename)
	}
	err = db.Init(&FileMetadata{})
	if err != nil {
		log.Fatalf("Couldn't initialize db ('%s') as FileMetadata", dbfilename)
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Couldn't get hostname.")
	}

	STARTDIR, err := filepath.Abs(startdir)
	if err != nil {
		log.Fatalf("Couldn't resolve %s to absolute path.", STARTDIR)
	}
	fmt.Printf("Beginning scan of '%s'...", STARTDIR)

	var scanned_dirs []string
	filepath.Walk(STARTDIR,
		func(mypath string, info os.FileInfo, err error) error {
			log.Printf("Walking...\t%s", mypath)
			if err != nil {
				log.Printf("Failed to walk path '%s' (%s)", mypath, err)
				return err
			}
			if info.IsDir() {
				scanned_dirs = append(scanned_dirs, mypath)
				return nil
			}

			ppath, filename := path.Split(mypath)
			parent := path.Base(path.Dir(ppath))
			mtype, err := mimetype.DetectFile(mypath)
			if err != nil {
				log.Printf("Failed to read mime type at '%s' (%s)", mypath, err)
				return err
			}

			file_metadata := FileMetadata{
				UUID:      uuid.New(),
				Host:      hostname,
				Filename:  filename,
				ParentDir: parent,
				Path:      mypath,
				Size:      info.Size(),
				Mime:      mtype.String(),
				LastScan:  time.Now(),
			}
			db.Save(&file_metadata)
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("Done!")
}
