package sd

import (
	"fmt"

	"github.com/italanleal/wfcsd/pkg/io"
)

func findPatternOccurrences(df *io.DataFrame, p Pattern) ([]int, error) {
	// If the dataframe has no rows, no pattern can match.
	if len(df.Rows) == 0 {
		return []int{}, nil
	}
	// For efficiency, first map the attribute names from the pattern to their column indices.
	// This avoids searching the header slice repeatedly inside the main loop.
	type condition struct {
		colIndex int
		min      float64
		max      float64
	}

	// Create a lookup map for dataframe headers.
	headerMap := make(map[string]int, len(df.Header))
	for i, name := range df.Header {
		headerMap[name] = i
	}

	// Build the list of conditions to check for each row.
	conditions := make([]condition, len(p.Items))
	for i, item := range p.Items {
		colIndex, found := headerMap[item.Attr]
		if !found {
			return nil, fmt.Errorf("%s | attribute '%s' from pattern not found in dataframe header", moduleName, item.Attr)
		}
		conditions[i] = condition{
			colIndex: colIndex,
			min:      item.Min,
			max:      item.Max,
		}
	}

	// Iterate through each row and check if it satisfies all conditions.
	var matchingIndices []int
	for rowIndex, row := range df.Rows {
		isMatch := true
		for _, cond := range conditions {
			cellValue := row[cond.colIndex]
			// If the value is outside the [min, max] range, this row is not a match.
			if cellValue < cond.min || cellValue > cond.max {
				isMatch = false
				break // No need to check other conditions for this row.
			}
		}

		if isMatch {
			// If isMatch is still true, the row satisfies all conditions.
			matchingIndices = append(matchingIndices, rowIndex)
		}
	}

	return matchingIndices, nil
}

func FindPatternOcurrences(df *io.DataFrame, p Pattern) ([]int, error) {
	return findPatternOccurrences(df, p)
}
