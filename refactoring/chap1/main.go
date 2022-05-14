package main

import (
	"fmt"
	"math"
)

type Invoice struct {
	Customer     string
	Performances []struct {
		PlayID   string
		Audience int
	}
}

type Play struct {
	Name      string
	PlaytType string
}

var plays = map[string]Play{
	"hamlet":  {"Hamlet", "tragedy"},
	"as-like": {"As You Like It", "comedy"},
	"othello": {"Othello", "tragedy"},
}

func statement(invoice Invoice) (string, error) {
	totalAmount := 0
	volumeCredit := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0

		switch play.PlaytType {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience
		default:
			return "", fmt.Errorf("unknown type: %s", play.PlaytType)
		}

		volumeCredit += int(math.Max(float64(perf.Audience-30), 0.0))
		if play.PlaytType == "comedy" {
			volumeCredit += int(math.Floor(float64(perf.Audience) / 5))
		}
		result += fmt.Sprintf("\t%s: %v (%v seat)\n", play.Name, thisAmount, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is %v\n", totalAmount)
	result += fmt.Sprintf("You earned %v credits\n", volumeCredit)
	return result, nil
}
