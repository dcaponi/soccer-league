package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dcaponi/soccer-league-2020/ranker"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("exactly one filepath is expected")
	}
	filename := os.Args[1]
	fileExtensions := strings.Split(filename, ".")
	fileExtension := fileExtensions[len(fileExtensions)-1]
	if fileExtension != "txt" {
		log.Fatalln("file must be .txt file")
	}

	p := os.Args[1]
	f, err := os.Open(p)
	if err != nil {
		log.Fatalln("unable to open file", p)
	}

	outcomes := ranker.CollectOutcomes(f)
	ranker.SortOutcomes(outcomes)
	for _, outcome := range outcomes {
		unit := "pts"
		if outcome.Score == 1 {
			unit = "pt"
		}
		fmt.Printf("%d. %s, %d %s\n", outcome.Ranking, outcome.Team, outcome.Score, unit)
	}
}
