package sd

import "github.com/italanleal/wfcsd/pkg/io"

func wracc(p Pattern, target []float64) float64 {
	N := float64(len(target))

	var totalPos float64
	for _, v := range target {
		if v == 1 {
			totalPos++
		}
	}

	nSub := float64(len(p.IndexP) + len(p.IndexN))
	if nSub == 0 {
		return 0.0
	}

	pSub := float64(len(p.IndexP))
	pTotal := totalPos

	return (nSub / N) * ((pSub / nSub) - (pTotal / N))
}

func WRAcc(p Pattern, target []float64) float64 {
	return wracc(p, target)
}

func ComputeWRAcc(df *io.DataFrame, targetColumn string, patterns []Pattern) error {
	target, err := df.ColumnByName(targetColumn)
	if err != nil {
		return err
	}

	for i := range patterns {
		patterns[i].WRAcc = wracc(patterns[i], target)
	}

	return nil
}
