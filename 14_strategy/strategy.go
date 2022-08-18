package strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

type StorageStrategy interface {
	Save(name string, data []byte) error
}

var Strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := Strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}

type fileStorage struct{}

func (s *fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

type encryptFileStorage struct{}

func (s *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	// 这里实现加密算法
	return data, nil
}
