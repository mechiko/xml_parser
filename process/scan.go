package process

import (
	"fmt"
	"strings"
	"xmlparser/domain"
)

func (p *process) Scan() error {
	if p.Documents == nil {
		return fmt.Errorf("хмл документ nil")
	}
	nameAp := p.Documents.Document.WayBillV3Transit.Content.Position.FullName
	codeAp := p.Documents.Document.WayBillV3Transit.Content.Position.Alccode
	p.Records = make([]*domain.Record, 0)
	for _, position := range p.Documents.Document.WayBillV3Transit.Content.Position.MarkInfo.Boxpos {
		boxNumber := strings.TrimSpace(position.Boxnumber)
		if _, ok := p.Koroba[boxNumber]; !ok {
			p.Koroba[boxNumber] = make([]*domain.Record, 0)
		}
		for _, amBox := range position.Amclist.Amc {
			rec := domain.Record{
				CodeAp: codeAp,
				NameAp: nameAp,
				Text:   strings.TrimSpace(amBox.Text),
				Korob:  boxNumber,
			}
			p.Records = append(p.Records, &rec)
			p.Koroba[boxNumber] = append(p.Koroba[boxNumber], &rec)
		}
	}
	for _, palet := range p.Documents.Document.WayBillV3Transit.Content.Position.BoxInfo.Boxtree.Bl {
		paletNumber := palet.Boxnum
		for _, box := range palet.Bl.Boxnum {
			if korobRecs, ok := p.Koroba[box]; ok {
				for i := range korobRecs {
					korobRecs[i].Palet = paletNumber
				}
			} else {
				return fmt.Errorf("коробка %s в палете %s не найдена", box, paletNumber)
			}
		}
	}
	return nil
}
