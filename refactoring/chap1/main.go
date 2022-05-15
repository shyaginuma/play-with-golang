package main

import (
	"fmt"
	"math"
)

type Invoice struct {
	Customer     string
	Performances []Performance
}

type Play struct {
	Name     string
	PlayType string
}

type Performance struct {
	PlayID   string
	Audience int
}

func statement(invoice Invoice) (string, error) {
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		thisAmount, err := amountFor(perf)
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf("\t%s: %v (%v seat)\n", playFor(perf).Name, thisAmount, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is %v\n", totalAmount(invoice))
	result += fmt.Sprintf("You earned %v credits\n", totalVolumeCredits(invoice))
	return result, nil
}

func totalVolumeCredits(invoice Invoice) int {
	result := 0
	for _, perf := range invoice.Performances {
		result += volumeCreditsFor(perf)
	}
	return result
}

func totalAmount(invoice Invoice) int {
	result := 0
	for _, perf := range invoice.Performances {
		thisAmount, err := amountFor(perf)
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

	switch playFor(perf).PlayType {
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
		return 0, fmt.Errorf("unknown type: %s", playFor(perf).PlayType)
	}

	return result, nil
}

func volumeCreditsFor(perf Performance) int {
	result := 0
	result += int(math.Max(float64(perf.Audience-30), 0.0))
	if playFor(perf).PlayType == "comedy" {
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
		[]Performance{
			{"hamlet", 55},
		},
	}

	result, err := statement(invoice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
