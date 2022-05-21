package calculator

import (
	"math"

	"github.com/play-with-golang/refactoring/performance"
)

type Calculator interface {
	Amount() int
	VolumeCredits() int
}

type PerformanceCalculator struct {
	Performance       *performance.Performance
	Play              performance.Play
	AmountFunc        func(performance.Performance) int
	VolumeCreditsFunc func(performance.Performance) int
}

func (pc *PerformanceCalculator) Amount() int {
	return pc.AmountFunc(pc.Performance)
}

func (pc *PerformanceCalculator) VolumeCredits() int {
	result := int(math.Max(float64(pc.Performance.Audience-30), 0.0))
	result += pc.VolumeCreditsFunc(pc.Performance)
	return result
}

func tragedyAmount(p performance.Performance) int {
	result := 40000
	if p.Audience > 30 {
		result += 1000 * (p.Audience - 30)
	}
	return result
}

func comedyAmount(p performance.Performance) int {
	result := 30000
	if p.Audience > 20 {
		result += 10000 + 500*(p.Audience-20)
	}
	result += 300 * p.Audience
	return result
}

func comeedyVolumeCredits(p performance.Performance) int {
	return int(math.Floor(float64(p.Audience) / 5))
}
