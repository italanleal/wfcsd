package sd

import (
	"fmt"

	"github.com/italanleal/wfcsd/pkg/io"
)

// Pattern represents a combination of exactly two items with associated positive row indices
type Pattern struct {
	Items   []Item  // subgroup items
	IndexP  []int   // indices of rows where this pattern occurs (positive cases)
	IndexN  []int   // indices of rows where this pattern occurs (negative cases)
	Options []int   // indices of possible merge options
	Freq    int     // absolute frequency = len(Index)
	WRAcc   float64 // quality
}

func MergePatterns(p1, p2 Pattern) (Pattern, bool) {
	// Regra 1: pelo menos um tem exatamente 2 items
	if !((len(p1.Items) == 2 && len(p2.Items) >= 2) || (len(p2.Items) == 2 && len(p1.Items) >= 2)) {
		return Pattern{}, false
	}

	// Encontrar itens comuns
	commonCount := 0
	var commonItem Item
	for _, i1 := range p1.Items {
		for _, i2 := range p2.Items {
			if i1 == i2 { // compara Attr e Value
				commonCount++
				commonItem = i1
			}
		}
	}

	// Regra 2: deve ter exatamente 1 item comum
	if commonCount != 1 {
		return Pattern{}, false
	}

	// Construir itens combinados (todos, sem duplicatas)
	combinedItems := []Item{commonItem}
	for _, i := range p1.Items {
		if i != commonItem {
			combinedItems = append(combinedItems, i)
		}
	}
	for _, i := range p2.Items {
		if i != commonItem {
			combinedItems = append(combinedItems, i)
		}
	}

	// Interseção de positivos (IndexP)
	indexPMap := make(map[int]struct{}, len(p1.IndexP))
	for _, idx := range p1.IndexP {
		indexPMap[idx] = struct{}{}
	}
	var combinedIndexP []int
	for _, idx := range p2.IndexP {
		if _, exists := indexPMap[idx]; exists {
			combinedIndexP = append(combinedIndexP, idx)
		}
	}

	// Interseção de negativos (IndexN)
	indexNMap := make(map[int]struct{}, len(p1.IndexN))
	for _, idx := range p1.IndexN {
		indexNMap[idx] = struct{}{}
	}
	var combinedIndexN []int
	for _, idx := range p2.IndexN {
		if _, exists := indexNMap[idx]; exists {
			combinedIndexN = append(combinedIndexN, idx)
		}
	}

	// Novo Pattern
	newPattern := Pattern{
		Items:  combinedItems,
		IndexP: combinedIndexP,
		IndexN: combinedIndexN,
		Freq:   len(combinedIndexP) + len(combinedIndexN),
	}

	return newPattern, true
}

// func ComputeMergeCandidates(patterns []Pattern, target Pattern) []int {
// 	var candidates []int

// 	for i, p := range patterns {

// 		// Regra: pelo menos um deve ter exatamente 2 items
// 		if len(p.Items) != 2 && len(target.Items) < 2 {
// 			continue
// 		}

// 		// Contar itens comuns
// 		commonCount := 0
// 		for _, i1 := range target.Items {
// 			for _, i2 := range p.Items {
// 				if i1.Attr == i2.Attr && i1.Value == i2.Value { // compara Attr e Value
// 					commonCount++
// 				}
// 			}

// 		}

// 		// Regra: deve ter exatamente 1 item comum
// 		if commonCount == 1 {
// 			candidates = append(candidates, i)
// 		}
// 	}

// 	return candidates
// }

// devolve os índices de patterns que contêm um item específico
func candidatesForItem(patterns []Pattern, item Item) []int {
	var idx []int
	for i, p := range patterns {
		// precisa conter exatamente esse item
		if containsItem(p.Items, item) {
			idx = append(idx, i)
		}
	}
	return idx
}

func containsItem(items []Item, item Item) bool {
	for _, it := range items {
		if it.Attr == item.Attr && it.Value == item.Value {
			return true
		}
	}
	return false
}

func hasAttrCollision(candidate, target []Item, anchor Item) bool {
	// cria um conjunto dos atributos do target, exceto o anchor
	attrSet := make(map[string]bool)
	for _, t := range target {
		if t.Attr != anchor.Attr {
			attrSet[t.Attr] = true
		}
	}
	// se o candidato tiver algum Attr já presente no target (exceto o anchor),
	// é colisão.
	for _, c := range candidate {
		if c.Attr != anchor.Attr && attrSet[c.Attr] {
			return true
		}
	}
	return false
}

func ComputeMergeCandidates(patterns []Pattern, target Pattern) []int {
	// conjunto (set) para evitar duplicação
	seen := make(map[int]bool)

	for _, tItem := range target.Items {
		cands := candidatesForItem(patterns, tItem)

		for _, ci := range cands {
			// além de conter o item, checa colisão com os demais itens do target
			if !hasAttrCollision(patterns[ci].Items, target.Items, tItem) {
				seen[ci] = true
			}
		}
	}

	// converte o set em slice
	result := make([]int, 0, len(seen))
	for i := range seen {
		result = append(result, i)
	}
	return result
}

func PrintPattern(df *io.DataFrame, pat Pattern) {
	// Monta as descrições de cada item
	itemStrs := make([]string, len(pat.Items))
	for j, item := range pat.Items {
		bs := df.DiscretScales[item.Attr]
		min, max := bs.BeamBounds(item.Value)
		itemStrs[j] = fmt.Sprintf("%s [%.2f, %.2f]", item.Attr, min, max)
	}

	// Junta com "<->"
	patternDesc := ""
	for j, s := range itemStrs {
		if j > 0 {
			patternDesc += " <-> "
		}
		patternDesc += s
	}

	// Imprime
	fmt.Printf(
		"Pattern %s | WRAcc: %.4f | len(IndexP)=%d, len(IndexN)=%d\n",
		patternDesc,
		pat.WRAcc,
		len(pat.IndexP),
		len(pat.IndexN),
	)
}
