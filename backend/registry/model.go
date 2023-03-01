package registry

import "sort"

type Model struct {
	repository Repository
}

func NewModel(repository Repository) *Model {
	return &Model{
		repository: repository,
	}
}

func (m *Model) SaveRecord(record Record) error {
	return m.repository.Add(record)
}

func (m *Model) GetRecords() []Record {
	records := m.repository.GetAll()

	sort.Sort(SortByRegisteredAt(records))

	return records
}
