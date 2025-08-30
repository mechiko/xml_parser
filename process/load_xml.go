package process

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func (p *process) Load() error {
	docs, err := readXml(p.File)
	if err != nil {
		return fmt.Errorf("ошибка получения xml %w", err)
	}
	if docs != nil {
		p.Documents = docs
	} else {
		return fmt.Errorf("загружен из хмл документ nil")
	}
	return nil

}

func readXml(filePath string) (*Documents, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла %w", err)
	}
	defer f.Close()
	byteValue, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error read file %w", err)
	}
	var docs Documents
	err = xml.Unmarshal(byteValue, &docs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal xml %w", err)
	}
	return &docs, nil
}
