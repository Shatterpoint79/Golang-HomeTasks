package storage

import (
	"api/binjson/bins"
	"api/binjson/files"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Storage struct {
	FilePath string
}

func NewStorage(filePath string) *Storage {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755) //В документации написано, что права для директорий в Wiтdows игнорируются
	}

	return &Storage{
		FilePath: filePath,
	}
}

func (s *Storage) SaveBinList(binList *bins.BinList) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return err
	}

	return files.WriteFile(data, s.FilePath)
}

func (s *Storage) LoadBinList() (*bins.BinList, error) {
	if !files.IsJSON(s.FilePath) {
		return nil, errors.New("это не JSON-файл")
	}

	data, err := files.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return bins.NewBinList(), nil
		}
		return nil, err
	}

	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, err
	}

	return &binList, nil
}

func (s *Storage) AddBin(bin *bins.Bin) error {
	binList, err := s.LoadBinList()
	if err != nil {
		return err
	}

	binList.AddBin(bin)
	return s.SaveBinList(binList)
}
