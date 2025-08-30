package ucexcel

import (
	"fmt"
	"xmlparser/domain"
	"xmlparser/ucexcel/address"
)

func (ue *ucexcel) Report(report []*domain.Record) error {
	ue.sheet = "Sheet1"
	// countRow = 1
	ue.address = address.New(1, 0)
	if err := ue.header(); err != nil {
		return fmt.Errorf("excel report header %w", err)
	}
	for _, ss := range report {
		if err := ue.line(ss); err != nil {
			return fmt.Errorf("excel report line %w", err)
		}
	}

	return nil
}

func (ue *ucexcel) header() error {
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Штрих-код товара EAN13"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Код товара ЕГАИС"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Наименование товара"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Data Matrix"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "PDF 417"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Код короба"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), "Код паллеты"); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextRow()
	return nil
}

func (ue *ucexcel) line(s *domain.Record) error {
	ue.address.NextCol()
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), s.CodeAp); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), s.NameAp); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), s.Text); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	ue.address.NextCol()
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), s.Korob); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextCol()
	if err := ue.file.SetCellStr(ue.sheet, ue.address.Address(), s.Palet); err != nil {
		return fmt.Errorf("excel error %w", err)
	}
	ue.address.NextRow()
	return nil
}
