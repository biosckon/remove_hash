package main

import (
	// "io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// for each file find if it has # in the name
	// check if file is a folder, if it is a folder then recurse into it

	err := filepath.Walk(".", visit)
	checkErr(err)

	log.Printf("done\n")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		log.Println(err)
		return nil
	}

	idx := strings.IndexByte(path, '#') // use byte for speed then use norm to replace
	if idx >= 0 {
		newpath := removeHash(path, idx)

		er := os.Rename(path, newpath)
		if er != nil {
			log.Println(er)
			return nil
		}
		log.Printf("ranming: \n%s\n%s\n", path, newpath)
	}
	return nil
}

func removeHash(s string, i int) (ret string) {
	// replace the # with _
	t := []byte(s)
	t[i] = '_'
	ret = string(t)

	// check if any more left
	idx := strings.IndexByte(ret, '#')
	if idx >= 0 {
		ret = removeHash(ret, idx)
	}

	return ret
}
