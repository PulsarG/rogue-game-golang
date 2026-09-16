package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteSave(filepath string, data any) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("Failed open/create file: %w", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("Failed encode: %w", err)
	}
	return nil
}

func ReadSave[T any](filepath string) (T, error) {
	var res T
	file, err := os.Open(filepath)
	if err != nil {
		return res, fmt.Errorf("Failed open file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&res); err != nil {
		return res, fmt.Errorf("Failed decode file: %w", err)
	}
	return res, nil
}
