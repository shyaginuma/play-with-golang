package main

import (
	"fmt"
	"math"
)

type Invoice struct {
	Customer     string
	Performances []*Performance
}

type Play struct {
	Name     string
	PlayType string
}

type Performance struct {
	PlayID   string
	Audience int
	Play     Play
}

type StatementData struct {
	Customer     string
	Performances []*Performance
}

func statement(invoice Invoice) (string, error) {
	data := StatementData{
		Customer:     invoice.Customer,
		Performances: invoice.Performances,
	}
	for _, perf := range data.Performances {
		perf = enrichPerformance(perf)
	}
	return renderPlainText(data)
}

func enrichPerformance(perf *Performance) *Performance {
	perf.Play = playFor(*perf)
	return perf
}

func renderPlainText(data StatementData) (string, error) {
	result := fmt.Sprintf("Statement for %s\n", data.Customer)
	for _, perf := range data.Performances {
		thisAmount, err := amountFor(*perf)
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf("\t%s: %v (%v seat)\n", perf.Play.Name, thisAmount, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is %v\n", totalAmount(data.Performances))
	result += fmt.Sprintf("You earned %v credits\n", totalVolumeCredits(data.Performances))
	return result, nil
}

func totalVolumeCredits(p []*Performance) int {
	result := 0
	for _, perf := range p {
		result += volumeCreditsFor(*perf)
	}
	return result
}

func totalAmount(p []*Performance) int {
	result := 0
	for _, perf := range p {
		thisAmount, err := amountFor(*perf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result += thisAmount
	}
	return result
}

func amountFor(perf Performance) (int, error) {
	result := 0

	switch perf.Play.PlayType {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += 1000 * (perf.Audience - 30)
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += 10000 + 500*(perf.Audience-20)
		}
		result += 300 * perf.Audience
	default:
		return 0, fmt.Errorf("unknown type: %s", perf.Play.PlayType)
	}

	return result, nil
}

func volumeCreditsFor(perf Performance) int {
	result := 0
	result += int(math.Max(float64(perf.Audience-30), 0.0))
	if perf.Play.PlayType == "comedy" {
		result += int(math.Floor(float64(perf.Audience) / 5))
	}
	return result
}

func playFor(perf Performance) Play {
	plays := map[string]Play{
		"hamlet":  {"Hamlet", "tragedy"},
		"as-like": {"As You Like It", "comedy"},
		"othello": {"Othello", "tragedy"},
	}
	return plays[perf.PlayID]
}

func main() {
	invoice := Invoice{
		"Yagi",
		[]*Performance{
			{PlayID: "hamlet", Audience: 55},
		},
	}

	result, err := statement(invoice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
