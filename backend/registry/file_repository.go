package registry

import (
	"encoding/json"
	"log"
	"os"
)

type FileRepository struct {
	fileName string
	records  []Record
}

func NewFileRepository(fileName string) *FileRepository {

	return &FileRepository{
		fileName: fileName,
		records:  []Record{},
	}
}

func (f *FileRepository) Add(record Record) error {
	f.records = append(f.records, record)
	file, err := os.Create(f.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if data, err := json.Marshal(f.records); err != nil {
		return err
	} else if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

func (f *FileRepository) GetAll() []Record {
	records := []Record{}

	if data, err := os.ReadFile(f.fileName); err != nil {
		log.Print("can't read file", f.fileName)
	} else if err = json.Unmarshal(data, &records); err != nil {
		log.Print("can't unmarshall file data")
	}

	f.records = records

	return records
}

var _ Repository = (*FileRepository)(nil)
