package ranker

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"sync"
)

// CollectOutcomes reads from a file (or any structure that implements io.Reader) and
// parses the input line-by-line then splits each line on the ", " string to get
// a and b teams. team name and score are ascertained by splitting on the space " " character
// and assuming the last string on each team string is the score while the rest of the strings
// are re-joined with a space character for preservation. Finally an Outcome is assembled by
// assigning points based on the rules (3 for a win, 1 for a draw and none otherwise).
func CollectOutcomes(inputFile io.Reader) []Outcome {
	sc := bufio.NewScanner(inputFile)
	sc.Split(bufio.ScanLines)

	teamPoints := map[string]int{}

	for sc.Scan() {
		teams := strings.Split(sc.Text(), ", ")

		aTeam := strings.Split(teams[0], " ")
		bTeam := strings.Split(teams[1], " ")

		aTeamNoints, aTeamScore := strings.Join(aTeam[0:len(aTeam)-1], " "), aTeam[len(aTeam)-1]
		bTeamNoints, bTeamScore := strings.Join(bTeam[0:len(bTeam)-1], " "), bTeam[len(bTeam)-1]

		if aTeamScore > bTeamScore {
			teamPoints[aTeamNoints] += 3
			teamPoints[bTeamNoints] += 0
		} else if aTeamScore < bTeamScore {
			teamPoints[aTeamNoints] += 0
			teamPoints[bTeamNoints] += 3
		} else if aTeamScore == bTeamScore {
			teamPoints[aTeamNoints] += 1
			teamPoints[bTeamNoints] += 1
		}
	}
	outcomes := make([]Outcome, len(teamPoints))
	i := 0
	for k, v := range teamPoints {
		outcomes[i] = Outcome{Team: k, Score: v}
		i++
	}
	return outcomes
}

// SortOutcomes sorts an array of Outcomes first by Score using the ByScore Comparator
// then by Team alphabetically if n Teams have the same number of points.
// This iterates through the sorted array looking for Teams with the same score,
// counts the number of occurrences, allocates a temporary array and
// sorts that using the ByName comparator before inserting the sorted sub array
// items back into their correct place in the main array.
func SortOutcomes(raw []Outcome) {
	sort.Sort(byScore(raw))

	i, rank := 0, 0
	wg := new(sync.WaitGroup)
	for i < len(raw) {
		rank++
		raw[i].Ranking = rank
		// if current and next outcomes have same score
		if i < len(raw)-1 && raw[i].Score == raw[i+1].Score {
			count, j := 0, i

			// move j forward until outcome j is different score while counting items to alphabetize
			for j < len(raw) && raw[i].Score == raw[j].Score {
				j++
				count++
			}

			// create sub-array to alphabetize and reset j
			t := make([]Outcome, count) // figure out count ahead of time to reduce dynamic allocations (i.e. avoid append)
			j = i

			// insert items to alphabetize
			for count > 0 {
				raw[j].Ranking = rank
				t[len(t)-count] = raw[j]
				j++
				count--
			}
			wg.Add(1) // tells wait group to wait for another goroutine
			go func(n int, src, toSort []Outcome, wg *sync.WaitGroup) {
				defer wg.Done() // report done at the end of the routine
				sort.Sort(byName(t))

				// re-insert alphabetized items into their correct slots in original array
				for _, o := range toSort {
					src[n] = o
					n++
				}
			}(i, raw, t, wg)
			i = j

		} else {
			i++
		}
	}
	wg.Wait() // blocks the SortOutcomes function until all outstanding goroutines report in done
}
