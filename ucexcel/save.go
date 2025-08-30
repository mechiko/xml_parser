package ucexcel

import (
	"fmt"
	"os"
	"path/filepath"
	_ "time/tzdata"
)

func (ue *ucexcel) Save(out string) error {
	if err := os.MkdirAll(out, 0o755); err != nil {
		return fmt.Errorf("create out dir %q: %w", out, err)
	}
	fname := ue.name + ".xlsx"
	fileAbs := filepath.Join(out, fname)
	file, err := os.OpenFile(fileAbs, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() error {
		if err := file.Close(); err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}()

	if _, err := ue.file.WriteTo(file); err != nil {
		return fmt.Errorf("%w", err)
	}

	return err
}
