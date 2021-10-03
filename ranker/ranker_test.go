package ranker

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectOutcomes(t *testing.T) {
	tests := map[string]struct {
		InputFile      io.Reader
		ExpectedOutput []Outcome
	}{
		"it collects the outcomes of the matches": {
			InputFile: strings.NewReader("Robots 3, Spammers 3\nThieves 4 lyfe 😈 2 1, FC Fraudsters 0\nRobots 1, FC Fraudsters 1\nSpammers 1, Thieves 4 lyfe 😈 2 3\nRobots 4, Grandparents 0\n"),
			ExpectedOutput: []Outcome{
				{Team: "Thieves 4 lyfe 😈 2", Score: 6, Ranking: 0},
				{Team: "Robots", Score: 5, Ranking: 0},
				{Team: "FC Fraudsters", Score: 1, Ranking: 0},
				{Team: "Spammers", Score: 1, Ranking: 0},
				{Team: "Grandparents", Score: 0, Ranking: 0},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			outcomes := CollectOutcomes(test.InputFile)
			for _, outcome := range test.ExpectedOutput {
				assert.Contains(t, outcomes, outcome)
			}
		})
	}
}

func TestSortOutcomes(t *testing.T) {
	tests := map[string]struct {
		Input          []Outcome
		ExpectedOutput []Outcome
	}{
		"it sorts the outcomes of the matches with no tie": {
			Input: []Outcome{
				{Team: "FC Fraudsters", Score: 1, Ranking: 0},
				{Team: "Robots", Score: 5, Ranking: 0},
				{Team: "Spammers", Score: 2, Ranking: 0},
				{Team: "Grandparents", Score: 0, Ranking: 0},
				{Team: "Thieves", Score: 6, Ranking: 0},
			},
			ExpectedOutput: []Outcome{
				{Team: "Thieves", Score: 6, Ranking: 1},
				{Team: "Robots", Score: 5, Ranking: 2},
				{Team: "Spammers", Score: 2, Ranking: 3},
				{Team: "FC Fraudsters", Score: 1, Ranking: 4},
				{Team: "Grandparents", Score: 0, Ranking: 5},
			},
		},
		"it sorts the outcomes of the matches alphabetically in case of a tie": {
			Input: []Outcome{
				{Team: "Spammers", Score: 1, Ranking: 0},
				{Team: "FC Fraudsters", Score: 1, Ranking: 0},
				{Team: "Robots", Score: 5, Ranking: 0},
				{Team: "Grandparents", Score: 0, Ranking: 0},
				{Team: "Thieves", Score: 6, Ranking: 0},
			},
			ExpectedOutput: []Outcome{
				{Team: "Thieves", Score: 6, Ranking: 1},
				{Team: "Robots", Score: 5, Ranking: 2},
				{Team: "FC Fraudsters", Score: 1, Ranking: 3},
				{Team: "Spammers", Score: 1, Ranking: 3},
				{Team: "Grandparents", Score: 0, Ranking: 4},
			},
		},
		"it sorts the outcomes of the matches alphabetically in case of 2 ties": {
			Input: []Outcome{
				{Team: "Spammers", Score: 1, Ranking: 0},
				{Team: "FC Fraudsters", Score: 1, Ranking: 0},
				{Team: "Robots", Score: 5, Ranking: 0},
				{Team: "Ding Dongs", Score: 5, Ranking: 0},
				{Team: "Grandparents", Score: 0, Ranking: 0},
				{Team: "Thieves", Score: 6, Ranking: 0},
			},
			ExpectedOutput: []Outcome{
				{Team: "Thieves", Score: 6, Ranking: 1},
				{Team: "Ding Dongs", Score: 5, Ranking: 2},
				{Team: "Robots", Score: 5, Ranking: 2},
				{Team: "FC Fraudsters", Score: 1, Ranking: 3},
				{Team: "Spammers", Score: 1, Ranking: 3},
				{Team: "Grandparents", Score: 0, Ranking: 4},
			},
		},
		"it sorts the outcomes of the matches alphabetically in case of all tie": {
			Input: []Outcome{
				{Team: "Spammers", Score: 1, Ranking: 0},
				{Team: "FC Fraudsters", Score: 1, Ranking: 0},
				{Team: "Robots", Score: 1, Ranking: 0},
				{Team: "Grandparents", Score: 1, Ranking: 0},
				{Team: "Thieves", Score: 1, Ranking: 0},
			},
			ExpectedOutput: []Outcome{
				{Team: "FC Fraudsters", Score: 1, Ranking: 1},
				{Team: "Grandparents", Score: 1, Ranking: 1},
				{Team: "Robots", Score: 1, Ranking: 1},
				{Team: "Spammers", Score: 1, Ranking: 1},
				{Team: "Thieves", Score: 1, Ranking: 1},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			SortOutcomes(test.Input)
			assert.Equal(t, test.ExpectedOutput, test.Input)
		})
	}
}
