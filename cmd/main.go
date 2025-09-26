package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/italanleal/wfcsd/pkg/io"
	"github.com/italanleal/wfcsd/pkg/sd"
)

func main() {
	df, err := io.ReadCSV("data/anemia_dataset.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = df.BuildBeams()

	if err != nil {
		log.Fatal(err)
	}

	tileList := sd.ExtractPatternsFromPairs(df, sd.WeaklyCorrelatedPairs(df, 0.5, "p"), "p")

	err = sd.ComputeWRAcc(df, "p", tileList)

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(tileList, func(i, j int) bool {
		return tileList[i].WRAcc > tileList[j].WRAcc
	})

	if err != nil {
		log.Fatal(err)
	}

	selectedPatterns := sd.SelectTopPatterns(tileList, 30)

	fmt.Println("Selected patterns")
	for _, pat := range selectedPatterns {
		sd.PrintPattern(df, pat)

		// Se quiser depurar os índices reais, descomente:
		// fmt.Printf("    Positives: %v\n", pat.IndexP)
		// fmt.Printf("    Negatives: %v\n", pat.IndexN)
	}

	collapsedPatterns := sd.PopulationCollapseFunction(df, selectedPatterns, tileList)
	collapsedPatterns = sd.PopulationCollapseFunction(df, collapsedPatterns, tileList)
	collapsedPatterns = sd.PopulationCollapseFunction(df, collapsedPatterns, tileList)

	err = sd.ComputeWRAcc(df, "p", collapsedPatterns)

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(collapsedPatterns, func(i, j int) bool {
		return collapsedPatterns[i].WRAcc > collapsedPatterns[j].WRAcc
	})

	fmt.Println("Collapsed patterns")
	for _, pat := range collapsedPatterns {
		sd.PrintPattern(df, pat)

		// Se quiser depurar os índices reais, descomente:
		// fmt.Printf("    Positives: %v\n", pat.IndexP)
		// fmt.Printf("    Negatives: %v\n", pat.IndexN)
	}

}
