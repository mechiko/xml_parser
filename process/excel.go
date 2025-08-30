package process

import (
	"fmt"
	"xmlparser/ucexcel"
)

func (p *process) Excel(name string, out string) error {
	p.NameFileWithoutExt = name
	excel := ucexcel.New(name)
	if err := excel.Open(); err != nil {
		return fmt.Errorf("excel open: %w", err)
	}
	if err := excel.Report(p.Records); err != nil {
		return fmt.Errorf("excel report: %w", err)
	}
	if err := excel.Save(out); err != nil {
		return fmt.Errorf("excel save: %w", err)
	}
	return nil
}
