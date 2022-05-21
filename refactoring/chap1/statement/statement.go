package statement

import (
	"github.com/play-with-golang/refactoring/chap1/calculator"
	"github.com/play-with-golang/refactoring/chap1/master"
	"github.com/play-with-golang/refactoring/chap1/model"
)

type StatementData struct {
	Customer           string
	Performances       []*model.Performance
	TotalAmount        int
	TotalVolumeCredits int
}

func NewStatement(invoice model.Invoice) StatementData {
	data := StatementData{
		Customer:     invoice.Customer,
		Performances: invoice.Performances,
	}
	for _, p := range data.Performances {
		p.Play = master.Plays[p.PlayID]
		c := calculator.NewPerformanceCalculator(p)
		p.Amount = c.Amount()
		p.VolumeCredits = c.VolumeCredits()
	}
	data.TotalAmount = totalAmount(data.Performances)
	data.TotalVolumeCredits = totalVolumeCredits(data.Performances)
	return data
}

func totalVolumeCredits(p []*model.Performance) int {
	result := 0
	for _, perf := range p {
		result += perf.VolumeCredits
	}
	return result
}

func totalAmount(p []*model.Performance) int {
	result := 0
	for _, perf := range p {
		result += perf.Amount
	}
	return result
}
