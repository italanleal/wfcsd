package io

import "math"

type DiscretScale struct {
	Min, Max float64
	Beams    int
}

// Beam returns the beam index (0-based) in which v lies.
func (b *DiscretScale) Beam(v float64) int {
	width := (b.Max - b.Min) / float64(b.Beams)
	return int(math.Floor((v - b.Min) / width))
}

// Beam returns the beam bounds by its index
func (b *DiscretScale) BeamBounds(idx int) (float64, float64) {
	width := (b.Max - b.Min) / float64(b.Beams)
	minBound := b.Min + float64(idx)*width
	maxBound := b.Min + float64(idx+1)*width

	return minBound, maxBound
}
