package statement

import (
	"math"

	"github.com/play-with-golang/refactoring/chap1/master"
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
		performances = append(performances, EnrichPerformance(p))
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

func EnrichPerformance(p model.Performance) model.EnrichedPerformance {
	play := master.Plays[p.PlayID]

	amount := 0
	volumeCredits := int(math.Max(float64(p.Audience-30), 0.0))
	switch play.PlayType {
	case "tragedy":
		amount = 40000
		if p.Audience > 30 {
			amount += 1000 * (p.Audience - 30)
		}
	case "comedy":
		amount = 30000
		if p.Audience > 20 {
			amount += 10000 + 500*(p.Audience-20)
		}
		amount += 300 * p.Audience

		volumeCredits += int(math.Floor(float64(p.Audience) / 5))
	}
	return model.EnrichedPerformance{
		Performance:   p,
		Play:          play,
		Amount:        amount,
		VolumeCredits: volumeCredits,
	}
}
