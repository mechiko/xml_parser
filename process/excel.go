package process

import (
	"log"
	"os"
	"xmlparser/ucexcel"
)

func checkErr(e error) {
	if e != nil {
		log.Printf("%v", e)
		os.Exit(-1)
	}
}

func (p *process) Excel(name string, out string) error {
	p.NameFileWithoutExt = name
	excel := ucexcel.New(name)
	err := excel.Open()
	checkErr(err)
	err = excel.Report(p.Records)
	checkErr(err)
	err = excel.Save(out)
	checkErr(err)

	return nil
}
