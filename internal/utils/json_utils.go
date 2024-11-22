package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJSONFromFile(filename string, v interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Если файл не существует, возвращаем nil, чтобы можно было его создать
		}
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() { _ = file.Close() }()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		return nil // Если файл пустой, возвращаем nil
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func WriteJSONToFile(filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}

	return nil
}
