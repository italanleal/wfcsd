package sd

import (
	"log"
	"sort"

	"github.com/italanleal/wfcsd/pkg/io"
)

// SelectTopPatterns retorna os primeiros n padrões (ou todos se n for maior que o tamanho).
func SelectTopPatterns(patterns []Pattern, n int) []Pattern {
	if n > len(patterns) {
		n = len(patterns)
	}
	return patterns[:n]
}

// PopulationCollapseFunction realiza o merge apenas sobre os padrões já selecionados.
func PopulationCollapseFunction(df *io.DataFrame, selectedPatterns, allPatterns []Pattern) []Pattern {
	target, err := df.ColumnByName("p")
	if err != nil {
		log.Fatal(err)
	}

	var mergedPatterns []Pattern
	for _, p := range selectedPatterns {
		// Calcula os candidatos de merge
		p.Options = ComputeMergeCandidates(allPatterns, p)

		if len(p.Options) == 0 {
			mergedPatterns = append(mergedPatterns, p)
			continue
		}

		// Ordena options pelo score
		sort.Slice(p.Options, func(i, j int) bool {
			scoreI := selectionScore(p, allPatterns[p.Options[i]])
			scoreJ := selectionScore(p, allPatterns[p.Options[j]])
			return scoreI > scoreJ
		})

		merged := false
		usedOptions := make(map[int]bool)

		for len(p.Options) > 0 {
			bestOptionIdx := p.Options[0]
			if usedOptions[bestOptionIdx] {
				p.Options = p.Options[1:]
				continue
			}
			usedOptions[bestOptionIdx] = true

			candidate := allPatterns[bestOptionIdx]
			newPattern, ok := MergePatterns(p, candidate)

			if !ok {
				p.Options = p.Options[1:]
				continue
			}
			newPattern.WRAcc = wracc(newPattern, target)

			if newPattern.WRAcc >= p.WRAcc {
				mergedPatterns = append(mergedPatterns, newPattern)
				merged = true

				break
			}
		}

		if !merged {
			mergedPatterns = append(mergedPatterns, p)
		}
	}

	return mergedPatterns
}

func selectionScore(p, candidate Pattern) float64 {
	// Positivos em comum
	commonPos := countCommonIndices(p.IndexP, candidate.IndexP)

	// Negativos distintos do candidate que não estão em p
	//distinctNeg := countDistinctIndices(p.IndexN, candidate.IndexN)

	// Score final: pondera positivos e negativos
	return float64(commonPos) //+ 0.2*float64(distinctNeg) // peso 0.5 para negativos
}

// Conta índices comuns entre dois slices
func countCommonIndices(a, b []int) int {
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	count := 0
	for _, v := range b {
		if _, exists := m[v]; exists {
			count++
		}
	}
	return count
}

// Conta índices do slice b que não estão em a (distintos)
func countDistinctIndices(a, b []int) int {
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	count := 0
	for _, v := range b {
		if _, exists := m[v]; !exists {
			count++
		}
	}
	return count
}
