package ranker

import (
	"bufio"
	"io"
	"sort"
	"strings"
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

	team_points := map[string]int{}

	for sc.Scan() {
		teams := strings.Split(sc.Text(), ", ")

		a_team := strings.Split(teams[0], " ")
		b_team := strings.Split(teams[1], " ")

		a_team_name, a_team_score := strings.Join(a_team[0:len(a_team)-1], " "), a_team[len(a_team)-1]
		b_team_name, b_team_score := strings.Join(b_team[0:len(b_team)-1], " "), b_team[len(b_team)-1]

		if a_team_score > b_team_score {
			team_points[a_team_name] += 3
			team_points[b_team_name] += 0
		} else if a_team_score < b_team_score {
			team_points[a_team_name] += 0
			team_points[b_team_name] += 3
		} else if a_team_score == b_team_score {
			team_points[a_team_name] += 1
			team_points[b_team_name] += 1
		}
	}
	outcomes := make([]Outcome, len(team_points))
	i := 0
	for k, v := range team_points {
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
				t[len(t)-count] = raw[j]
				j++
				count--
			}

			sort.Sort(byName(t))

			// re-insert alphabetized items into their correct slots in original array
			for _, o := range t {
				o.Ranking = rank
				raw[i] = o
				i++
			}

		} else {
			i++
		}
	}
}
