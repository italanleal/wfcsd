package io

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

// readCSV loads CSV into a DataFrame with header + numeric rows
func readCSV(path string) (*DataFrame, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("%s | CSV must have header + at least one row", moduleName)
	}

	header := records[0]
	data := make([][]float64, len(records)-1)

	for i := 1; i < len(records); i++ {
		row := make([]float64, len(records[i]))
		for j, v := range records[i] {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("%s | invalid float at row %d col %d: %v", moduleName, i, j, err)
			}
			f = math.Round(f*1) / 1 //rounding values due to the short amount of samples from dataset
			row[j] = f
		}
		data[i-1] = row
	}

	return &DataFrame{Header: header, Rows: data}, nil
}

// ReadCSV is the exported wrapper
func ReadCSV(path string) (*DataFrame, error) {
	return readCSV(path)
}
