package registry

type Repository interface {
	Add(record Record) error
	GetAll() []Record
}
