package template

import "strings"

type ISort interface {
	Init() string
	Shuffle() string
	Merge() string
}

type SortA struct{}

func (s *SortA) Init() string {
	return "SortA init"
}
func (s *SortA) Shuffle() string {
	return "SortA shuffle"
}
func (s *SortA) Merge() string {
	return "SortA merge"
}

type SortB struct{}

func (s *SortB) Init() string {
	return "SortB init"
}
func (s *SortB) Shuffle() string {
	return "SortB shuffle"
}
func (s *SortB) Merge() string {
	return "SortB merge"
}

type ISortTemplate struct {
	isort ISort
}

func (s *ISortTemplate) SetSort(sort ISort) {
	s.isort = sort
}

func (s *ISortTemplate) Sort() string {
	var output []string
	output = append(output, s.isort.Init())
	output = append(output, s.isort.Shuffle())
	output = append(output, s.isort.Merge())
	return strings.Join(output, ",")
}
