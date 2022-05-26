package main

import (
	"math"
	"sort"
)

type Province struct {
	Name      string
	Producers []Producer
	Demand    int
	Price     int
}

type Producer struct {
	Name       string
	Cost       int
	Production int
}

func (p *Province) ShortFall() int {
	return p.Demand - p.TotalProduction()
}

func (p *Province) Profit() int {
	satisfiedDemand := math.Min(float64(p.Demand), float64(p.TotalProduction()))
	demandValue := satisfiedDemand * float64(p.Price)

	demandCost := 0
	remainingDemand := 0
	sort.SliceStable(p.Producers, func(i, j int) bool { return p.Producers[i].Cost < p.Producers[j].Cost })
	for _, producer := range p.Producers {
		contribution := math.Min(float64(remainingDemand), float64(producer.Production))
		remainingDemand -= int(contribution)
		demandCost += int(contribution) * producer.Cost
	}
	return int(demandValue) - demandCost
}

func (p *Province) TotalProduction() int {
	result := 0
	for _, producer := range p.Producers {
		result += producer.Production
	}
	return result
}

func SampleProvinceData() Province {
	return Province{
		Name: "Asia",
		Producers: []Producer{
			{"Byzantium", 10, 9},
			{"Attalia", 12, 10},
			{"Sinope", 10, 6},
		},
		Demand: 30,
		Price:  20,
	}
}
