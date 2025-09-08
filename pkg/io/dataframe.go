package io

import "fmt"

type DataFrame struct {
	Header []string    // column names
	Rows   [][]float64 // numeric data
}

func (df *DataFrame) ColumnByName(name string) ([]float64, error) {
	for j, colName := range df.Header {
		if colName == name {
			col := make([]float64, len(df.Rows))
			for i := 0; i < len(df.Rows); i++ {
				col[i] = df.Rows[i][j]
			}
			return col, nil
		}
	}
	return nil, fmt.Errorf("%s | column %s not found", moduleName, name)
}
