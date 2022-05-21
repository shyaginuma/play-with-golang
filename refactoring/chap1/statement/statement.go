package statement

import (
	"github.com/play-with-golang/refactoring/chap1/calculator"
	"github.com/play-with-golang/refactoring/chap1/model"
)

type StatementData struct {
	Customer           string
	Performances       []model.EnrichedPerformance
	TotalAmount        int
	TotalVolumeCredits int
}

func NewStatement(invoice model.Invoice) StatementData {
	var performances []model.EnrichedPerformance
	for _, p := range invoice.Performances {
		c := calculator.NewPerformanceCalculator(p)
		amount := c.Amount()
		volumeCredits := c.VolumeCredits()
		performances = append(
			performances,
			model.EnrichedPerformance{Performance: p, Amount: amount, VolumeCredits: volumeCredits},
		)
	}

	data := StatementData{
		Customer:           invoice.Customer,
		Performances:       performances,
		TotalAmount:        totalAmount(performances),
		TotalVolumeCredits: totalVolumeCredits(performances),
	}
	return data
}

func totalVolumeCredits(p []model.EnrichedPerformance) int {
	result := 0
	for _, perf := range p {
		result += perf.VolumeCredits
	}
	return result
}

func totalAmount(p []model.EnrichedPerformance) int {
	result := 0
	for _, perf := range p {
		result += perf.Amount
	}
	return result
}
