package mocks

import (
	"fmt"

	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/internal/transactions"
)

type MockStorage struct {
	DataMock []transactions.Transaction
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]transactions.Transaction)
	*a = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]transactions.Transaction)
	m.DataMock = append(m.DataMock, a[len(a)-1])
	return nil
}
