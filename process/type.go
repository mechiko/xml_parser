package process

import (
	"fmt"
	"path/filepath"
	"strings"
	"xmlparser/domain"

	"github.com/mechiko/utility"
)

type process struct {
	NameFileWithoutExt string
	File               string
	Records            []*domain.Record
	Koroba             map[string][]*domain.Record
	Documents          *Documents
}

func New(file string) (*process, error) {
	if file == "" {
		return nil, fmt.Errorf("file name empty")
	}
	if !utility.PathOrFileExists(file) {
		return nil, fmt.Errorf("file not found")
	}
	name := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	p := &process{
		File:               file,
		NameFileWithoutExt: name,
	}
	return p, nil
}
