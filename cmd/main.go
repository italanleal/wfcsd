package main

import (
	"fmt"
	"log"

	"github.com/italanleal/wfcsd/pkg/io"
	"github.com/italanleal/wfcsd/pkg/math"
	"github.com/italanleal/wfcsd/pkg/sd"
)

func main() {
	// Load the CSV into a DataFrame
	df, err := io.ReadCSV("data/anemia_dataset.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Columns:", df.Header)

	numCols := len(df.Header)
	columns := make([][]float64, numCols)

	// Extract each column
	for j := 0; j < numCols; j++ {
		col, err := df.ColumnByName(df.Header[j])
		if err != nil {
			log.Fatal(err)
		}
		columns[j] = col
	}

	// Compute correlation matrix with column names
	fmt.Printf("%12s", "") // top-left empty
	for _, name := range df.Header {
		fmt.Printf("%12s", name)
	}
	fmt.Println()

	for i, rowName := range df.Header {
		fmt.Printf("%12s", rowName)
		for j := 0; j < numCols; j++ {
			corrVal := math.Correlation(columns[i], columns[j])
			fmt.Printf("%12.6f ", corrVal)
		}
		fmt.Println()
	}

	pairs := sd.WeaklyCorrelatedPairs(df, 0.5, "p")

	// Iterate over all pairs
	for idx, p := range pairs {
		col1 := p[0]
		col2 := p[1]
		fmt.Printf("Pair %d: %s <-> %s\n", idx+1, col1, col2)
	}

	// Initialize tileList to store all patterns
	tileList := sd.ExtractPatternsFromPairs(df, pairs, "p")

	for i, pat := range tileList {
		fmt.Printf(
			"Pattern %d: %s=%.2f <-> %s=%.2f | Freq: %d/%d\n",
			i+1,
			pat.Items[0].Attr, pat.Items[0].Value,
			pat.Items[1].Attr, pat.Items[1].Value,
			len(pat.Index),
			pat.Freq,
		)
	}

}
