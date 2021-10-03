ship:
	go build ./... && bash ship.sh github.com/dcaponi/soccer-league-2020

test:
	go get github.com/jpoles1/gopherbadger
	gopherbadger -md="readme.md" -png=false

install:
	go build ./... && go install .

run:
	go build ./... && go run .