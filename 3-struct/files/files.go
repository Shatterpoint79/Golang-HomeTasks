package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения %s: %w", filepath.Base(filename), err)
	}
	return data, nil
}

func IsJSON(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}

func WriteFile(data []byte, filename string) error {
	if len(data) == 0 {
		return fmt.Errorf("нет данных в файле %s", filename)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil { //владелец - чтение и запись, остальные - только чтение
		return fmt.Errorf("ошибка записи в %s: %w", filepath.Base(filename), err)
	}
	return nil
}
