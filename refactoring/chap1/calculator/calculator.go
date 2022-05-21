package calculator

import (
	"math"

	"github.com/play-with-golang/refactoring/chap1/model"
)

type Calculator interface {
	Amount() int
	VolumeCredits() int
}

type PerformanceCalculator struct {
	Performance       model.Performance
	amountFunc        func(model.Performance) int
	volumeCreditsFunc func(model.Performance) int
}

func (pc *PerformanceCalculator) Amount() int {
	return pc.amountFunc(pc.Performance)
}

func (pc *PerformanceCalculator) VolumeCredits() int {
	result := int(math.Max(float64(pc.Performance.Audience-30), 0.0))
	result += pc.volumeCreditsFunc(pc.Performance)
	return result
}

func NewPerformanceCalculator(p model.EnrichedPerformance) PerformanceCalculator {
	var amountFunc func(model.Performance) int
	var volumeCreditsFunc func(model.Performance) int

	switch p.Play.PlayType {
	case "tragedy":
		amountFunc = tragedyAmount
	case "comedy":
		amountFunc = comedyAmount
		volumeCreditsFunc = comedyVolumeCredits
	}
	return PerformanceCalculator{
		p,
		amountFunc,
		volumeCreditsFunc,
	}
}

func tragedyAmount(p model.Performance) int {
	result := 40000
	if p.Audience > 30 {
		result += 1000 * (p.Audience - 30)
	}
	return result
}

func comedyAmount(p model.Performance) int {
	result := 30000
	if p.Audience > 20 {
		result += 10000 + 500*(p.Audience-20)
	}
	result += 300 * p.Audience
	return result
}

func comedyVolumeCredits(p model.Performance) int {
	return int(math.Floor(float64(p.Audience) / 5))
}
