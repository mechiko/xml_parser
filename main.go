package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"xmlparser/process"

	"github.com/mechiko/utility"
)

const inDir = `IN`
const outDir = `OUT`

func checkErr(e error) {
	if e != nil {
		log.Printf("%v", e)
		os.Exit(-1)
	}
}

func main() {
	re, err := regexp.Compile(`.*\.xml$`)
	checkErr(err)
	files, err := utility.FilteredSearchOfDirectoryTree(re, inDir)
	checkErr(err)
	checkErr(os.MkdirAll(outDir, 0o755))
	for _, file := range files {
		p, err := process.New(file)
		// заполняем короба палеты и КМ
		checkErr(err)
		err = p.Load()
		checkErr(err)
		name := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		err = p.Scan()
		checkErr(err)
		err = p.Excel(name, "OUT")
		checkErr(err)

	}
}
