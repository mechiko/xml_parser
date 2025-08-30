package ucexcel

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	_ "time/tzdata"
)

func (ue *ucexcel) Save(out string) (retErr error) {
	if ue.file == nil {
		return fmt.Errorf("excel workbook is not initialized; call Open() first")
	}
	if err := os.MkdirAll(out, 0o755); err != nil {
		return fmt.Errorf("create out dir %q: %w", out, err)
	}
	fname := ue.name + ".xlsx"
	fileAbs := filepath.Join(out, fname)
	file, err := os.OpenFile(fileAbs, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o660)
	if err != nil {
		return fmt.Errorf("open %q: %w", fileAbs, err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			retErr = errors.Join(retErr, fmt.Errorf("close %q: %w", fileAbs, cerr))
		}
	}()

	if _, err := ue.file.WriteTo(file); err != nil {
		return fmt.Errorf("write workbook to %q: %w", fileAbs, err)
	}

	// Free resources held by excelize (temp files, buffers, etc.).
	if cerr := ue.file.Close(); cerr != nil {
		retErr = errors.Join(retErr, fmt.Errorf("close workbook: %w", cerr))
	}
	return retErr
}
