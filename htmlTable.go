package dom

type TableS struct {
	ElementI

	Rows           []ElementI
	DefaultStyling map[string]string
}

type TableRowT struct {
	ElementI

	Table   *TableS
	Columns []ElementI
}

func NewTable() *TableS {
	ret := &TableS{
		ElementI:       Doc.CreateElement("table"),
		DefaultStyling: make(map[string]string),
	}
	return ret
}

func (s *TableS) AddDefaultStyling(name, value string) *TableS {
	s.DefaultStyling[name] = value
	return s
}

func (s *TableS) AddRow() *TableRowT {
	row := s.NewChild("tr")
	ret := &TableRowT{
		ElementI: row,
		Table:    s,
	}
	return ret
}

func (s *TableRowT) AddDataCell() ElementI {
	cell := s.NewChild("td")
	s.Columns = append(s.Columns, cell)

	for name, value := range s.Table.DefaultStyling {
		// fmt.Println(name, value)
		cell.Style().Set(name, value)
	}

	return cell
}
