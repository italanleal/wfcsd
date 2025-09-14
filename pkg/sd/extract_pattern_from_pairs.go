package sd

import (
	"log"

	"github.com/italanleal/wfcsd/pkg/io"
)

// func extractPatternsFromPairs(df *io.DataFrame, pairs [][2]string, targetColumn string) []Pattern {
// 	colIdx := make(map[string]int)
// 	for i, name := range df.Header {
// 		colIdx[name] = i
// 	}

// 	targetIdx, ok := colIdx[targetColumn]
// 	if !ok {
// 		log.Fatalf("Target column not found: %s", targetColumn)
// 	}

// 	var patterns []Pattern

// 	for _, pair := range pairs {
// 		attr1, attr2 := pair[0], pair[1]
// 		idx1, ok1 := colIdx[attr1]
// 		idx2, ok2 := colIdx[attr2]
// 		if !ok1 || !ok2 {
// 			continue
// 		}

// 		patMap := make(map[[2]float64]*Pattern) // chave = valores (v1, v2)

// 		for rowIdx, row := range df.Rows {
// 			v1 := row[idx1]
// 			v2 := row[idx2]
// 			key := [2]float64{v1, v2}

// 			p, exists := patMap[key]
// 			if !exists {
// 				p = &Pattern{
// 					Items: []Item{
// 						{Attr: attr1, Value: v1},
// 						{Attr: attr2, Value: v2},
// 					},
// 					Index: []int{},
// 					Freq:  0,
// 				}
// 				patMap[key] = p
// 			}

// 			// Sempre incrementa a frequência
// 			p.Freq++

// 			// Se target == 1, adiciona índice positivo
// 			if row[targetIdx] == 1 {
// 				p.Index = append(p.Index, rowIdx)
// 			}
// 		}

// 		// Adiciona todos os patterns do par de colunas à lista final
// 		for _, p := range patMap {
// 			patterns = append(patterns, *p)
// 		}
// 	}

// 	return patterns
// }

func extractPatternsFromPairs(df *io.DataFrame, pairs [][2]string, targetColumn string) []Pattern {
	// Mapear nomes de colunas para índice
	colIdx := make(map[string]int)
	for i, name := range df.Header {
		colIdx[name] = i
	}

	targetIdx, ok := colIdx[targetColumn]
	if !ok {
		log.Fatalf("%s | Target column not found: %s", moduleName, targetColumn)
	}

	var patterns []Pattern

	for _, pair := range pairs {
		attr1, attr2 := pair[0], pair[1]
		idx1, ok1 := colIdx[attr1]
		idx2, ok2 := colIdx[attr2]
		if !ok1 || !ok2 {
			continue
		}

		// Pega os BeamScalers para cada coluna
		scaler1 := df.DiscretScales[attr1]
		scaler2 := df.DiscretScales[attr2]

		// chave = (beam1, beam2)
		patMap := make(map[[2]int]*Pattern)

		for rowIdx, row := range df.Rows {
			b1 := scaler1.Beam(row[idx1])
			b2 := scaler2.Beam(row[idx2])
			key := [2]int{b1, b2}

			p, exists := patMap[key]
			if !exists {
				p = &Pattern{
					Items: []Item{
						{Attr: attr1, Value: b1}, // Armazeno o índice como float64 se Item.Value for float64
						{Attr: attr2, Value: b2},
					},
					IndexP: []int{},
					IndexN: []int{},
					Freq:   0,
				}
				patMap[key] = p
			}

			// Incrementa frequência
			p.Freq++

			// Se target == 1, adiciona índice positivo
			if row[targetIdx] == 1 {
				p.IndexP = append(p.IndexP, rowIdx)
			} else {
				p.IndexN = append(p.IndexN, rowIdx)
			}
		}

		// Adiciona todos os patterns do par de colunas à lista final
		for _, p := range patMap {
			patterns = append(patterns, *p)
		}
	}

	return patterns
}

// ExtractPatternsFromPairs is the exported wrapper
func ExtractPatternsFromPairs(df *io.DataFrame, pairs [][2]string, targetColumn string) []Pattern {
	return extractPatternsFromPairs(df, pairs, targetColumn)
}

// Helper: find index of a string in slice
func indexOf(slice []string, val string) int {
	for i, s := range slice {
		if s == val {
			return i
		}
	}
	return -1
}
