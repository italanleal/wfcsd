package sd

import (
	m2 "math"

	"github.com/italanleal/wfcsd/pkg/io"
	m1 "github.com/italanleal/wfcsd/pkg/math"
)

// WeaklyCorrelatedPairs finds column pairs in df with |correlation| below threshold,
// ignoring the target column. Returns unordered column name pairs.
func weaklyCorrelatedPairs(df *io.DataFrame, threshold float64, targetColumn string) [][2]string {
	numCols := len(df.Header)
	columns := make([][]float64, numCols)

	// Extract all columns
	for j := 0; j < numCols; j++ {
		col, _ := df.ColumnByName(df.Header[j])
		columns[j] = col
	}

	var result [][2]string

	for i := 0; i < numCols; i++ {
		colNameI := df.Header[i]
		if colNameI == targetColumn {
			continue
		}
		for j := i + 1; j < numCols; j++ { // unordered pairs
			colNameJ := df.Header[j]
			if colNameJ == targetColumn {
				continue
			}
			corr := m1.Correlation(columns[i], columns[j])
			if m2.Abs(corr) < threshold {
				result = append(result, [2]string{colNameI, colNameJ})
			}
		}
	}

	return result
}

// WeaklyCorrelatedPairs is the exported wrapper
func WeaklyCorrelatedPairs(df *io.DataFrame, threshold float64, targetColumn string) [][2]string {
	return weaklyCorrelatedPairs(df, threshold, targetColumn)
}
