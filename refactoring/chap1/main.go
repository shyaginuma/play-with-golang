package main

type Invoice struct {
	Customer     string
	Performances []struct {
		playID   string
		audience int
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

func statement(invoice Invoice) string {
	// TODO: implement
	return plays["hamlet"].Name
}
