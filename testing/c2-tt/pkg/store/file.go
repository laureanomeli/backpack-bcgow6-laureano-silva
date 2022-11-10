package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	AddMock(mock *Mock)
	ClearMock()
}

type Mock struct {
	Data        []byte
	Error       error
	ReadInvoked bool
}

type FileStore struct {
	FilePath string
	Mock     *Mock
}

func NewStore(FilePath string) Store {
	return &FileStore{FilePath: FilePath}
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Write(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Error != nil {
			return nil
		}

		return nil
	}
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Error != nil {
			return fs.Mock.Error
		}

		fs.Mock.ReadInvoked = true
		return json.Unmarshal(fs.Mock.Data, &data)
	}
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
