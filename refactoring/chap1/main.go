package main

import (
	"fmt"

	"github.com/play-with-golang/refactoring/chap1/model"
	"github.com/play-with-golang/refactoring/chap1/statement"
)

func main() {
	invoice := model.Invoice{
		Customer: "Yagi",
		Performances: []model.Performance{
			{PlayID: "hamlet", Audience: 55},
		},
	}
	data := statement.NewStatement(invoice)
	fmt.Println(renderPlainText(data))
}

func renderPlainText(data statement.StatementData) string {
	result := fmt.Sprintf("Statement for %s\n", data.Customer)
	for _, perf := range data.Performances {
		result += fmt.Sprintf("\t%s: %v (%v seat)\n", perf.PlayID, perf.Amount, perf.Audience)
	}
	result += fmt.Sprintf("Amount owed is %v\n", data.TotalAmount)
	result += fmt.Sprintf("You earned %v credits\n", data.TotalVolumeCredits)
	return result
}
