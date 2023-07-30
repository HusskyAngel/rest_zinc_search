package pathsgetter

import (
	"log"
	"os"
	"path/filepath"
)


//iterate over a directory and do an action with every contained file
func PathsGetter(path string, do func(string)) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
    
		if info.Mode().IsRegular() {
			do(path)
		} 
		return nil
	})

  if err != nil{
    log.Fatal("error getting the files")
  }

}
