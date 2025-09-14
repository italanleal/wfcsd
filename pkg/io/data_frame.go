package io

import (
	"fmt"
	"math"
)

type DataFrame struct {
	Header        []string    // column names
	Rows          [][]float64 // numeric data
	DiscretScales map[string]*DiscretScale
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

func (df *DataFrame) BuildBeams() error {
	if len(df.Rows) == 0 {
		return fmt.Errorf("%s | empty dataframe", moduleName)
	}

	df.DiscretScales = make(map[string]*DiscretScale)

	k := int(math.Ceil(1 + 3.322*math.Log10(float64(len(df.Rows)))))
	if k < 1 {
		k = 1
	}

	for j, name := range df.Header {
		min, max := math.Inf(1), math.Inf(-1)
		for i := 0; i < len(df.Rows); i++ {
			v := df.Rows[i][j]
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		df.DiscretScales[name] = &DiscretScale{
			Min:   min,
			Max:   max,
			Beams: k,
		}
	}
	return nil
}
