package master

import "github.com/play-with-golang/refactoring/chap1/model"

var plays = map[string]model.Play{
	"hamlet":  {Name: "Hamlet", PlayType: "tragedy"},
	"as-like": {Name: "As You Like It", PlayType: "comedy"},
	"othello": {Name: "Othello", PlayType: "tragedy"},
}
