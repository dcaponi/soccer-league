package ranker

import "strings"

// Outcome represents the team on the tournament leaderboard
type Outcome struct {
	Team    string
	Score   int
	Ranking int
}

type byScore []Outcome

func (a byScore) Len() int           { return len(a) }
func (a byScore) Less(i, j int) bool { return a[i].Score > a[j].Score }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type byName []Outcome

func (a byName) Len() int      { return len(a) }
func (a byName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool {
	return strings.Compare(strings.ToLower(a[i].Team), strings.ToLower(a[j].Team)) < 0
}
