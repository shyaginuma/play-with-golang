package model

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

type EnrichedPerformance struct {
	Performance
	Play          Play
	Amount        int
	VolumeCredits int
}
